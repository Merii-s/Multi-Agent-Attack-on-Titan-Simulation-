package titans

import (
	env "AOT/agt/env"
	obj "AOT/pkg/obj"
	types "AOT/pkg/types"
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
	behavior   env.BehaviorI
}

func NewBasicTitan(id types.Id, tl types.Position, life int, reach int, strength int, speed int, vision int, obj types.ObjectName, regen int) *BasicTitan {
	if obj != types.BasicTitan1 && obj != types.BasicTitan2 {
		return nil
	}
	atts := NewTitan(id, tl, life, reach, strength, speed, vision, obj, regen)
	bt := &BasicTitan{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
	}
	behavior := &BasicTitanBehavior{bt: bt}
	bt.SetBehavior(behavior)
	return bt
}

// Setter and getter methods for BasicTitan
func (bt *BasicTitan) SyncChan() chan string {
	return bt.syncChan
}

func (bt *BasicTitan) StopCh() chan struct{} {
	return bt.stopCh
}

func (bt *BasicTitan) Behavior() *env.BehaviorI {
	return &bt.behavior
}

func (bt *BasicTitan) SetBehavior(b env.BehaviorI) {
	bt.behavior = b
}

// Methods for BasicTitan
func (bt *BasicTitan) Percept(e *env.Environment) {
	bt.behavior.Percept(e)
}

func (bt *BasicTitan) Deliberate() {
	bt.behavior.Deliberate()

}

func (bt *BasicTitan) Act(e *env.Environment) {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	bt.behavior.Act(e)
}

func (bt *BasicTitan) Start(e *env.Environment) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	go func() {
		for {
			println("BasicTitan Start")
			bt.behavior.Percept(e)
			time.Sleep(100 * time.Millisecond)
			bt.behavior.Deliberate()
			bt.behavior.Act(e)
		}
	}()

}

func (bt *BasicTitan) Id() types.Id {
	return bt.attributes.agentAttributes.Id()
}

func (bt *BasicTitan) Move(pos types.Position) {
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

func (bt *BasicTitan) Attack(agt env.AgentI) {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < bt.AttackSuccess(bt.attributes.agentAttributes.Speed(), agt.Agent().Speed()) {
		// If the attack is successful, the agent loses HP
		agt.Agent().SetHp(agt.Agent().Hp() - bt.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %s : %s lost  %d HP \n", bt.Id(), agt.Id(), agt.Agent().Hp())
	} else {
		fmt.Println("Attack unsuccessful.")
		// If the attack is unsuccessful, nothing happens
	}
}

func (bt *BasicTitan) Pos() types.Position {
	return bt.attributes.agentAttributes.Pos()
}

func (bt *BasicTitan) Vision() int {
	return bt.attributes.agentAttributes.Vision()
}

func (bt *BasicTitan) Object() obj.Object {
	return bt.attributes.agentAttributes.Object()
}

func (bt *BasicTitan) PerceivedObjects() []obj.Object {
	return bt.attributes.agentAttributes.PerceivedObjects()
}

func (bt *BasicTitan) PerceivedAgents() []env.AgentI {
	return bt.attributes.agentAttributes.PerceivedAgents()
}

func (bt *BasicTitan) Agent() *env.Agent { return &bt.attributes.agentAttributes }

func (bt *BasicTitan) SetPos(pos types.Position) { bt.attributes.agentAttributes.SetPos(pos) }

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
	bt *BasicTitan
}

func (btb *BasicTitanBehavior) Percept(e *env.Environment) {
	println("BasicTitan Percept")
	// Get the perceived objects and agents
	perceivedObjects, perceivedAgents := btb.bt.attributes.agentAttributes.GetVision(e)

	// Add the percepted agents to the list of percepted agents
	for _, obj := range perceivedObjects {
		fmt.Printf("Percepted object: %s\n", obj.Name())
		btb.bt.attributes.agentAttributes.AddPerceivedObject(obj)
	}

	// Add the percepted agents to the list of percepted agents
	for _, agt := range perceivedAgents {
		fmt.Printf("Percepted agent: %s\n", agt.Id())
		btb.bt.attributes.agentAttributes.AddPerceivedAgent(agt)
	}

	time.Sleep(100 * time.Millisecond)
}

func (btb *BasicTitanBehavior) Deliberate() {}

func (btb *BasicTitanBehavior) Act(e *env.Environment) {
}
