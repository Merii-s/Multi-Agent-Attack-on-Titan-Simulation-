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
	pkg.BehaviorI
}

func NewBasicTitan(id pkg.Id, pos pkg.Position, hp int, reach int, speed int, strength int, height int, regen int) *BasicTitan {
	atts := NewTitan(id, pkg.Titan, pos, hp, reach, strength, speed, height, regen)
	return &BasicTitan{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
		BehaviorI:  &BasicTitanBehavior{},
	}
}

// Setter and getter methods for BasicTitan
func (bt *BasicTitan) SyncChan() chan string {
	return bt.syncChan
}

func (bt *BasicTitan) StopCh() chan struct{} {
	return bt.stopCh
}

func (bt *BasicTitan) Behavior() pkg.BehaviorI {
	return bt.BehaviorI
}

func (bt *BasicTitan) SetBehavior(b pkg.BehaviorI) {
	bt.BehaviorI = b
}

// Methods for BasicTitan
func (bt *BasicTitan) Percept(e *pkg.Environment) {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	bt.BehaviorI.Percept(e)
}

func (bt *BasicTitan) Deliberate() {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	bt.BehaviorI.Deliberate()

}

func (bt *BasicTitan) Act(e *pkg.Environment) {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	bt.BehaviorI.Act(e)
}

func (bt *BasicTitan) Start() {
	// launch the agent goroutine Percept-Act cycle
	go func() {
		for {
			// Percept
			// TODO : Percept
			// Deliberate
			// TODO : Deliberate
			// Act
			// TODO : Act
		}
	}()

}

func (bt *BasicTitan) Id() pkg.Id {
	return bt.attributes.agentAttributes.Id()
}

func (bt *BasicTitan) Move(pos pkg.Position) {
	// TODO : Move randomly or towards a target --> not only in a straight line (top right here)
	bt.attributes.agentAttributes.SetPos(pos)
}

func (bt *BasicTitan) Eat() {
	// TODO: Eat humans
}

func (*BasicTitan) Sleep() {
	// It never sleeps
	time.Sleep(0)
}

// Return a value between 0 and 1 representing success of an attack
func (bt *BasicTitan) AttackSuccess(spdAtk int, spdDef int) float64 {
	// If the speed of the attacker is greater than the speed of the defender, the attack is successful
	if spdAtk > spdDef {
		return 1
	} else {
		// If the speed of the attacker is less than the speed of the defender, the attack is successful with a probability of
		// (speed of the attacker)/(speed of the defender)
		return float64(spdAtk) / float64(spdDef)
	}
}

func (bt *BasicTitan) attack(agt pkg.Agent) {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < bt.AttackSuccess(bt.attributes.agentAttributes.Speed(), agt.Speed()) {
		// If the attack is successful, the agent loses HP
		agt.SetHp(agt.Hp() - bt.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %s : %s lost  %d HP \n", bt.Id(), agt.Id(), agt.Hp())
	} else {
		fmt.Println("Attack unsuccessful.")
		// If the attack is unsuccessful, nothing happens
	}
}

func (bt *BasicTitan) Pos() pkg.Position {
	return bt.attributes.agentAttributes.Pos()
}

func (bt *BasicTitan) SeenPositions() map[pkg.Position]pkg.ObjectName {
	return bt.attributes.agentAttributes.SeenPositions()
}

func (bt *BasicTitan) Vision() int {
	return bt.attributes.agentAttributes.Vision()
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

// Define the behavior struct of the BasicTitan :
type BasicTitanBehavior struct {
}

func (btb *BasicTitanBehavior) Percept(e *pkg.Environment) {}

func (btb *BasicTitanBehavior) Deliberate() {}

func (btb *BasicTitanBehavior) Act(e *pkg.Environment) {
}
