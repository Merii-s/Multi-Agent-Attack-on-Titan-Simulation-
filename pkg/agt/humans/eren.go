package agt

import (
	pkg "AOT/pkg"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ErenI interface {
	HumanI
}

type Eren struct {
	attributes Human
	stopCh     chan struct{}
	syncChan   chan string
	mu         sync.Mutex
	behavior   pkg.BehaviorI
}

func NewEren(id pkg.Id, tl pkg.Position, life int, reach int, strength int, speed int, vision int, obj pkg.ObjectName) *Eren {
	atts := NewHuman(id, tl, life, reach, strength, speed, vision, obj)
	eren := &Eren{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
	}
	behavior := &ErenBehavior{eren: eren}
	eren.SetBehavior(behavior)
	return eren
}

// Setter and getter methods for Eren
func (eren *Eren) SyncChan() chan string {
	return eren.syncChan
}

func (eren *Eren) StopCh() chan struct{} {
	return eren.stopCh
}

func (eren *Eren) Behavior() *pkg.BehaviorI {
	return &eren.behavior
}

func (eren *Eren) SetBehavior(b pkg.BehaviorI) {
	eren.behavior = b
}

// Methods for Eren
func (eren *Eren) Percept(e *pkg.Environment) {
	eren.behavior.Percept(e)
}

func (eren *Eren) Deliberate() {
	eren.behavior.Deliberate()
}

func (eren *Eren) Act(e *pkg.Environment) {
	eren.mu.Lock()
	defer eren.mu.Unlock()
	eren.behavior.Act(e)
}

func (eren *Eren) Id() pkg.Id {
	return eren.attributes.agentAttributes.Id()
}

func (eren *Eren) Start(e *pkg.Environment) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	go func() {
		for {
			println("Eren Start")
			eren.behavior.Percept(e)
			time.Sleep(100 * time.Millisecond)
			eren.behavior.Deliberate()
			eren.behavior.Act(e)
		}
	}()
}

func (eren *Eren) move(pos pkg.Position) {
	eren.attributes.agentAttributes.SetPos(pos)
}

func (eren *Eren) eat() {

}

func (eren *Eren) sleep() {

}

func (*Eren) Guard() {

}

func (*Eren) attack_success(spd_atk int, reachAtk int, spd_def int) float64 {
	// If the speed of the attacker is greater than the speed of the defender, the attack is successful
	if spd_atk > spd_def {
		return 1
	} else {
		// If the speed of the attacker is less than the speed of the defender, the attack is successful with a probability of
		// (speed of the attacker)/(speed of the defender)
		return float64(spd_atk) / float64(spd_def)
	}
}

func (eren *Eren) attack(agt pkg.Agent) {
	eren.mu.Lock()
	defer eren.mu.Unlock()
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < eren.attack_success(eren.attributes.agentAttributes.Speed(), eren.attributes.agentAttributes.Reach(), agt.Speed()) {
		// If the attack is successful, the agent loses HP
		agt.SetHp(agt.Hp() - eren.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %eren : %eren lost  %d HP \n", eren.Id(), agt.Id(), agt.Hp())
	} else {
		fmt.Println("Attack unsuccessful.")
		// If the attack is unsuccessful, nothing happens
	}
}

func (eren *Eren) Pos() pkg.Position {
	return eren.attributes.agentAttributes.Pos()
}

func (eren *Eren) Vision() int {
	return eren.attributes.agentAttributes.Vision()
}

func (eren *Eren) Object() pkg.Object {
	return eren.attributes.agentAttributes.Object()
}

func (eren *Eren) PerceivedObjects() []pkg.Object {
	return eren.attributes.agentAttributes.PerceivedObjects()
}

func (eren *Eren) PerceivedAgents() []pkg.AgentI {
	return eren.attributes.agentAttributes.PerceivedAgents()
}

// Define the behavior struct of Eren
type ErenBehavior struct {
	eren *Eren
}

func (eb *ErenBehavior) Percept(e *pkg.Environment) {
	println("Eren Percept")
	// Get the perceived objects and agents
	perceivedObjects, perceivedAgents := eb.eren.attributes.agentAttributes.GetVision(e)

	// Add the percepted agents to the list of percepted agents
	for _, obj := range perceivedObjects {
		fmt.Printf("Percepted object: %eren\n", obj.Name())
		eb.eren.attributes.agentAttributes.AddPerceivedObject(obj)
	}

	// Add the percepted agents to the list of percepted agents
	for _, agt := range perceivedAgents {
		fmt.Printf("Percepted agent: %eren\n", agt.Id())
		eb.eren.attributes.agentAttributes.AddPerceivedAgent(agt)
	}

	time.Sleep(100 * time.Millisecond)
}

func (eb *ErenBehavior) Deliberate() {
	// TODO: Implement deliberation logic for Eren
}

func (eb *ErenBehavior) Act(e *pkg.Environment) {
	// TODO: Implement action logic for Eren
}
