package env

import (
	obj "AOT/pkg/obj"
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"
)

type Agent struct {
	id               types.Id
	reach            int
	strength         int
	speed            int
	maxHP            int
	vision           int
	attack           bool
	syncChan         chan int
	nextPosition     types.Position
	agentToAttack    *AgentI
	perceivedObjects []*obj.Object
	perceivedAgents  []*AgentI
	cantSeeBehind    []types.ObjectName
	object           *obj.Object
}

func NewAgent(id types.Id, tl types.Position, life int, reach int, strength int, speed int, vision int, name types.ObjectName) *Agent {
	return &Agent{
		id:               id,
		reach:            reach,
		strength:         strength,
		speed:            speed,
		vision:           vision,
		attack:           false,
		syncChan:         make(chan int),
		nextPosition:     tl,
		agentToAttack:    nil,
		perceivedObjects: []*obj.Object{},
		perceivedAgents:  []*AgentI{},
		cantSeeBehind:    []types.ObjectName{},
		object:           obj.NewObject(name, tl, life),
	}
}

func (t *Agent) SyncChan() chan int { return t.syncChan }

func (t *Agent) Id() types.Id {
	return t.id
}

func (t *Agent) Pos() types.Position { return t.object.TL() }

func (t *Agent) Reach() int {
	return t.reach
}

func (t *Agent) Strength() int {
	return t.strength
}

func (t *Agent) Speed() int {
	return t.speed
}

func (t *Agent) MaxHP() int {
	return t.maxHP
}

func (t *Agent) Hp() int { return t.object.Life() }

func (t *Agent) SetPos(pos types.Position) { t.object.SetPosition(pos) }

func (t *Agent) SetAttack(b bool) { t.attack = b }

func (t *Agent) SetNextPos(nextPos types.Position) { t.nextPosition = nextPos }

func (t *Agent) SetAgentToAttack(agt *AgentI) { t.agentToAttack = agt }

func (t *Agent) SetHp(hp int) { t.object.SetLife(hp) }

func (t *Agent) Vision() int { return t.vision }

func (t *Agent) NextPos() types.Position { return t.nextPosition }

func (t *Agent) Attack() bool { return t.attack }

func (t *Agent) AgentToAttack() *AgentI { return t.agentToAttack }

func (t *Agent) PerceivedObjects() []*obj.Object { return t.perceivedObjects }

func (t *Agent) PerceivedAgents() []*AgentI { return t.perceivedAgents }

func (t *Agent) CantSeeBehind() []types.ObjectName { return t.cantSeeBehind }

func (t *Agent) SetCantSeeBehind(cantSeeBehindObjects []types.ObjectName) {
	for _, objects := range cantSeeBehindObjects {
		t.cantSeeBehind = append(t.cantSeeBehind, objects)
	}
}

func (t *Agent) Object() obj.Object { return *t.object }

func (t *Agent) ObjectP() *obj.Object { return t.object }

func (t *Agent) AddPerceivedObject(obj *obj.Object) {
	t.perceivedObjects = append(t.perceivedObjects, obj)
}

func (t *Agent) AddPerceivedAgent(agt *AgentI) { t.perceivedAgents = append(t.perceivedAgents, agt) }

func (t *Agent) NextPosition() types.Position { return t.nextPosition }

func (t *Agent) AttackValue() bool { return t.attack }

func (t *Agent) GetName() types.ObjectName { return t.object.Name() }

func (t *Agent) SetName(name types.ObjectName) { t.object.SetName(name) }

func (t *Agent) ResetPerception() {
	//println("Reset perception")
	t.perceivedObjects = []*obj.Object{}
	t.perceivedAgents = []*AgentI{}
}

func (t *Agent) SetSpeed(speed int) { t.speed = speed }

func (t *Agent) SetStrength(strength int) { t.strength = strength }

func (t *Agent) SetVision(vision int) { t.vision = vision }

// returns a list of objects that the agent can see
// the vision is a square centered on the agent position
func (t *Agent) GetVision(e *Environment) ([]*obj.Object, []*AgentI) {

	// Get the top left and bottom right positions of the vision square
	agtCenter := t.ObjectP().Center()
	agtVision := t.Vision()
	topLeft := types.Position{X: agtCenter.X - agtVision, Y: agtCenter.Y - agtVision}
	bottomRight := types.Position{X: agtCenter.X + agtVision, Y: agtCenter.Y + agtVision}

	// Get the positions inside the vision square from the environment
	perceivedObjects := e.PerceivedObjects(topLeft, bottomRight)
	//fmt.Println("In vision before : env perceivedObjects:", perceivedObjects)
	perceivedAgents := e.PerceivedAgents(topLeft, bottomRight, t.id)
	//fmt.Println(t.Id(), "In vision before : env perceivedAgents:", len(perceivedAgents))

	// Get the objects not seen by the agent
	CantSeeBehindObjects := []obj.Object{}
	for i, obj := range perceivedObjects {
		if utils.Contains(t.CantSeeBehind(), obj.Name()) {
			CantSeeBehindObjects = append(CantSeeBehindObjects, *perceivedObjects[i])
		}
	}
	// Filter out positions to avoid regarding the agent position
	// If a CantSeeBehindObjects object is in the vision square, agent can't see behind it,
	// depending on the angle between the agent and the position to avoid, the agent can't see every positions behind it following the perspective logic

	// Get the positions behind the CantSeeBehindObjects objects
	noSeeableSquaresBehindObjects := map[*obj.Object][]types.Position{}

	for i, object := range perceivedObjects {
		if !utils.Contains(CantSeeBehindObjects, *perceivedObjects[i]) {
			continue
		} else {
			// If a objCantSeeBehindObject is in the vision square, the agent can't see behind it
			// Get the angle between the agent and the position to avoid
			objectCenter := object.Center()
			angle := utils.GetAngle(t.ObjectP().Center(), objectCenter)

			// Get the perceivedObjects behind the current position to avoid
			noSeeableSquaresBehindObjects[object] = utils.GetNotSeeableBoxBehindObject(t.Object(), angle, topLeft, bottomRight)
			//fmt.Println(t.Id(), " Can't see in :", noSeeableSquaresBehindObjects[object], "because of ", object.Name(), "at", objectCenter, "with angle", angle)

		}
	}
	//fmt.Println(t.Id(), " Position ", t.Pos())
	//fmt.Println(t.Id(), " Can't see behind:", CantSeeBehindObjects)
	//fmt.Println(t.Id(), " Can't see this number of positions:", len(positionsBehindObjects))
	//fmt.Println(t.Id(), " at these positions:", positionsBehindObjects)

	// Checks if the perceivedObjects are in a positionsBehindObjects position and remove them if they are
	perceivedObjects = utils.RemoveNoSeeableObjects(perceivedObjects, noSeeableSquaresBehindObjects)
	perceivedAgents = RemoveNoSeeableAgents(perceivedAgents, noSeeableSquaresBehindObjects)

	return perceivedObjects, perceivedAgents
}

func ClosestAgent(agents []*AgentI, position types.Position) (*AgentI, types.Position) {
	// Get the closest position from the list
	closestAgent := agents[0]
	closestAgentPosition := (*agents[0]).Agent().ObjectP().Hitbox()[0]
	for i, agt := range agents {
		for _, pos := range utils.GetPositionsInHitbox((*agt).Agent().ObjectP().Hitbox()[0], (*agt).Agent().ObjectP().Hitbox()[1]) {
			if position.Distance(pos) < position.Distance(closestAgentPosition) {
				closestAgent = agents[i]
				closestAgentPosition = pos
			}
		}
	}
	return closestAgent, closestAgentPosition
}
