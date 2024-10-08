package humans

import (
	env "AOT/agt/env"
	obj "AOT/pkg/obj"
	params "AOT/pkg/parameters"
	types "AOT/pkg/types"
	pkg "AOT/pkg/utilitaries"
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
	eren.attributes.agentAttributes.SetCantSeeBehind([]types.ObjectName{types.Wall, types.Field, types.Dungeon, types.BigHouse, types.SmallHouse})
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

func (eren *Eren) AgtSyncChan() chan int {
	return eren.attributes.agentAttributes.SyncChan()
}

func (eren *Eren) Behavior() *env.BehaviorI {
	return &eren.behavior
}

func (eren *Eren) SetBehavior(b env.BehaviorI) {
	eren.behavior = b
}

// Methods for Eren
func (eren *Eren) Percept(e *env.Environment /*, wgPercept *sync.WaitGroup*/) {
	//defer wgPercept.Done()
	eren.behavior.Percept(e)
}

func (eren *Eren) Deliberate( /*wgDeliberate *sync.WaitGroup*/ ) {
	//defer wgDeliberate.Done()
	eren.behavior.Deliberate()
}

func (eren *Eren) Act(e *env.Environment /*, wgAct *sync.WaitGroup*/) {
	//defer wgAct.Done()
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

func (eren *Eren) Start(e *env.Environment /*, wgStart *sync.WaitGroup, wgPercept *sync.WaitGroup, wgDeliberate *sync.WaitGroup, wgAct *sync.WaitGroup*/) {
	println("Start ", eren.Id())
	time.Sleep(100 * time.Millisecond)
	go func() {
		var step int
		for {
			step = <-eren.AgtSyncChan()
			if eren.Agent().ObjectP().Life() > 0 {

				eren.Percept(e)
				eren.Deliberate()
				eren.Act(e)

				time.Sleep(30 * time.Millisecond)
				eren.AgtSyncChan() <- step
			} else {
				time.Sleep(30 * time.Millisecond)
				eren.AgtSyncChan() <- step
			}
		}
	}()
}

func (eren *Eren) Move(pos types.Position) {
	eren.attributes.agentAttributes.SetPos(pos)
}

func (eren *Eren) Eat() {

}

func (eren *Eren) Sleep() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
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

func (eren *Eren) Attack(agt *env.AgentI) {
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < eren.AttackSuccess(eren.attributes.agentAttributes.Speed(), (*agt).Agent().Speed()) {
		// If the attack is successful, the agent loses HP
		(*agt).Agent().SetHp((*agt).Agent().Hp() - eren.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %s : %s lost  %d HP \n", eren.Id(), (*agt).Id(), (*agt).Agent().Hp())
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

func (eren *Eren) PerceivedObjects() []*obj.Object {
	return eren.attributes.agentAttributes.PerceivedObjects()
}

func (eren *Eren) PerceivedAgents() []*env.AgentI {
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

	// Add the perceived agents to the list of perceived agents
	for _, object := range perceivedObjects {
		eb.eren.attributes.agentAttributes.AddPerceivedObject(object)
	}
	// Add the perceived agents to the list of perceived agents
	for _, agt := range perceivedAgents {
		eb.eren.attributes.agentAttributes.AddPerceivedAgent(agt)
	}
	println(eb.eren.Id(), " Perceived agents: ", len(eb.eren.attributes.agentAttributes.PerceivedAgents()))
	println(eb.eren.Id(), " Perceived objects: ", len(eb.eren.attributes.agentAttributes.PerceivedObjects()))

	time.Sleep(100 * time.Millisecond)
}

func (eb *ErenBehavior) Deliberate() {
	println("Eren Deliberate")

	var interestingAgents []*env.AgentI
	agtPos := eb.eren.attributes.agentAttributes.Pos()

	for _, agt := range eb.eren.attributes.agentAttributes.PerceivedAgents() {
		if (*agt).Agent().GetName() == types.BasicTitan1 ||
			(*agt).Agent().GetName() == types.BasicTitan2 ||
			(*agt).Agent().GetName() == types.BeastTitan ||
			(*agt).Agent().GetName() == types.ColossalTitan ||
			(*agt).Agent().GetName() == types.ArmoredTitan ||
			(*agt).Agent().GetName() == types.FemaleTitan ||
			(*agt).Agent().GetName() == types.JawTitan {
			interestingAgents = append(interestingAgents, agt)
			fmt.Println(eb.eren.Id(), " Eren will transform into Eren Titan")
			eb.eren.transform = true
			fmt.Println(eb.eren.Id(), " Eren transform: ", eb.eren.transform)
		}
	}
	println(eb.eren.Id(), " Interesting agents: ", len(interestingAgents))

	eb.eren.attributes.agentAttributes.ResetPerception()

	// Checks first if there are interesting agents to attack and if not, the nearest agent to go to
	if len(interestingAgents) != 0 {
		closestAgent, closestAgentPosition := env.ClosestAgent(interestingAgents, agtPos)

		fmt.Println(eb.eren.Id(), " Eren is interested by ", (*closestAgent).Id())

		if pkg.DetectCollision((*closestAgent).Object(), eb.eren.Object()) {
			eb.eren.attributes.agentAttributes.SetAttack(true)
			eb.eren.attributes.agentAttributes.SetAgentToAttack(closestAgent)
			fmt.Println(eb.eren.Id(), " Eren will attack ", (*closestAgent).Id())
		} else {
			eb.eren.attributes.agentAttributes.SetAttack(false)

			neighborAgentPositions := pkg.GetNeighbors(agtPos, eb.eren.attributes.agentAttributes.Speed())
			nextPos := closestAgentPosition.ClosestPosition(neighborAgentPositions)

			eb.eren.attributes.agentAttributes.SetNextPos(nextPos)
		}
	} else {
		// If there are no interesting agents, Eren moves randomly
		var nextPos types.Position

		if rand.Intn(10) < 5 {
			nextPos = types.Position{X: eb.eren.attributes.agentAttributes.Pos().X + rand.Intn(5), Y: eb.eren.attributes.agentAttributes.Pos().Y + rand.Intn(5)}
		} else {
			nextPos = types.Position{X: eb.eren.attributes.agentAttributes.Pos().X - rand.Intn(5), Y: eb.eren.attributes.agentAttributes.Pos().Y - rand.Intn(5)}
		}

		println(eb.eren.Id(), "Agent position: ", eb.eren.attributes.agentAttributes.Pos().X, eb.eren.attributes.agentAttributes.Pos().Y)
		println(eb.eren.Id(), " Next position: ", nextPos.X, nextPos.Y)

		eb.eren.attributes.agentAttributes.SetNextPos(nextPos)
	}
}

func (eb *ErenBehavior) Act(e *env.Environment) {
	println("Eren Act")
	if eb.eren.transform {
		eb.eren.attributes.agentAttributes.SetName(types.ErenTitanS)
		eb.eren.attributes.agentAttributes.SetSpeed(params.EREN_TITAN_SPEED)
		eb.eren.attributes.agentAttributes.SetStrength(params.EREN_TITAN_STRENGTH)
		eb.eren.attributes.agentAttributes.SetVision(params.EREN_TITAN_VISION)
		eb.eren.attributes.agentAttributes.SetCantSeeBehind([]types.ObjectName{types.Wall})
		fmt.Println("Eren transforms into Eren Titan")
		eb.eren.transform = false
	}
	if eb.eren.attributes.agentAttributes.Attack() {
		eb.eren.attributes.agentAttributes.SetName(types.ErenTitanS)
		eb.eren.attributes.agentAttributes.SetSpeed(params.EREN_TITAN_SPEED)
		eb.eren.attributes.agentAttributes.SetStrength(params.EREN_TITAN_STRENGTH)
		eb.eren.attributes.agentAttributes.SetVision(params.EREN_TITAN_VISION)
		eb.eren.attributes.agentAttributes.SetCantSeeBehind([]types.ObjectName{types.Wall})
		fmt.Println("Eren transforms into Eren Titan")
		fmt.Println("Eren attacks a titan")
		eb.eren.Move(eb.eren.attributes.agentAttributes.NextPos())
		eb.eren.Attack(eb.eren.attributes.agentAttributes.AgentToAttack())
		// Reset the parameters
		eb.eren.attributes.agentAttributes.SetAttack(false)
		eb.eren.attributes.agentAttributes.SetAgentToAttack(nil)
	} else {
		if env.IsNextPositionValid(eb.eren, e) {
			eb.eren.Move(eb.eren.attributes.agentAttributes.NextPos())
		} else {
			eb.eren.Agent().SetNextPos(env.FirstValidPositionToCityCenter(eb.eren, e))
		}
	}
}
