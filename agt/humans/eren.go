package humans

import (
	env "AOT/agt/env"
	obj "AOT/pkg/obj"
	types "AOT/pkg/types"
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
	behavior   env.BehaviorI
	transform  bool
}

func NewEren(id types.Id, tl types.Position, life int, reach int, strength int, speed int, vision int, obj types.ObjectName) *Eren {
	atts := NewHuman(id, tl, life, reach, strength, speed, vision, obj)
	eren := &Eren{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
		transform:  false,
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

func (eren *Eren) Behavior() *env.BehaviorI {
	return &eren.behavior
}

func (eren *Eren) SetBehavior(b env.BehaviorI) {
	eren.behavior = b
}

// Methods for Eren
func (eren *Eren) Percept(e *env.Environment, wgPercept *sync.WaitGroup) {
	defer wgPercept.Done()
	eren.behavior.Percept(e)
}

func (eren *Eren) Deliberate(wgDeliberate *sync.WaitGroup) {
	defer wgDeliberate.Done()
	eren.behavior.Deliberate()
}

func (eren *Eren) Act(e *env.Environment, wgAct *sync.WaitGroup) {
	defer wgAct.Done()
	eren.mu.Lock()
	defer eren.mu.Unlock()
	eren.behavior.Act(e)
}

func (eren *Eren) Id() types.Id {
	return eren.attributes.agentAttributes.Id()
}

func (eren *Eren) Agent() *env.Agent {
	return &eren.attributes.agentAttributes
}

func (eren *Eren) Start(e *env.Environment, wgStart *sync.WaitGroup, wgPercept *sync.WaitGroup, wgDeliberate *sync.WaitGroup, wgAct *sync.WaitGroup) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	go func() {
		println("Eren Start")
		for {
			wgStart.Done()
			wgStart.Wait()

			wgPercept.Add(1)
			eren.Percept(e, wgPercept)
			wgPercept.Wait()

			wgDeliberate.Add(1)
			eren.Deliberate(wgDeliberate)
			wgDeliberate.Wait()

			wgAct.Add(1)
			eren.Act(e, wgAct)
			wgAct.Wait()
		}
	}()
}

func (eren *Eren) Move(pos types.Position) {
	eren.attributes.agentAttributes.SetPos(pos)
}

func (eren *Eren) Eat() {

}

func (eren *Eren) Sleep() {

}

func (*Eren) Guard() {

}

func (*Eren) AttackSuccess(spdAtk int, spdDef int) float64 {
	// If the speed of the attacker is greater than the speed of the defender, the attack is successful
	if spdAtk > spdDef {
		return 1
	} else {
		// If the speed of the attacker is less than the speed of the defender, the attack is successful with a probability of
		// (speed of the attacker)/(speed of the defender)
		return float64(spdAtk) / float64(spdDef)
	}
}

func (eren *Eren) Attack(agt env.AgentI) {
	eren.mu.Lock()
	defer eren.mu.Unlock()
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < eren.AttackSuccess(eren.attributes.agentAttributes.Speed(), agt.Agent().Speed()) {
		// If the attack is successful, the agent loses HP
		agt.Agent().SetHp(agt.Agent().Hp() - eren.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %sren : %sren lost  %d HP \n", eren.Id(), agt.Id(), agt.Agent().Hp())
	} else {
		fmt.Println("Attack unsuccessful.")
		// If the attack is unsuccessful, nothing happens
	}
}

func (eren *Eren) Pos() types.Position {
	return eren.attributes.agentAttributes.Pos()
}

func (eren *Eren) Vision() int {
	return eren.attributes.agentAttributes.Vision()
}

func (eren *Eren) Object() obj.Object {
	return eren.attributes.agentAttributes.Object()
}

func (eren *Eren) PerceivedObjects() []obj.Object {
	return eren.attributes.agentAttributes.PerceivedObjects()
}

func (eren *Eren) PerceivedAgents() []env.AgentI {
	return eren.attributes.agentAttributes.PerceivedAgents()
}

func (eren *Eren) SetPos(pos types.Position) { eren.attributes.agentAttributes.SetPos(pos) }

// Define the behavior struct of Eren
type ErenBehavior struct {
	eren *Eren
}

func (eb *ErenBehavior) Percept(e *env.Environment) {
	println("Eren Percept")
	// Get the perceived objects and agents
	perceivedObjects, perceivedAgents := eb.eren.attributes.agentAttributes.GetVision(e)

	// Add the percepted agents to the list of percepted agents
	for _, obj := range perceivedObjects {
		fmt.Printf("Percepted object: %sren\n", obj.Name())
		eb.eren.attributes.agentAttributes.AddPerceivedObject(obj)
	}

	// Add the percepted agents to the list of percepted agents
	for _, agt := range perceivedAgents {
		fmt.Printf("Percepted agent: %sren\n", agt.Id())
		eb.eren.attributes.agentAttributes.AddPerceivedAgent(agt)
	}

	time.Sleep(100 * time.Millisecond)
}

func (eb *ErenBehavior) Deliberate() {
	// Initialize variables for counting titans
	numTitans := 0
	var agentToAttack env.AgentI

	// Count the number of titans and store the position of the titan to attack
	for _, agt := range eb.eren.attributes.agentAttributes.PerceivedAgents() {
		//if the agent is a special titan, Eren decides to transform to titan and attacks
		if agt.Object().GetName() == "BeastTitan" || agt.Object().GetName() == "ColossalTitan" || agt.Object().GetName() == "ArmoredTitan" || agt.Object().GetName() == "FemaleTitan" || agt.Object().GetName() == "JawTitan" {
			agentToAttack = agt
			eb.eren.attributes.agentAttributes.SetAttack(true)
			//Eren decides to transform to titan
			eb.eren.transform = true
			break
		}
		if agt.Object().GetName() == "BasicTitan1" || agt.Object().GetName() == "BasicTitan2" {
			numTitans++
			if numTitans == 1 {
				agentToAttack = agt
				eb.eren.attributes.agentAttributes.SetAttack(true)
			}
		}
	}

	// Decide action based on the number of titans
	if numTitans > 1 {
		//Eren decides to transform to titan
		eb.eren.transform = true
	}

	if eb.eren.attributes.agentAttributes.Attack() {
		eb.eren.attributes.agentAttributes.SetAgentToAttack(agentToAttack)
		eb.eren.attributes.agentAttributes.SetNextPos(agentToAttack.Pos())
	} else {
		// Move randomly
		var randPos types.Position
		randPos.X = eb.eren.attributes.agentAttributes.Pos().X + rand.Intn(10)
		randPos.Y = eb.eren.attributes.agentAttributes.Pos().Y + rand.Intn(10)

		eb.eren.attributes.agentAttributes.SetNextPos(randPos)
	}
}

func (eb *ErenBehavior) Act(e *env.Environment) {
	if eb.eren.transform {
		//TODO: Eren transforms to titan
	}

	if eb.eren.attributes.agentAttributes.Attack() {
		eb.eren.Move(eb.eren.attributes.agentAttributes.NextPos())
		eb.eren.Attack(eb.eren.attributes.agentAttributes.AgentToAttack())
		// Reset the parameters
		eb.eren.attributes.agentAttributes.SetAttack(false)
		eb.eren.attributes.agentAttributes.SetAgentToAttack(nil)
		if eb.eren.transform {
			eb.eren.transform = false
			//TODO: Eren transforms back to human
		}
	} else {
		// Move towards the specified position
		eb.eren.Move(eb.eren.attributes.agentAttributes.NextPos())
	}
}
