package agt

import (
	pkg "AOT/pkg"
	"fmt"
	"math/rand"
	"sync"
)

type SoldierI interface {
	HumanI
}

type Soldier struct {
	attributes Human
	mu         sync.Mutex
	pkg.BehaviorI
}

func NewSoldier(id pkg.Id, t pkg.Type, topLeft pkg.Position, hp int, reach int, strength int, speed int) *Soldier {
	atts := NewHuman(id, t, topLeft, hp, reach, strength, speed)
	return &Soldier{
		attributes: *atts,
		mu:         sync.Mutex{},
		BehaviorI:  &SoldierBehavior{},
	}
}

func (bt *Soldier) Id() pkg.Id {
	return bt.attributes.agentAttributes.Id()
}

func (*Soldier) Start() {

}

func (*Soldier) move() {

}

func (*Soldier) eat() {

}

func (*Soldier) sleep() {

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

func (s *Soldier) attack(agt pkg.Agent) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < s.attack_success(s.attributes.agentAttributes.Speed(), s.attributes.agentAttributes.Reach(), agt.Speed()) {
		// If the attack is successful, the agent loses HP
		agt.SetHp(agt.Hp() - s.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %s : %s lost  %d HP \n", s.Id(), agt.Id(), agt.Hp())
	} else {
		fmt.Println("Attack unsuccessful.")
		// If the attack is unsuccessful, nothing happens
	}
}

// Define the behavior struct of the BasicTitan :
type SoldierBehavior struct {
}

func (btb *SoldierBehavior) Percept(e *pkg.Environment) {

}

func (btb *SoldierBehavior) Deliberate() {

}

func (btb *SoldierBehavior) Act(e *pkg.Environment) {

}
