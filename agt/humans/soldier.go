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
	s.attributes.agentAttributes.SetCantSeeBehind([]types.ObjectName{types.Wall, types.Field, types.Dungeon, types.BigHouse, types.SmallHouse})
	behavior := &SoldierBehavior{s: s}
	s.SetBehavior(behavior)
	return s
}

// Setter and getter methods for Soldier
func (s *Soldier) SyncChan() chan string { return s.syncChan }

func (s *Soldier) StopCh() chan struct{} { return s.stopCh }

func (s *Soldier) AgtSyncChan() chan int {
	return s.attributes.agentAttributes.SyncChan()
}

func (s *Soldier) Behavior() *env.BehaviorI { return &s.behavior }

func (s *Soldier) SetBehavior(b env.BehaviorI) { s.behavior = b }

// Methods for Soldier
func (s *Soldier) Percept(e *env.Environment /*, wgPercept *sync.WaitGroup*/) {
	//defer wgPercept.Done()
	s.behavior.Percept(e)
}

func (s *Soldier) Deliberate( /*wgDeliberate *sync.WaitGroup*/ ) {
	//defer wgDeliberate.Done()
	s.behavior.Deliberate()
}

func (s *Soldier) Act(e *env.Environment /*, wgAct *sync.WaitGroup*/) {
	//defer wgAct.Done()
	s.mu.Lock()
	defer s.mu.Unlock()
	s.behavior.Act(e)
}

func (s *Soldier) Id() types.Id { return s.attributes.agentAttributes.Id() }

func (s *Soldier) Start(e *env.Environment /*, wgStart *sync.WaitGroup, wgPercept *sync.WaitGroup, wgDeliberate *sync.WaitGroup, wgAct *sync.WaitGroup*/) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	//wgStart.Done()
	//wgStart.Wait()
	go func() {
		println("Soldier Start")
		for {
			//wgPercept.Add(1)
			s.Percept(e /*, wgPercept*/)
			//wgPercept.Wait()

			//wgDeliberate.Add(1)
			s.Deliberate( /*wgDeliberate*/ )
			//wgDeliberate.Wait()

			//wgAct.Add(1)
			s.Act(e /*, wgAct*/)
			//wgAct.Wait()
		}
	}()
}

func (s *Soldier) Move(pos types.Position) { s.attributes.agentAttributes.SetPos(pos) }

func (s *Soldier) Eat() {

}

func (s *Soldier) Sleep() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
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

func (s *Soldier) Attack(agt *env.AgentI) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// TODO : Verif reachable
	// If the percentage is less than the success rate, the attack is successful
	if rand.Float64() < s.AttackSuccess(s.attributes.agentAttributes.Speed(), (*agt).Agent().Speed()) {
		// If the attack is successful, the agent loses HP
		(*agt).Agent().SetHp((*agt).Agent().Hp() - s.attributes.agentAttributes.Strength())
		fmt.Printf("Attack successful from %s : %s lost  %d HP \n", s.Id(), (*agt).Id(), (*agt).Agent().Hp())
	} else {
		fmt.Println("Attack unsuccessful.")
		// If the attack is unsuccessful, nothing happens
	}
}

func (s *Soldier) SetPos(pos types.Position) { s.attributes.agentAttributes.SetPos(pos) }

func (s *Soldier) Pos() types.Position { return s.attributes.agentAttributes.Pos() }

func (s *Soldier) Vision() int { return s.attributes.agentAttributes.Vision() }

func (s *Soldier) Object() obj.Object { return s.attributes.agentAttributes.Object() }

func (s *Soldier) Agent() *env.Agent { return &s.attributes.agentAttributes }

func (s *Soldier) PerceivedObjects() []*obj.Object {
	return s.attributes.agentAttributes.PerceivedObjects()
}

func (s *Soldier) PerceivedAgents() []*env.AgentI {
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

	// Add the perceived agents to the list of perceived agents
	for _, object := range perceivedObjects {
		sb.s.attributes.agentAttributes.AddPerceivedObject(object)
	}
	// Add the perceived agents to the list of perceived agents
	for _, agt := range perceivedAgents {
		sb.s.attributes.agentAttributes.AddPerceivedAgent(agt)
	}
	println("Perceived agents: ", len(sb.s.attributes.agentAttributes.PerceivedAgents()))
	println("Perceived objects: ", len(sb.s.attributes.agentAttributes.PerceivedObjects()))
}

func (sb *SoldierBehavior) Deliberate() {
	println("Soldier Deliberate")

	var interestingAgents []*env.AgentI
	agtPos := sb.s.attributes.agentAttributes.Pos()

	BasicTitansNumber := 0
	SpecialTitanIn := false
	//println("Interesting objects: ", len(interestingObjects))

	for _, agt := range sb.s.attributes.agentAttributes.PerceivedAgents() {
		if (*agt).Agent().GetName() == types.BasicTitan1 ||
			(*agt).Agent().GetName() == types.BasicTitan2 {
			interestingAgents = append(interestingAgents, agt)
			BasicTitansNumber++
		}
		if (*agt).Agent().GetName() == types.BeastTitan ||
			(*agt).Agent().GetName() == types.ColossalTitan ||
			(*agt).Agent().GetName() == types.ArmoredTitan ||
			(*agt).Agent().GetName() == types.FemaleTitan ||
			(*agt).Agent().GetName() == types.JawTitan {
			interestingAgents = append(interestingAgents, agt)
			SpecialTitanIn = true
			// A voir si on veut quand même récupérer les reste des agents
			//break
		}
	}
	//println("Interesting agents: ", len(interestingAgents))

	sb.s.attributes.agentAttributes.ResetPerception()

	// Checks first if there are interesting agents to attack and if not, the nearest agent to go to
	if len(interestingAgents) != 0 && !SpecialTitanIn && BasicTitansNumber < 2 {
		closestAgent, closestAgentPosition := env.ClosestAgent(interestingAgents, agtPos)

		if pkg.DetectCollision((*closestAgent).Object(), sb.s.Object()) {
			sb.s.attributes.agentAttributes.SetAttack(true)
			sb.s.attributes.agentAttributes.SetAgentToAttack(closestAgent)
		} else {
			sb.s.attributes.agentAttributes.SetAttack(false)

			neighborAgentPositions := pkg.GetNeighbors(agtPos, sb.s.attributes.agentAttributes.Speed())
			nextPos := closestAgentPosition.ClosestPosition(neighborAgentPositions)

			sb.s.attributes.agentAttributes.SetNextPos(nextPos)
		}
	} else if SpecialTitanIn || BasicTitansNumber >= 2 {
		//if the soldier is in danger, they run away
		_, closestAgentPosition := env.ClosestAgent(interestingAgents, agtPos)
		sb.s.attributes.agentAttributes.SetNextPos(pkg.OppositeDirection(sb.s.attributes.agentAttributes.Pos(), closestAgentPosition))
	} else {
		// If there are no titans around, the soldier moves randomly
		var nextPos types.Position

		if rand.Intn(10) < 5 {
			nextPos = types.Position{X: sb.s.attributes.agentAttributes.Pos().X + rand.Intn(5), Y: sb.s.attributes.agentAttributes.Pos().Y + rand.Intn(5)}
		} else {
			nextPos = types.Position{X: sb.s.attributes.agentAttributes.Pos().X - rand.Intn(5), Y: sb.s.attributes.agentAttributes.Pos().Y - rand.Intn(5)}
		}

		println("Agent position: ", sb.s.attributes.agentAttributes.Pos().X, sb.s.attributes.agentAttributes.Pos().Y)
		println("Next position: ", nextPos.X, nextPos.Y)

		sb.s.attributes.agentAttributes.SetNextPos(nextPos)
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
		if env.IsNextPositionValid(sb.s, e) {
			sb.s.Move(sb.s.attributes.agentAttributes.NextPos())
		} else {
			sb.s.Agent().SetNextPos(sb.s.Pos())
		}
	}
}
