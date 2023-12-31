package humans

import (
	env "AOT/agt/env"
	obj "AOT/pkg/obj"
	types "AOT/pkg/types"
	pkg "AOT/pkg/utilitaries"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MikasaI interface {
	HumanI
}

type Mikasa struct {
	attributes Human
	stopCh     chan struct{}
	syncChan   chan string
	mu         sync.Mutex
	behavior   env.BehaviorI
}

func NewMikasa(id types.Id, tl types.Position, life int, reach int, strength int, speed int, vision int, obj types.ObjectName) *Mikasa {
	atts := NewHuman(id, tl, life, reach, strength, speed, vision, obj)
	m := &Mikasa{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
	}
	behavior := &MikasaBehavior{m: m}
	m.SetBehavior(behavior)
	return m
}

// Setter and getter methods for Mikasa
func (m *Mikasa) SyncChan() chan string {
	return m.syncChan
}

func (m *Mikasa) StopCh() chan struct{} {
	return m.stopCh
}

func (m *Mikasa) Behavior() *env.BehaviorI {
	return &m.behavior
}

func (m *Mikasa) SetBehavior(b env.BehaviorI) {
	m.behavior = b
}

// Methods for Mikasa
func (m *Mikasa) Percept(e *env.Environment, wgPercept *sync.WaitGroup) {
	defer wgPercept.Done()
	m.behavior.Percept(e)
}

func (m *Mikasa) Deliberate(wgDeliberate *sync.WaitGroup) {
	defer wgDeliberate.Done()
	m.behavior.Deliberate()
}

func (m *Mikasa) Act(e *env.Environment, wgAct *sync.WaitGroup) {
	defer wgAct.Done()
	m.mu.Lock()
	defer m.mu.Unlock()
	m.behavior.Act(e)
}

func (m *Mikasa) Id() types.Id {
	return m.attributes.agentAttributes.Id()
}

func (m *Mikasa) Agent() *env.Agent {
	return &m.attributes.agentAttributes
}

func (m *Mikasa) Start(e *env.Environment, wgStart *sync.WaitGroup, wgPercept *sync.WaitGroup, wgDeliberate *sync.WaitGroup, wgAct *sync.WaitGroup) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	wgStart.Done()
	wgStart.Wait()
	go func() {
		println("Mikasa Start")
		for {
			wgPercept.Add(1)
			m.Percept(e, wgPercept)
			wgPercept.Wait()

			wgDeliberate.Add(1)
			m.Deliberate(wgDeliberate)
			wgDeliberate.Wait()

			wgAct.Add(1)
			m.Act(e, wgAct)
			wgAct.Wait()
		}
	}()
}

func (m *Mikasa) Move(pos types.Position) {
	m.attributes.agentAttributes.SetPos(pos)
}

func (m *Mikasa) Eat() {

}

func (m *Mikasa) Sleep() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
}

func (m *Mikasa) Guard() {

}

func (m *Mikasa) AttackSuccess(spdAtk int, spdDef int) float64 {
	// If the speed of the attacker is greater than the speed of the defender, the attack is successful
	if spdAtk > spdDef {
		return 1
	} else {
		// If the speed of the attacker is less than the speed of the defender, the attack is successful with a probability of
		// (speed of the attacker)/(speed of the defender)
		return float64(spdAtk) / float64(spdDef)
	}
}

func (m *Mikasa) Attack(agt env.AgentI) {
	m.mu.Lock()
	defer m.mu.Unlock()
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < m.AttackSuccess(m.attributes.agentAttributes.Speed(), agt.Agent().Speed()) {
		// If the attack is successful, the agent loses HP
		agt.Agent().SetHp(agt.Agent().Hp() - m.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %s : %s lost  %d HP \n", m.Id(), agt.Id(), agt.Agent().Hp())
	} else {
		fmt.Println("Attack unsuccessful.")
		// If the attack is unsuccessful, nothing happens
	}
}

func (m *Mikasa) Pos() types.Position {
	return m.attributes.agentAttributes.Pos()
}

func (m *Mikasa) Vision() int {
	return m.attributes.agentAttributes.Vision()
}

func (m *Mikasa) Object() obj.Object {
	return m.attributes.agentAttributes.Object()
}

func (m *Mikasa) PerceivedObjects() []obj.Object {
	return m.attributes.agentAttributes.PerceivedObjects()
}

func (m *Mikasa) PerceivedAgents() []env.AgentI {
	return m.attributes.agentAttributes.PerceivedAgents()
}

func (m *Mikasa) SetPos(pos types.Position) { m.attributes.agentAttributes.SetPos(pos) }

// Define the behavior struct of Mikasa
type MikasaBehavior struct {
	m *Mikasa
}

func (mb *MikasaBehavior) Percept(e *env.Environment) {
	println("Mikasa Percept")
	// Get the perceived objects and agents
	perceivedObjects, perceivedAgents := mb.m.attributes.agentAttributes.GetVision(e)

	// Add the perceived agents to the list of perceived agents
	for _, object := range perceivedObjects {
		mb.m.attributes.agentAttributes.AddPerceivedObject(object)
	}
	// Add the perceived agents to the list of perceived agents
	for _, agt := range perceivedAgents {
		mb.m.attributes.agentAttributes.AddPerceivedAgent(agt)
	}
	println("Perceived agents: ", len(mb.m.attributes.agentAttributes.PerceivedAgents()))
	println("Perceived objects: ", len(mb.m.attributes.agentAttributes.PerceivedObjects()))

	time.Sleep(100 * time.Millisecond)
}

func (mb *MikasaBehavior) Deliberate() {
	println("Mikasa Deliberate")

	//TODO : Find where to put GetAvoidancePositions function
	// Checks hitbox around to avoid collisions
	toAvoid := []types.Position{}
	for _, object := range mb.m.attributes.agentAttributes.PerceivedObjects() {
		for _, pos := range pkg.GetPositionsInHitbox(object.TL(), object.Hitbox()[1]) {
			toAvoid = append(toAvoid, pos)
		}
	}

	for _, agt := range mb.m.attributes.agentAttributes.PerceivedAgents() {
		for _, pos := range pkg.GetPositionsInHitbox(agt.Agent().ObjectP().TL(), agt.Agent().ObjectP().Hitbox()[1]) {
			toAvoid = append(toAvoid, pos)
			toAvoid = append(toAvoid, types.Position{X: pos.X - mb.m.Agent().ObjectP().Hitbox()[0].X, Y: pos.Y - mb.m.Agent().ObjectP().Hitbox()[0].Y})
		}
	}
	var interestingAgents []env.AgentI
	agtPos := mb.m.attributes.agentAttributes.Pos()

	numberTitans := 0
	ErenAround := false
	//println("Interesting objects: ", len(interestingObjects))

	for _, agt := range mb.m.attributes.agentAttributes.PerceivedAgents() {
		if agt.Agent().GetName() == types.BasicTitan1 ||
			agt.Agent().GetName() == types.BasicTitan2 ||
			agt.Agent().GetName() == types.BeastTitan ||
			agt.Agent().GetName() == types.ColossalTitan ||
			agt.Agent().GetName() == types.ArmoredTitan ||
			agt.Agent().GetName() == types.FemaleTitan ||
			agt.Agent().GetName() == types.JawTitan {
			interestingAgents = append(interestingAgents, agt)
			numberTitans++
		}
		if agt.Agent().GetName() == types.Eren {
			// Mikasa always goes to Eren first
			interestingAgents = append(interestingAgents, agt)
			ErenAround = true
			mb.m.attributes.agentAttributes.SetNextPos(agt.Agent().Pos())
		}
	}
	//println("Interesting agents: ", len(interestingAgents))

	mb.m.attributes.agentAttributes.ResetPerception()

	// If Eren is not around, Mikasa checks first if there are interesting agents to attack and if not, the nearest agent to go to
	if !ErenAround {
		if len(interestingAgents) != 0 {
			closestAgent, closestAgentPosition := env.ClosestAgent(interestingAgents, agtPos)

			if pkg.DetectCollision(closestAgent.Object(), mb.m.Object()) {
				mb.m.attributes.agentAttributes.SetAttack(true)
				mb.m.attributes.agentAttributes.SetAgentToAttack(closestAgent)
			} else {
				mb.m.attributes.agentAttributes.SetAttack(false)

				neighborAgentPositions := pkg.GetNeighbors(agtPos, mb.m.attributes.agentAttributes.Speed(), toAvoid)
				nextPos := closestAgentPosition.ClosestPosition(neighborAgentPositions)

				mb.m.attributes.agentAttributes.SetNextPos(nextPos)
			}
		} else {
			// If there are no interesting agents, Mikasa moves randomly
			var nextPos types.Position

			if rand.Intn(10) < 5 {
				nextPos = types.Position{X: mb.m.attributes.agentAttributes.Pos().X + rand.Intn(5), Y: mb.m.attributes.agentAttributes.Pos().Y + rand.Intn(5)}
			} else {
				nextPos = types.Position{X: mb.m.attributes.agentAttributes.Pos().X - rand.Intn(5), Y: mb.m.attributes.agentAttributes.Pos().Y - rand.Intn(5)}
			}

			println("Agent position: ", mb.m.attributes.agentAttributes.Pos().X, mb.m.attributes.agentAttributes.Pos().Y)
			println("Next position: ", nextPos.X, nextPos.Y)

			mb.m.attributes.agentAttributes.SetNextPos(nextPos)
		}
	}
}

func (mb *MikasaBehavior) Act(e *env.Environment) {
	// Perform the action based on the parameters
	if mb.m.attributes.agentAttributes.Attack() {
		mb.m.Move(mb.m.attributes.agentAttributes.NextPos())
		mb.m.Attack(mb.m.attributes.agentAttributes.AgentToAttack())
		// Reset the parameters
		mb.m.attributes.agentAttributes.SetAttack(false)
		mb.m.attributes.agentAttributes.SetAgentToAttack(nil)
	} else {
		// Move towards the specified position
		mb.m.Move(mb.m.attributes.agentAttributes.NextPos())
	}
}
