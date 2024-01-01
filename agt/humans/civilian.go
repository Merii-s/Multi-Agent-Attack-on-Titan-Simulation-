package humans

import (
	env "AOT/agt/env"
	obj "AOT/pkg/obj"
	types "AOT/pkg/types"
	pkg "AOT/pkg/utilitaries"
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
func (c *Civilian) SyncChan() chan string { return c.syncChan }

func (c *Civilian) StopCh() chan struct{} { return c.stopCh }

func (c *Civilian) Behavior() *env.BehaviorI { return &c.behavior }

func (c *Civilian) SetBehavior(b env.BehaviorI) { c.behavior = b }

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
	wgStart.Done()
	wgStart.Wait()
	go func() {
		println("Civilian Start")
		for {
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
	time.Sleep(time.Duration(rand.Intn(30)+10) * time.Second)
}

func (c *Civilian) AttackSuccess(spdAtk int, spdDef int) float64 { return 0 }

func (c *Civilian) Attack(agt *env.AgentI) {
	c.attributes.agentAttributes.SetAgentToAttack(nil)
	c.attributes.agentAttributes.SetAttack(false)
}

func (c *Civilian) SetPos(pos types.Position) { c.attributes.agentAttributes.SetPos(pos) }

// func (c *Civilian) build() {

// }

// func (c *Civilian) getFood() {

// }

func (c *Civilian) Pos() types.Position { return c.attributes.agentAttributes.Pos() }

func (c *Civilian) Vision() int { return c.attributes.agentAttributes.Vision() }

func (c *Civilian) Object() obj.Object { return c.attributes.agentAttributes.Object() }

func (c *Civilian) PerceivedObjects() []*obj.Object {
	return c.attributes.agentAttributes.PerceivedObjects()
}

func (c *Civilian) PerceivedAgents() []*env.AgentI {
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

	// Add the perceived agents to the list of perceived agents
	for i, _ := range perceivedObjects {
		cb.c.attributes.agentAttributes.AddPerceivedObject(perceivedObjects[i])
	}
	// Add the perceived agents to the list of perceived agents
	for i, _ := range perceivedAgents {
		cb.c.attributes.agentAttributes.AddPerceivedAgent(perceivedAgents[i])
	}
	println("Perceived agents: ", len(cb.c.attributes.agentAttributes.PerceivedAgents()))
	println("Perceived objects: ", len(cb.c.attributes.agentAttributes.PerceivedObjects()))

	time.Sleep(100 * time.Millisecond)
}

func (cb *CivilianBehavior) Deliberate() {
	println("Civilian Deliberate")

	var interestingAgents []*env.AgentI
	agtPos := cb.c.attributes.agentAttributes.Pos()

	titanFound := false
	//println("Interesting objects: ", len(interestingObjects))

	for _, agt := range cb.c.attributes.agentAttributes.PerceivedAgents() {
		if (*agt).Agent().GetName() == types.BasicTitan1 ||
			(*agt).Agent().GetName() == types.BasicTitan2 ||
			(*agt).Agent().GetName() == types.BeastTitan ||
			(*agt).Agent().GetName() == types.ColossalTitan ||
			(*agt).Agent().GetName() == types.ArmoredTitan ||
			(*agt).Agent().GetName() == types.FemaleTitan ||
			(*agt).Agent().GetName() == types.JawTitan {
			titanFound = true
			interestingAgents = append(interestingAgents, agt)
		}
	}
	//println("Interesting agents: ", len(interestingAgents))

	cb.c.attributes.agentAttributes.ResetPerception()

	// Checks first if there are interesting agents to escape
	if len(interestingAgents) != 0 && titanFound {
		_, closestAgentPosition := env.ClosestAgent(interestingAgents, agtPos)

		cb.c.attributes.agentAttributes.SetNextPos(pkg.OppositeDirection(cb.c.attributes.agentAttributes.Pos(), closestAgentPosition))
	} else {
		// If there are no interesting agents, the civilian moves randomly
		var nextPos types.Position

		if rand.Intn(10) < 5 {
			nextPos = types.Position{X: cb.c.attributes.agentAttributes.Pos().X + rand.Intn(5), Y: cb.c.attributes.agentAttributes.Pos().Y + rand.Intn(5)}
		} else {
			nextPos = types.Position{X: cb.c.attributes.agentAttributes.Pos().X - rand.Intn(5), Y: cb.c.attributes.agentAttributes.Pos().Y - rand.Intn(5)}
		}

		println("Agent position: ", cb.c.attributes.agentAttributes.Pos().X, cb.c.attributes.agentAttributes.Pos().Y)
		println("Next position: ", nextPos.X, nextPos.Y)

		cb.c.attributes.agentAttributes.SetNextPos(nextPos)
	}

}

func (cb *CivilianBehavior) Act(e *env.Environment) {
	if env.IsNextPositionValid(cb.c, e) {
		cb.c.Move(cb.c.attributes.agentAttributes.NextPos())
	} else {
		cb.c.Agent().SetNextPos(cb.c.Pos())
	}
}
