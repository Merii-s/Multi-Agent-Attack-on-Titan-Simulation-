package titans

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

type BasicTitanI interface {
	TitanI
}

type BasicTitan struct {
	attributes        Titan
	stopCh            chan struct{}
	syncChan          chan string
	mu                sync.Mutex
	behavior          env.BehaviorI
	firstWallInitDone bool
	wallToGo          *obj.Object
	wallToGoPos       types.Position
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
	bt.attributes.agentAttributes.SetCantSeeBehind([]types.ObjectName{types.Wall})
	behavior := &BasicTitanBehavior{bt: bt}
	bt.SetBehavior(behavior)
	return bt
}

// Setter and getter methods for BasicTitan
func (bt *BasicTitan) SyncChan() chan int {
	return bt.attributes.agentAttributes.SyncChan()
}

func (bt *BasicTitan) StopCh() chan struct{} {
	return bt.stopCh
}

func (bt *BasicTitan) AgtSyncChan() chan int {
	return bt.attributes.agentAttributes.SyncChan()
}

func (bt *BasicTitan) Behavior() *env.BehaviorI {
	return &bt.behavior
}

func (bt *BasicTitan) SetBehavior(b env.BehaviorI) { bt.behavior = b }

func (bt *BasicTitan) SetWallToGo(wall *obj.Object) { bt.wallToGo = wall }

func (bt *BasicTitan) WallToGo() *obj.Object { return bt.wallToGo }

func (bt *BasicTitan) WallToGoPos() types.Position { return bt.wallToGoPos }

func (bt *BasicTitan) SetWallToGoPos(pos types.Position) { bt.wallToGoPos = pos }

// Methods for BasicTitan
func (bt *BasicTitan) Percept(e *env.Environment /*, wgPercept *sync.WaitGroup*/) {
	//defer wgPercept.Done()
	bt.behavior.Percept(e)
	fmt.Println("Percept Done  ", bt.Id())
}

func (bt *BasicTitan) Deliberate( /*wgDeliberate *sync.WaitGroup*/ ) {
	//defer wgDeliberate.Done()
	bt.behavior.Deliberate()

}

func (bt *BasicTitan) Act(e *env.Environment /*, wgAct *sync.WaitGroup*/) {
	//defer wgAct.Done()
	bt.mu.Lock()
	defer bt.mu.Unlock()
	bt.behavior.Act(e)
	//fmt.Println("Act Done  ", bt.Id())
}

func (bt *BasicTitan) Start(e *env.Environment /*, wgStart *sync.WaitGroup, wgPercept *sync.WaitGroup, wgDeliberate *sync.WaitGroup, wgAct *sync.WaitGroup*/) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	//wgStart.Done()
	println("Start ", bt.Id())
	//wgStart.Wait()
	time.Sleep(100 * time.Millisecond)
	// go func() {
	// 	for {
	// 		//wgPercept.Add(1)
	// 		fmt.Println("Percept ", bt.Id())
	// 		bt.Percept(e /*, wgPercept*/)
	// 		//wgPercept.Done()
	// 		//wgPercept.Wait()
	// 		time.Sleep(15 * time.Millisecond)

	// 		//wgDeliberate.Add(1)
	// 		fmt.Println("Deliberate ", bt.Id())
	// 		//bt.Deliberate(wgDeliberate)
	// 		//wgDeliberate.Wait()
	// 		time.Sleep(15 * time.Millisecond)

	// 		//wgAct.Add(1)
	// 		fmt.Println("Act ", bt.Id())
	// 		//bt.Act(e, wgAct)
	// 		//wgAct.Wait()
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()
	go func() {
		var step int
		for {
			step = <-bt.AgtSyncChan()

			bt.Percept(e)
			bt.Deliberate()
			bt.Act(e)

			time.Sleep(100 * time.Millisecond)
			bt.AgtSyncChan() <- step
		}
	}()

}

func (bt *BasicTitan) Id() types.Id {
	return bt.attributes.agentAttributes.Id()
}

func (bt *BasicTitan) Move(pos types.Position) { bt.attributes.agentAttributes.SetPos(pos) }

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

func (bt *BasicTitan) Attack(agt *env.AgentI) {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	if pkg.DetectCollision((*agt).Object(), bt.Object()) {
		// If the percentage is less than the success rate, the attack is successful
		if rand.Float64() < bt.AttackSuccess(bt.attributes.agentAttributes.Speed(), (*agt).Agent().Speed()) {
			// If the attack is successful, the agent loses HP
			(*agt).Agent().SetHp((*agt).Agent().Hp() - bt.attributes.agentAttributes.Strength())
			fmt.Printf("Attack successful from %s : %s has now %d HP \n", bt.Id(), (*agt).Id(), (*agt).Agent().Hp())
		} else {
			fmt.Println("Attack unsuccessful.")
			// If the attack is unsuccessful, nothing happens
		}
	} else {
		fmt.Println(bt.Id(), " : Attack unsuccessful, not valid anymore.")
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

func (bt *BasicTitan) PerceivedObjects() []*obj.Object {
	return bt.attributes.agentAttributes.PerceivedObjects()
}

func (bt *BasicTitan) PerceivedAgents() []*env.AgentI {
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
	//println("BasicTitan Percept")
	// If the titan is out of the screen, it sets the closest wall position as the wall to go
	if pkg.IsOutOfWalls(btb.bt.attributes.agentAttributes.Pos()) && !btb.bt.firstWallInitDone {
		//println("BasicTitan is out of walls")
		// Checks hitbox around to avoid collisions
		btb.bt.firstWallInitDone = true

		walls := env.GetWallPositions(e)

		closestPosition := types.Position{X: -2 * params.ScreenWidth, Y: -2 * params.ScreenHeight}
		closestWall := &obj.Object{}

		for wall, posTab := range walls {
			agtCenter := btb.bt.Agent().ObjectP().Center()
			closestInCurrentWall := agtCenter.ClosestPosition(posTab)
			if agtCenter.Distance(closestInCurrentWall) < agtCenter.Distance(closestPosition) {
				closestPosition = closestInCurrentWall
				closestWall = wall
			}
		}

		btb.bt.SetWallToGoPos(closestPosition)
		fmt.Println("Wall to go position: ", btb.bt.WallToGoPos().X, btb.bt.WallToGoPos().Y)
		btb.bt.SetWallToGo(closestWall)
		fmt.Println("Wall to go: ", btb.bt.WallToGo().TL())
	}
	println("BasicTitan", btb.bt.Id(), "Position: ", btb.bt.attributes.agentAttributes.Pos().X, btb.bt.attributes.agentAttributes.Pos().Y)

	// Get the perceived objects and agents
	perceivedObjects, perceivedAgents := btb.bt.attributes.agentAttributes.GetVision(e)

	// Add the perceived agents to the list of perceived agents
	for i := range perceivedObjects {
		btb.bt.attributes.agentAttributes.AddPerceivedObject(perceivedObjects[i])
	}
	// Add the perceived agents to the list of perceived agents
	for i := range perceivedAgents {
		btb.bt.attributes.agentAttributes.AddPerceivedAgent(perceivedAgents[i])
	}
	//println("Perceived agents: ", len(btb.bt.attributes.agentAttributes.PerceivedAgents()))
	//println("Perceived objects: ", len(btb.bt.attributes.agentAttributes.PerceivedObjects()))

}

func (btb *BasicTitanBehavior) Deliberate() {
	//println("BasicTitan Deliberate")

	//TODO : Find where to put GetAvoidancePositions function
	// Checks hitbox around to avoid collisions
	toAvoid := []types.Position{}
	for _, object := range btb.bt.attributes.agentAttributes.PerceivedObjects() {
		for _, pos := range pkg.GetPositionsInHitbox(object.TL(), object.Hitbox()[1]) {
			toAvoid = append(toAvoid, pos)
		}
	}

	for _, agt := range btb.bt.attributes.agentAttributes.PerceivedAgents() {
		for _, pos := range pkg.GetPositionsInHitbox((*agt).Agent().ObjectP().TL(), (*agt).Agent().ObjectP().Hitbox()[1]) {
			toAvoid = append(toAvoid, pos)
			toAvoid = append(toAvoid, types.Position{X: pos.X - btb.bt.Agent().ObjectP().Hitbox()[0].X, Y: pos.Y - btb.bt.Agent().ObjectP().Hitbox()[0].Y})
		}
	}

	if pkg.IsOutOfWalls(btb.bt.attributes.agentAttributes.Pos()) && !pkg.DetectCollision(*btb.bt.wallToGo, btb.bt.Object()) {
		neighborAgentPositions := pkg.GetNeighbors(btb.bt.attributes.agentAttributes.Pos(), btb.bt.attributes.agentAttributes.Speed(), toAvoid)
		nextPos := btb.bt.wallToGo.Center().ClosestPosition(neighborAgentPositions)
		println("BasicTitan", btb.bt.Id(), "Next position: ", nextPos.X, nextPos.Y)
		btb.bt.attributes.agentAttributes.SetNextPos(nextPos)
	} else {
		var interestingObjects []*obj.Object
		var interestingAgents []*env.AgentI
		agtPos := btb.bt.attributes.agentAttributes.Pos()

		for i, object := range btb.bt.attributes.agentAttributes.PerceivedObjects() {
			fmt.Println("Perceived object: ", btb.bt.attributes.agentAttributes.PerceivedObjects()[i].Name())
			if object.Name() == types.Wall || object.Name() == types.Field || object.Name() == types.Dungeon {
				fmt.Println("Interesting object: ", btb.bt.attributes.agentAttributes.PerceivedObjects()[i].Name())
				interestingObjects = append(interestingObjects, btb.bt.attributes.agentAttributes.PerceivedObjects()[i])
			}
		}
		//println("Interesting objects: ", len(interestingObjects))

		for i, agt := range btb.bt.attributes.agentAttributes.PerceivedAgents() {
			if (*agt).Agent().GetName() == types.MaleCivilian ||
				(*agt).Agent().GetName() == types.FemaleCivilian ||
				(*agt).Agent().GetName() == types.Eren ||
				(*agt).Agent().GetName() == types.Mikasa ||
				(*agt).Agent().GetName() == types.MaleSoldier ||
				(*agt).Agent().GetName() == types.FemaleSoldier {
				interestingAgents = append(interestingAgents, btb.bt.attributes.agentAttributes.PerceivedAgents()[i])
			}
		}
		//println("Interesting agents: ", len(interestingAgents))

		btb.bt.attributes.agentAttributes.ResetPerception()

		// Checks first if there are interesting agents to attack and if not, the nearest agent to go to
		if len(interestingAgents) != 0 {
			closestAgent, closestAgentPosition := env.ClosestAgent(interestingAgents, agtPos)

			fmt.Println("Closest agent: ", (*closestAgent).Agent().GetName(), " at ", closestAgentPosition.X)

			if pkg.DetectCollision((*closestAgent).Object(), btb.bt.Object()) {
				btb.bt.attributes.agentAttributes.SetAttack(true)
				btb.bt.attributes.SetAttackObject(false)
				btb.bt.attributes.agentAttributes.SetAgentToAttack(closestAgent)
				println("BasicTitan", btb.bt.Id(), " : Attacks agent: ", (*btb.bt.attributes.agentAttributes.AgentToAttack()).Object().GetName())
			} else {
				btb.bt.attributes.agentAttributes.SetAttack(false)
				btb.bt.attributes.SetAttackObject(false)

				neighborAgentPositions := pkg.GetNeighbors(agtPos, btb.bt.attributes.agentAttributes.Speed(), toAvoid)
				nextPos := closestAgentPosition.ClosestPosition(neighborAgentPositions)

				btb.bt.attributes.agentAttributes.SetNextPos(nextPos)
				println("Agent position: ", btb.bt.attributes.agentAttributes.Pos().X, btb.bt.attributes.agentAttributes.Pos().Y)
				println("Next position : ", nextPos.X, nextPos.Y)
			}

			// If there are no interesting agents, the titan goes towards the nearest interesting object (wall or field)
		} else if len(interestingObjects) != 0 {
			closestObject, closestObjectPosition := pkg.ClosestObject(interestingObjects, agtPos)
			println("Closest object: ", closestObject.Name())
			println("Closest object Position: ", closestObjectPosition.X, closestObjectPosition.Y)

			if pkg.DetectCollision(*closestObject, btb.bt.Object()) {
				btb.bt.attributes.agentAttributes.SetAttack(false)
				btb.bt.attributes.SetAttackObject(true)
				btb.bt.attributes.SetObjectToAttack(closestObject)
				println("BasicTitan", btb.bt.Id(), " : Attacks object: ", btb.bt.attributes.ObjectToAttack().GetName())
				//println("Attack ", btb.bt.attributes.agentAttributes.Attack())
				//println("Attack Object", btb.bt.attributes.AttackObjectBool())
				//println("Object to attack: ", btb.bt.attributes.ObjectToAttack().GetName())
			} else {
				btb.bt.attributes.agentAttributes.SetAttack(false)
				btb.bt.attributes.SetAttackObject(false)

				neighborAgentPositions := pkg.GetNeighbors(agtPos, btb.bt.attributes.agentAttributes.Speed(), toAvoid)
				nextPos := closestObjectPosition.ClosestPosition(neighborAgentPositions)
				println("Agent position: ", btb.bt.attributes.agentAttributes.Pos().X, btb.bt.attributes.agentAttributes.Pos().Y)
				println("Next position: ", nextPos.X, nextPos.Y)

				btb.bt.attributes.agentAttributes.SetNextPos(nextPos)
			}
		}
	}
}

func (btb *BasicTitanBehavior) Act(e *env.Environment) {
	//println("BasicTitan Act")
	// If the titan is attacking an agent, it attacks it
	if btb.bt.attributes.agentAttributes.Attack() {
		//println("Attack")
		btb.bt.Attack(btb.bt.attributes.agentAttributes.AgentToAttack())
		btb.bt.attributes.agentAttributes.SetAttack(false)
		btb.bt.attributes.agentAttributes.SetAgentToAttack(nil)

		// If the titan is attacking an object, it attacks it
	} else if btb.bt.attributes.AttackObjectBool() {
		//println("Attack object")
		btb.bt.attributes.AttackObject(btb.bt.attributes.ObjectToAttackP())
		btb.bt.attributes.SetAttackObject(false)
		btb.bt.attributes.SetObjectToAttack(nil)

		// If the titan is not attacking anything, it moves towards the next position
	} else if !btb.bt.attributes.AttackObjectBool() && !btb.bt.attributes.agentAttributes.Attack() {
		if env.IsNextPositionValid(btb.bt, e) {
			btb.bt.Move(btb.bt.attributes.agentAttributes.NextPos())
			fmt.Println("Basic Titan", btb.bt.Id(), " : Can move to next position", btb.bt.attributes.agentAttributes.NextPos())
		} else {
			btb.bt.Agent().SetNextPos(btb.bt.Pos())
			fmt.Println("Basic Titan", btb.bt.Id(), " : Next position is not valid")
		}
	}
}
