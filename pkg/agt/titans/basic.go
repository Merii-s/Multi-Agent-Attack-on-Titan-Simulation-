package agt

import (
	pkg "AOT/pkg"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BasicTitanI interface {
	TitanI
}

type BasicTitan struct {
	attributes Titan
	stopCh     chan struct{}
	syncChan   chan string
	mu         sync.Mutex
}

func NewBasicTitan(id pkg.Id, pos pkg.Position, hp int, reach int, speed int, strength int, height int, regen int) *BasicTitan {
	atts := NewTitan(id, pkg.Titan, pos, hp, reach, strength, speed, height, regen)
	return &BasicTitan{attributes: *atts, syncChan: make(chan string), mu: sync.Mutex{}, stopCh: make(chan struct{})}
}

func (*BasicTitan) Percept(e *pkg.Environment) {

}

func (*BasicTitan) Deliberate() {

}

func (*BasicTitan) Act(e *pkg.Environment) {

}

func (*BasicTitan) Start() {

}

func (bt *BasicTitan) Id() pkg.Id {
	return bt.attributes.agentAttributes.Id()
}

func (bt *BasicTitan) move() {
	new_X_pos := bt.attributes.agentAttributes.Pos().X + bt.attributes.agentAttributes.Speed()
	new_Y_pos := bt.attributes.agentAttributes.Pos().Y + bt.attributes.agentAttributes.Speed()
	new_pos := pkg.Position{X: new_X_pos, Y: new_Y_pos}
	bt.attributes.agentAttributes.SetPos(new_pos)
}

func (bt *BasicTitan) eat() {
	// TODO : Eat humans
}

func (*BasicTitan) sleep() {
	// It never sleeps
	time.Sleep(0)
}

// Return a value between 0 and 1 representing success of an attack
func (*BasicTitan) attack_success(spd_atk int, reach_atk int, spd_def int) float64 {
	// If the speed of the attacker is greater than the speed of the defender, the attack is successful
	if spd_atk > spd_def {
		return 1
	} else {
		// If the speed of the attacker is less than the speed of the defender, the attack is successful with a probability of
		// (speed of the attacker)/(speed of the defender)
		return float64(spd_atk) / float64(spd_def)
	}
}

func (bt *BasicTitan) attack(agt pkg.Agent) {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < bt.attack_success(bt.attributes.agentAttributes.Speed(), bt.attributes.agentAttributes.Reach(), agt.Speed()) {
		// If the attack is successful, the agent loses HP
		agt.SetHp(agt.Hp() - bt.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %s : %s lost  %d HP \n", bt.Id(), agt.Id(), agt.Hp())
	} else {
		fmt.Println("Attack unsuccessful.")
		// If the attack is unsuccessful, nothing happens
	}
}

// Regenerate method for BasicTitan
func (bt *BasicTitan) Regenerate() {
	// Create a channel to signal the stop of regeneration
	bt.stopCh = make(chan struct{})

	// Start a goroutine for regeneration
	go func() {
		// Create a ticker that ticks every 10 seconds
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// Regenerate HP only if it's not full
				bt.mu.Lock()
				if bt.attributes.agentAttributes.Hp() < bt.attributes.agentAttributes.MaxHP() {
					bt.attributes.agentAttributes.SetHp(bt.attributes.agentAttributes.Hp() + bt.attributes.RegenRate())
					if bt.attributes.agentAttributes.Hp() > bt.attributes.agentAttributes.MaxHP() {
						bt.attributes.agentAttributes.SetHp(bt.attributes.agentAttributes.MaxHP())
					}
					fmt.Printf("Regenerated HP: %d\n", bt.attributes.agentAttributes.Hp())
				}
				bt.mu.Unlock()

				// Use syncChan to signal regeneration completion
				bt.syncChan <- "RegenerationComplete"
			case <-bt.stopCh:
				// Stop the regeneration goroutine when signaled
				return
			}
		}
	}()
}

// StopRegeneration stops the regeneration process
func (bt *BasicTitan) StopRegeneration() {
	close(bt.stopCh)
}
