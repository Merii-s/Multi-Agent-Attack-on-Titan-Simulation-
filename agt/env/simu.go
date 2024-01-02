package env

import (
	"sync"
	"time"
)

type Simulation struct {
	env         *Environment // pb de copie avec les locks...
	agents      []*AgentI
	maxStep     int
	maxDuration time.Duration
	step        int // Stats
	start       time.Time
	syncChans   sync.Map
	uiChan      chan *Environment
}

func NewSimulation(NB_AGENTS int, maxStep int, maxDuration time.Duration, e *Environment, uiC chan *Environment) (simu *Simulation) {
	simu = &Simulation{}
	simu.agents = make([]*AgentI, 0)
	simu.maxStep = maxStep
	simu.maxDuration = maxDuration
	simu.uiChan = uiC

	simu.env = e

	for i := range e.Agts {
		simu.syncChans.Store(e.Agts[i].Id(), e.Agts[i].AgtSyncChan())
		simu.agents = append(simu.agents, &e.Agts[i])
	}

	return simu
}

func (simu *Simulation) Run() {

	// Démarrage des agents
	for i := range simu.agents {
		(*simu.agents[i]).Start(simu.env)
	}

	// On sauvegarde la date du début de la simulation
	simu.start = time.Now()

	// Lancement de l'orchestration de tous les agents
	// simu.step += 1 // plus de sens
	for i := range simu.agents {
		go func(i int) {
			step := 0
			for {
				step++
				c, _ := simu.syncChans.Load((*simu.agents[i]).Id())
				c.(chan int) <- step               // /!\ utilisation d'un "Type Assertion"
				time.Sleep(100 * time.Millisecond) // "cool down"
				<-c.(chan int)
			}
		}(i)
	}

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			simu.uiChan <- simu.env
		}
	}()

	time.Sleep(simu.maxDuration)
}
