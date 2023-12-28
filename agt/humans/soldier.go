package humans

import (
	env "AOT/agt/env"
	obj "AOT/pkg/obj"
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"
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
	behavior   env.BehaviorI
}

func NewSoldier(id types.Id, tl types.Position, life int, reach int, strength int, speed int, vision int, obj types.ObjectName) *Soldier {
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

func (s *Soldier) Behavior() *env.BehaviorI {
	return &s.behavior
}

func (s *Soldier) SetBehavior(b env.BehaviorI) {
	s.behavior = b
}

// Methods for Soldier
func (s *Soldier) Percept(e *env.Environment) {
	s.behavior.Percept(e)
}

func (s *Soldier) Deliberate() {
	s.behavior.Deliberate()
}

func (s *Soldier) Act(e *env.Environment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.behavior.Act(e)
}

func (s *Soldier) Id() types.Id {
	return s.attributes.agentAttributes.Id()
}

func (s *Soldier) Agent() *env.Agent {
	return &s.attributes.agentAttributes
}

func (s *Soldier) Start(e *env.Environment) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	go func() {
		for {
			println("Soldier Start")
			s.behavior.Percept(e)
			time.Sleep(100 * time.Millisecond)
			s.behavior.Deliberate()
			s.behavior.Act(e)
		}
	}()
}

func (s *Soldier) Move(pos types.Position) {
	// TODO : Move randomly or towards a target --> not only in a straight line (top right here)
	s.attributes.agentAttributes.SetPos(pos)
}

func (s *Soldier) Eat() {

}

func (s *Soldier) Sleep() {
	//time.Sleep(?)
}

func (*Soldier) Gard() {

}

// Return a value between 0 and 1 representing success of an attack
func (*Soldier) AttackSuccess(spdAtk int, spdDef int) float64 {
	// If the speed of the attacker is greater than the speed of the defender, the attack is successful
	if spdAtk > spdDef {
		return 1
	} else {
		// If the speed of the attacker is less than the speed of the defender, the attack is successful with a probability of
		// (speed of the attacker)/(speed of the defender)
		return float64(spdAtk) / float64(spdDef)
	}
}

func (s *Soldier) Attack(agt env.AgentI) {
	if agt.Id() != s.attributes.agentAttributes.Id() {
		s.mu.Lock()
		defer s.mu.Unlock()
		// TODO: consider the reach of the agent
		// If the percentage is less than the success rate, the attack is successful
		if rand.Float64() < s.AttackSuccess(s.attributes.agentAttributes.Speed(), agt.Agent().Speed()) {
			// If the attack is successful, the agent loses HP
			agt.Agent().SetHp(agt.Agent().Hp() - s.attributes.agentAttributes.Strength())
			fmt.Printf("Attack successful from %s : %s lost  %d HP \n", s.Id(), agt.Id(), agt.Agent().Hp())
		} else {
			fmt.Println("Attack unsuccessful.")
			// If the attack is unsuccessful, nothing happens
		}
	}
}

func (s *Soldier) SetPos(pos types.Position) { s.attributes.agentAttributes.SetPos(pos) }

func (s *Soldier) Pos() types.Position {
	return s.attributes.agentAttributes.Pos()
}

func (s *Soldier) Vision() int {
	return s.attributes.agentAttributes.Vision()
}

func (s *Soldier) Object() obj.Object {
	return s.attributes.agentAttributes.Object()
}

func (s *Soldier) PerceivedObjects() []obj.Object {
	return s.attributes.agentAttributes.PerceivedObjects()
}

func (s *Soldier) PerceivedAgents() []env.AgentI {
	return s.attributes.agentAttributes.PerceivedAgents()
}

// Define the behavior struct of the Soldier :
type SoldierBehavior struct {
	s *Soldier
}

func (sb *SoldierBehavior) Percept(e *env.Environment) {
	println("Soldier Percept")
	// Get the perceived objects and agents
	perceivedObjects, perceivedAgents := sb.s.attributes.agentAttributes.GetVision(e)

	// Add the percepted agents to the list of percepted agents
	for _, obj := range perceivedObjects {
		fmt.Printf("Percepted object: %s\n", obj.GetName())
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
	var agentToAttack env.AgentI

	// Count the number of titans and store the position of the titan to attack
	for _, agt := range sb.s.attributes.agentAttributes.PerceivedAgents() {
		// if the agent is a special titan, the soldier moves away in the opposite direction
		if agt.Object().GetName() == "BeastTitan" || agt.Object().GetName() == "ColossalTitan" || agt.Object().GetName() == "ArmoredTitan" || agt.Object().GetName() == "FemaleTitan" || agt.Object().GetName() == "JawTitan" {
			sb.s.attributes.agentAttributes.SetNextPos(utils.OppositeDirection(sb.s.attributes.agentAttributes.Pos(), agt.Pos()))
			break
		}
		if agt.Object().GetName() == "BasicTitan1" || agt.Object().GetName() == "BasicTitan2" {
			numTitans++
			if numTitans == 1 {
				agentToAttack = agt
			}
		}
	}
	// Decide action based on the number of titans
	if numTitans < 2 {
		sb.s.attributes.agentAttributes.SetAgentToAttack(agentToAttack)
		sb.s.attributes.agentAttributes.SetAttack(true)
		sb.s.attributes.agentAttributes.SetNextPos(agentToAttack.Pos())
	} else {
		sb.s.attributes.agentAttributes.SetNextPos(utils.OppositeDirection(sb.s.attributes.agentAttributes.Pos(), agentToAttack.Pos()))
	}
}

func (sb *SoldierBehavior) Act(e *env.Environment) {
	// Perform the action based on the parameters
	if sb.s.attributes.agentAttributes.Attack() {
		sb.s.Move(sb.s.attributes.agentAttributes.NextPos())
		sb.s.Attack(sb.s.attributes.agentAttributes.AgentToAttack())
		// Reset the parameters
		sb.s.attributes.agentAttributes.SetAttack(false)
		sb.s.attributes.agentAttributes.SetAgentToAttack(nil)
	} else {
		// Move towards the specified position
		sb.s.Move(sb.s.attributes.agentAttributes.NextPos())
	}
}
