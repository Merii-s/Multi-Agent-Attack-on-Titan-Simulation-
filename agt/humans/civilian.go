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

type CivilianI interface {
	HumanI
	build()
	getFood()
}

type Civilian struct {
	attributes Human
	stopCh     chan struct{}
	syncChan   chan string
	mu         sync.Mutex
	behavior   env.BehaviorI
}

func NewCivilian(id types.Id, tl types.Position, life int, reach int, strength int, speed int, vision int, obj types.ObjectName) *Civilian {
	atts := NewHuman(id, tl, life, reach, strength, speed, vision, obj)
	c := &Civilian{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
	}
	behavior := &CivilianBehavior{c: c}
	c.SetBehavior(behavior)
	return c
}

// Setter and getter methods for Civilian
func (c *Civilian) SyncChan() chan string {
	return c.syncChan
}

func (c *Civilian) StopCh() chan struct{} {
	return c.stopCh
}

func (c *Civilian) Behavior() *env.BehaviorI {
	return &c.behavior
}

func (c *Civilian) SetBehavior(b env.BehaviorI) {
	c.behavior = b
}

// methods for civilian
func (c *Civilian) Percept(e *env.Environment, wgPercept *sync.WaitGroup) {
	defer wgPercept.Done()
	c.behavior.Percept(e)
}

func (c *Civilian) Deliberate(wgDeliberate *sync.WaitGroup) {
	defer wgDeliberate.Done()
	c.behavior.Deliberate()

}

func (c *Civilian) Act(e *env.Environment, wgAct *sync.WaitGroup) {
	defer wgAct.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.behavior.Act(e)
}

func (c *Civilian) Id() types.Id {
	return c.attributes.agentAttributes.Id()
}

func (c *Civilian) Agent() *env.Agent {
	return &c.attributes.agentAttributes
}

func (c *Civilian) Start(e *env.Environment, wgStart *sync.WaitGroup, wgPercept *sync.WaitGroup, wgDeliberate *sync.WaitGroup, wgAct *sync.WaitGroup) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	go func() {
		println("Civilian Start")
		for {
			wgStart.Done()
			wgStart.Wait()

			wgPercept.Add(1)
			c.Percept(e, wgPercept)
			wgPercept.Wait()

			wgDeliberate.Add(1)
			c.Deliberate(wgDeliberate)
			wgDeliberate.Wait()

			wgAct.Add(1)
			c.Act(e, wgAct)
			wgAct.Wait()
		}
	}()
}

func (c *Civilian) Move(pos types.Position) { c.attributes.agentAttributes.SetPos(pos) }

func (c *Civilian) Eat() {

}

func (*Civilian) Sleep() {
	//time.Sleep(?)
}

func (c *Civilian) AttackSuccess(spdAtk int, spdDef int) float64 {
	return 0
}

func (c *Civilian) Attack(agt env.AgentI) { return }

func (c *Civilian) SetPos(pos types.Position) { c.attributes.agentAttributes.SetPos(pos) }

func (c *Civilian) build() {

}

func (c *Civilian) getFood() {

}

func (c *Civilian) Pos() types.Position {
	return c.attributes.agentAttributes.Pos()
}

func (c *Civilian) Vision() int {
	return c.attributes.agentAttributes.Vision()
}

func (c *Civilian) Object() obj.Object {
	return c.attributes.agentAttributes.Object()
}

func (c *Civilian) PerceivedObjects() []obj.Object {
	return c.attributes.agentAttributes.PerceivedObjects()
}

func (c *Civilian) PerceivedAgents() []env.AgentI {
	return c.attributes.agentAttributes.PerceivedAgents()
}

// Define the behavior struct of the Civilian :
type CivilianBehavior struct {
	c *Civilian
}

func (cb *CivilianBehavior) Percept(e *env.Environment) {
	println("Civilian Percept")
	// Get the perceived objects and agents
	perceivedObjects, perceivedAgents := cb.c.attributes.agentAttributes.GetVision(e)

	// Add the percepted agents to the list of percepted agents
	for _, object := range perceivedObjects {
		fmt.Printf("Percepted object: %s\n", object.Name())
		cb.c.attributes.agentAttributes.AddPerceivedObject(object)
	}

	// Add the percepted agents to the list of percepted agents
	for _, agt := range perceivedAgents {
		fmt.Printf("Percepted agent: %s\n", agt.Id())
		cb.c.attributes.agentAttributes.AddPerceivedAgent(agt)
	}

	time.Sleep(100 * time.Millisecond)
}

func (cb *CivilianBehavior) Deliberate() {
	titanFound := false

	for _, agt := range cb.c.attributes.agentAttributes.PerceivedAgents() {
		// if the agent is a titan, the civilian escapes
		if agt.Object().GetName() == "BeastTitan" || agt.Object().GetName() == "ColossalTitan" || agt.Object().GetName() == "ArmoredTitan" || agt.Object().GetName() == "FemaleTitan" || agt.Object().GetName() == "JawTitan" || agt.Object().GetName() == "BasicTitan1" || agt.Object().GetName() == "BasicTitan2" {
			cb.c.attributes.agentAttributes.SetNextPos(utils.OppositeDirection(cb.c.attributes.agentAttributes.Pos(), agt.Pos()))
			titanFound = true
			break
		}
	}
	if !titanFound {
		// Move randomly
		var randPos types.Position
		randPos.X = cb.c.attributes.agentAttributes.Pos().X + rand.Intn(10)
		randPos.Y = cb.c.attributes.agentAttributes.Pos().Y + rand.Intn(10)

		cb.c.attributes.agentAttributes.SetNextPos(randPos)
	}
}

func (cb *CivilianBehavior) Act(e *env.Environment) {
	// Move to the next position
	cb.c.Move(cb.c.attributes.agentAttributes.NextPos())
}
