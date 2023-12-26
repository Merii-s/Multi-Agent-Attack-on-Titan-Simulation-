package agt

import (
	pkg "AOT/pkg"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SoldierI interface {
	HumanI
}

type Soldier struct {
	attributes Human
	stopCh     chan struct{}
	syncChan   chan string
	mu         sync.Mutex
	bahavior   pkg.BehaviorI
}

func NewSoldier(id pkg.Id, tl pkg.Position, life int, reach int, strength int, speed int, vision int, obj pkg.ObjectName) *Soldier {
	atts := NewHuman(id, tl, life, reach, strength, speed, vision, obj)
	s := &Soldier{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
	}
	behavior := &SoldierBehavior{s: s}
	s.SetBehavior(behavior)
	return s
}

// Setter and getter methods for Soldier
func (s *Soldier) SyncChan() chan string {
	return s.syncChan
}

func (s *Soldier) StopCh() chan struct{} {
	return s.stopCh
}

func (s *Soldier) Behavior() *pkg.BehaviorI {
	return &s.bahavior
}

func (s *Soldier) SetBehavior(b pkg.BehaviorI) {
	s.bahavior = b
}

// Methods for Soldier
func (s *Soldier) Percept(e *pkg.Environment) {
	s.bahavior.Percept(e)
}

func (s *Soldier) Deliberate() {
	s.bahavior.Deliberate()
}

func (s *Soldier) Act(e *pkg.Environment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bahavior.Act(e)
}

func (s *Soldier) Id() pkg.Id {
	return s.attributes.agentAttributes.Id()
}

func (s *Soldier) Agent() *pkg.Agent {
	return &s.attributes.agentAttributes
}

func (s *Soldier) Start(e *pkg.Environment) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	go func() {
		for {
			println("Soldier Start")
			s.bahavior.Percept(e)
			time.Sleep(100 * time.Millisecond)
			s.bahavior.Deliberate()
			s.bahavior.Act(e)
		}
	}()
}

func (s *Soldier) move(pos pkg.Position) {
	// TODO : Move randomly or towards a target --> not only in a straight line (top right here)
	s.attributes.agentAttributes.SetPos(pos)
}

func (s *Soldier) eat() {

}

func (s *Soldier) sleep() {
	//time.Sleep(?)
}

func (*Soldier) Gard() {

}

// Return a value between 0 and 1 representing success of an attack
func (*Soldier) attack_success(spd_atk int, reach_atk int, spd_def int) float64 {
	// If the speed of the attacker is greater than the speed of the defender, the attack is successful
	if spd_atk > spd_def {
		return 1
	} else {
		// If the speed of the attacker is less than the speed of the defender, the attack is successful with a probability of
		// (speed of the attacker)/(speed of the defender)
		return float64(spd_atk) / float64(spd_def)
	}
}

func (s *Soldier) Attack(agt pkg.AgentI) {
	if agt.Id() != s.attributes.agentAttributes.Id() {
		s.mu.Lock()
		defer s.mu.Unlock()
		// If the percentage is less than the success rate, the attack is successful
		if rand.Float64() < s.attack_success(s.attributes.agentAttributes.Speed(), s.attributes.agentAttributes.Reach(), agt.Agent().Speed()) {
			// If the attack is successful, the agent loses HP
			agt.Agent().SetHp(agt.Agent().Hp() - s.attributes.agentAttributes.Strength())
			fmt.Printf("Attack successful from %s : %s lost  %d HP \n", s.Id(), agt.Id(), agt.Agent().Hp())
		} else {
			fmt.Println("Attack unsuccessful.")
			// If the attack is unsuccessful, nothing happens
		}
	}
}

func (s *Soldier) Pos() pkg.Position {
	return s.attributes.agentAttributes.Pos()
}

func (s *Soldier) Vision() int {
	return s.attributes.agentAttributes.Vision()
}

func (s *Soldier) Object() pkg.Object {
	return s.attributes.agentAttributes.Object()
}

func (s *Soldier) PerceivedObjects() []pkg.Object {
	return s.attributes.agentAttributes.PerceivedObjects()
}

func (s *Soldier) PerceivedAgents() []pkg.AgentI {
	return s.attributes.agentAttributes.PerceivedAgents()
}

// Define the behavior struct of the Soldier :
type SoldierBehavior struct {
	s *Soldier
}

func (sb *SoldierBehavior) Percept(e *pkg.Environment) {
	println("Soldier Percept")
	// Get the perceived objects and agents
	perceivedObjects, perceivedAgents := sb.s.attributes.agentAttributes.GetVision(e)

	// Add the percepted agents to the list of percepted agents
	for _, obj := range perceivedObjects {
		fmt.Printf("Percepted object: %s\n", obj.Name())
		sb.s.attributes.agentAttributes.AddPerceivedObject(obj)
	}

	// Add the percepted agents to the list of percepted agents
	for _, agt := range perceivedAgents {
		fmt.Printf("Percepted agent: %s\n", agt.Id())
		sb.s.attributes.agentAttributes.AddPerceivedAgent(agt)
	}

	time.Sleep(100 * time.Millisecond)
}

func (sb *SoldierBehavior) Deliberate() {
	// Initialize variables for counting titans
	numTitans := 0
	var firstTitanPos pkg.Position
	idAgentToAttack := sb.s.Id()

	// TO DO: randomize the choice

	// Count the number of titans and store the position of the first titan
	for _, agt := range sb.s.attributes.agentAttributes.PerceivedAgents() {
		// if the agent is a special titan, the soldier moves away in the opposite direction
		if agt.Object().Name() == "BeastTitan" || agt.Object().Name() == "ColossalTitan" || agt.Object().Name() == "ArmoredTitan" || agt.Object().Name() == "FemaleTitan" || agt.Object().Name() == "JawTitan" {
			// TO DO: Move away from the titan
		}
		if agt.Object().Name() == "BasicTitan1" || agt.Object().Name() == "BasicTitan2" {
			numTitans++
			if numTitans == 1 {
				firstTitanPos = agt.Pos()
				AgentToAttack := agt
			}
		}
	}
	// Decide action based on the number of titans
	if numTitans < 2 {
		sb.s.Attack(AgentToAttack.Agent())
	} else {
		// TO DO: Move away
	}
}

func (sb *SoldierBehavior) Act(e *pkg.Environment) {
	// Perform the action based on the parameters
	if attack {
		sb.s.attributes.agentAttributes.SetPos(pos)
		sb.s.attack(e.GetAgent(agtId))
	} else {
		// Move towards the specified position
		sb.s.move(pos)
	}
}
