package agt

import (
	pkg "AOT/pkg"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MikasaI interface {
	HumanI
}

type Mikasa struct {
	attributes Human
	stopCh     chan struct{}
	syncChan   chan string
	mu         sync.Mutex
	behavior   pkg.BehaviorI
}

func NewMikasa(id pkg.Id, tl pkg.Position, life int, reach int, strength int, speed int, vision int, obj pkg.ObjectName) *Mikasa {
	atts := NewHuman(id, tl, life, reach, strength, speed, vision, obj)
	m := &Mikasa{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
	}
	behavior := &MikasaBehavior{m: m}
	m.SetBehavior(behavior)
	return m
}

// Setter and getter methods for Mikasa
func (m *Mikasa) SyncChan() chan string {
	return m.syncChan
}

func (m *Mikasa) StopCh() chan struct{} {
	return m.stopCh
}

func (m *Mikasa) Behavior() *pkg.BehaviorI {
	return &m.behavior
}

func (m *Mikasa) SetBehavior(b pkg.BehaviorI) {
	m.behavior = b
}

// Methods for Mikasa
func (m *Mikasa) Percept(e *pkg.Environment) {
	m.behavior.Percept(e)
}

func (m *Mikasa) Deliberate() {
	m.behavior.Deliberate()
}

func (m *Mikasa) Act(e *pkg.Environment) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.behavior.Act(e)
}

func (m *Mikasa) Id() pkg.Id {
	return m.attributes.agentAttributes.Id()
}

func (m *Mikasa) Start(e *pkg.Environment) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	go func() {
		for {
			println("Mikasa Start")
			m.behavior.Percept(e)
			time.Sleep(100 * time.Millisecond)
			m.behavior.Deliberate()
			m.behavior.Act(e)
		}
	}()
}

func (m *Mikasa) Move(pos pkg.Position) {
	m.attributes.agentAttributes.SetPos(pos)
}

func (m *Mikasa) Eat() {

}

func (m *Mikasa) Sleep() {

}

func (m *Mikasa) Guard() {

}

func (m *Mikasa) attack_success(spd_atk int, reachAtk int, spd_def int) float64 {
	// If the speed of the attacker is greater than the speed of the defender, the attack is successful
	if spd_atk > spd_def {
		return 1
	} else {
		// If the speed of the attacker is less than the speed of the defender, the attack is successful with a probability of
		// (speed of the attacker)/(speed of the defender)
		return float64(spd_atk) / float64(spd_def)
	}
}

func (m *Mikasa) Attack(agt pkg.Agent) {
	m.mu.Lock()
	defer m.mu.Unlock()
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < m.attack_success(m.attributes.agentAttributes.Speed(), m.attributes.agentAttributes.Reach(), agt.Speed()) {
		// If the attack is successful, the agent loses HP
		agt.SetHp(agt.Hp() - m.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %s : %s lost  %d HP \n", m.Id(), agt.Id(), agt.Hp())
	} else {
		fmt.Println("Attack unsuccessful.")
		// If the attack is unsuccessful, nothing happens
	}
}

// Define the behavior struct of Mikasa
type MikasaBehavior struct {
	m *Mikasa
}

func (mb *MikasaBehavior) Percept(e *pkg.Environment) {

}

func (mb *MikasaBehavior) Deliberate() {

}

func (mb *MikasaBehavior) Act(e *pkg.Environment) {

}
