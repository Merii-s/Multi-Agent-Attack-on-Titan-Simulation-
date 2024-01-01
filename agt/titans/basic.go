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
	wallToGo          types.Position
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

func (bt *BasicTitan) SetBehavior(b env.BehaviorI) { bt.behavior = b }

func (bt *BasicTitan) SetWallToGo(pos types.Position) { bt.wallToGo = pos }

func (bt *BasicTitan) WallToGo() types.Position { return bt.wallToGo }

// Methods for BasicTitan
func (bt *BasicTitan) Percept(e *env.Environment, wgPercept *sync.WaitGroup) {
	defer wgPercept.Done()
	bt.behavior.Percept(e)
}

func (bt *BasicTitan) Deliberate(wgDeliberate *sync.WaitGroup) {
	defer wgDeliberate.Done()
	bt.behavior.Deliberate()

}

func (bt *BasicTitan) Act(e *env.Environment, wgAct *sync.WaitGroup) {
	defer wgAct.Done()
	bt.mu.Lock()
	defer bt.mu.Unlock()
	bt.behavior.Act(e)
}

func (bt *BasicTitan) Start(e *env.Environment, wgStart *sync.WaitGroup, wgPercept *sync.WaitGroup, wgDeliberate *sync.WaitGroup, wgAct *sync.WaitGroup) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	wgStart.Done()
	wgStart.Wait()
	go func() {
		println("BasicTitan Start")
		for {
			wgPercept.Add(1)
			bt.Percept(e, wgPercept)
			wgPercept.Wait()

			wgDeliberate.Add(1)
			bt.Deliberate(wgDeliberate)
			wgDeliberate.Wait()

			wgAct.Add(1)
			bt.Act(e, wgAct)
			wgAct.Wait()
			time.Sleep(100 * time.Millisecond)
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

func (bt *BasicTitan) Attack(agt env.AgentI) {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	// TO
	// TODO : Verif reachable
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
	// If the titan is out of the screen, it sets the closest wall position as the wall to go
	if pkg.IsOutOfScreen(btb.bt.attributes.agentAttributes.Pos(), params.ScreenWidth, params.ScreenHeight) && !btb.bt.firstWallInitDone {
		println("BasicTitan is out of screen")
		// Checks hitbox around to avoid collisions
		btb.bt.firstWallInitDone = true
		wallPositions := env.GetWallPositions(e)

		//println("BasicTitan Position: ", btb.bt.attributes.agentAttributes.Pos().X, btb.bt.attributes.agentAttributes.Pos().Y)
		btb.bt.SetWallToGo(btb.bt.attributes.agentAttributes.Pos().ClosestPosition(wallPositions))
		println("Wall to go: ", btb.bt.WallToGo().X, btb.bt.WallToGo().Y)
	}
	println("BasicTitan Position: ", btb.bt.attributes.agentAttributes.Pos().X, btb.bt.attributes.agentAttributes.Pos().Y)

	// Get the perceived objects and agents
	perceivedObjects, perceivedAgents := btb.bt.attributes.agentAttributes.GetVision(e)

	// Add the perceived agents to the list of perceived agents
	for i, _ := range perceivedObjects {
		btb.bt.attributes.agentAttributes.AddPerceivedObject(perceivedObjects[i])
	}
	// Add the perceived agents to the list of perceived agents
	for i, _ := range perceivedAgents {
		btb.bt.attributes.agentAttributes.AddPerceivedAgent(perceivedAgents[i])
	}
	println("Perceived agents: ", len(btb.bt.attributes.agentAttributes.PerceivedAgents()))
	println("Perceived objects: ", len(btb.bt.attributes.agentAttributes.PerceivedObjects()))

}

func (btb *BasicTitanBehavior) Deliberate() {
	println("BasicTitan Deliberate")

	//TODO : Find where to put GetAvoidancePositions function
	// Checks hitbox around to avoid collisions
	toAvoid := []types.Position{}
	for _, object := range btb.bt.attributes.agentAttributes.PerceivedObjects() {
		for _, pos := range pkg.GetPositionsInHitbox(object.TL(), object.Hitbox()[1]) {
			toAvoid = append(toAvoid, pos)
		}
	}

	for _, agt := range btb.bt.attributes.agentAttributes.PerceivedAgents() {
		for _, pos := range pkg.GetPositionsInHitbox(agt.Agent().ObjectP().TL(), agt.Agent().ObjectP().Hitbox()[1]) {
			toAvoid = append(toAvoid, pos)
			toAvoid = append(toAvoid, types.Position{X: pos.X - btb.bt.Agent().ObjectP().Hitbox()[0].X, Y: pos.Y - btb.bt.Agent().ObjectP().Hitbox()[0].Y})
		}
	}
	if pkg.IsOutOfScreen(btb.bt.attributes.agentAttributes.Pos(), params.ScreenWidth, params.ScreenHeight) {
		println("BasicTitan is out of screen")
		neighborAgentPositions := pkg.GetNeighbors(btb.bt.attributes.agentAttributes.Pos(), btb.bt.attributes.agentAttributes.Speed(), toAvoid)
		nextPos := btb.bt.wallToGo.ClosestPosition(neighborAgentPositions)
		println("Next position: ", nextPos.X, nextPos.Y)
		btb.bt.attributes.agentAttributes.SetNextPos(nextPos)

	} else {
		var interestingObjects []obj.Object
		var interestingAgents []env.AgentI
		agtPos := btb.bt.attributes.agentAttributes.Pos()

		for i, object := range btb.bt.attributes.agentAttributes.PerceivedObjects() {
			if object.Name() == types.Wall || object.Name() == types.Field {
				interestingObjects = append(interestingObjects, btb.bt.attributes.agentAttributes.PerceivedObjects()[i])
			}
		}
		//println("Interesting objects: ", len(interestingObjects))

		for i, agt := range btb.bt.attributes.agentAttributes.PerceivedAgents() {
			if agt.Agent().GetName() == types.MaleCivilian ||
				agt.Agent().GetName() == types.FemaleCivilian ||
				agt.Agent().GetName() == types.Eren ||
				agt.Agent().GetName() == types.Mikasa ||
				agt.Agent().GetName() == types.MaleSoldier ||
				agt.Agent().GetName() == types.FemaleSoldier {
				interestingAgents = append(interestingAgents, btb.bt.attributes.agentAttributes.PerceivedAgents()[i])
			}
		}
		//println("Interesting agents: ", len(interestingAgents))

		btb.bt.attributes.agentAttributes.ResetPerception()

		// Checks first if there are interesting agents to attack and if not, the nearest agent to go to
		if len(interestingAgents) != 0 {
			closestAgent, closestAgentPosition := env.ClosestAgent(interestingAgents, agtPos)

			if pkg.DetectCollision(closestAgent.Object(), btb.bt.Object()) {
				btb.bt.attributes.agentAttributes.SetAttack(true)
				btb.bt.attributes.SetAttackObject(false)
				btb.bt.attributes.agentAttributes.SetAgentToAttack(closestAgent)
			} else {
				btb.bt.attributes.agentAttributes.SetAttack(false)
				btb.bt.attributes.SetAttackObject(false)

				neighborAgentPositions := pkg.GetNeighbors(agtPos, btb.bt.attributes.agentAttributes.Speed(), toAvoid)
				nextPos := closestAgentPosition.ClosestPosition(neighborAgentPositions)

				btb.bt.attributes.agentAttributes.SetNextPos(nextPos)
			}

			// If there are no interesting agents, the titan goes towards the nearest interesting object (wall or field)
		} else if len(interestingObjects) != 0 {
			closestObject, closestObjectPosition := pkg.ClosestObject(interestingObjects, agtPos)
			println("Closest object: ", closestObject.Name())
			println("Closest object Position: ", closestObjectPosition.X, closestObjectPosition.Y)

			if pkg.DetectCollision(closestObject, btb.bt.Object()) {
				//println("Attack object: ", btb.bt.attributes.ObjectToAttack().GetName())
				btb.bt.attributes.agentAttributes.SetAttack(false)
				btb.bt.attributes.SetAttackObject(true)
				btb.bt.attributes.SetObjectToAttack(&closestObject)
				println("Attack ", btb.bt.attributes.agentAttributes.Attack())
				println("Attack Object", btb.bt.attributes.AttackObjectBool())
				println("Object to attack: ", btb.bt.attributes.ObjectToAttack().GetName())
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
	println("BasicTitan Act")
	// If the titan is attacking an agent, it attacks it
	if btb.bt.attributes.agentAttributes.Attack() {
		println("Attack")
		btb.bt.Attack(btb.bt.attributes.agentAttributes.AgentToAttack())
		btb.bt.attributes.agentAttributes.SetAttack(false)
		btb.bt.attributes.agentAttributes.SetAgentToAttack(nil)
	}
	// If the titan is attacking an object, it attacks it
	if btb.bt.attributes.AttackObjectBool() {
		println("Attack object")
		println("Wall life before attack: ", btb.bt.attributes.GetObjectToAttackLife())
		btb.bt.attributes.AttackObject(btb.bt.attributes.ObjectToAttackP())
		println("Wall life after attack: ", btb.bt.attributes.GetObjectToAttackLife())
		btb.bt.attributes.SetAttackObject(false)
		btb.bt.attributes.SetObjectToAttack(nil)
	}
	// If the titan is not attacking anything, it moves towards the next position
	if !btb.bt.attributes.AttackObjectBool() && !btb.bt.attributes.agentAttributes.Attack() {
		btb.bt.Move(btb.bt.attributes.agentAttributes.NextPosition())
	}
}
