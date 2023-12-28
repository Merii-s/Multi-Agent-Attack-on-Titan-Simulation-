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
	nextPosition     types.Position
	agentToAttack    AgentI
	perceivedObjects []obj.Object
	perceivedAgents  []AgentI
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
		nextPosition:     tl,
		agentToAttack:    nil,
		perceivedObjects: []obj.Object{},
		perceivedAgents:  []AgentI{},
		cantSeeBehind:    []types.ObjectName{},
		object:           obj.NewObject(name, tl, life),
	}
}

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

func (t *Agent) SetAgentToAttack(agt AgentI) { t.agentToAttack = agt }

func (t *Agent) SetHp(hp int) { t.object.SetLife(hp) }

func (t *Agent) Vision() int { return t.vision }

func (t *Agent) NextPos() types.Position { return t.nextPosition }

func (t *Agent) Attack() bool { return t.attack }

func (t *Agent) AgentToAttack() AgentI { return t.agentToAttack }

func (t *Agent) PerceivedObjects() []obj.Object { return t.perceivedObjects }

func (t *Agent) PerceivedAgents() []AgentI { return t.perceivedAgents }

func (t *Agent) CantSeeBehind() []types.ObjectName { return t.cantSeeBehind }

func (t *Agent) Object() obj.Object { return *t.object }

func (t *Agent) AddPerceivedObject(obj obj.Object) {
	t.perceivedObjects = append(t.perceivedObjects, obj)
}

func (t *Agent) AddPerceivedAgent(agt AgentI) { t.perceivedAgents = append(t.perceivedAgents, agt) }

func (t *Agent) SetNextPosition(pos Position) { t.nextPosition = pos }

func (t *Agent) NextPosition() Position { return t.nextPosition }

func (t *Agent) AttackValue() bool { return t.attack }

func (t *Agent) GetName() ObjectName { return t.object.name }

// returns a list of objects that the agent can see
// the vision is a square centered on the agent position
func (t *Agent) GetVision(e *Environment) ([]obj.Object, []AgentI) {
	// Get the top left and bottom right positions of the vision square
	topLeft := types.Position{X: t.Pos().X - t.Vision(), Y: t.Pos().Y + t.Vision()}
	bottomRight := types.Position{X: t.Pos().X + t.Vision(), Y: t.Pos().Y - t.Vision()}

	// Get the positions inside the vision square from the environment
	perceivedObjects := e.PerceivedObjects(topLeft, bottomRight)
	perceivedAgents := e.PerceivedAgents(topLeft, bottomRight, t.id)

	// Get the objects not seen by the agent
	CantSeeBehindObjects := []obj.Object{}
	for _, obj := range perceivedObjects {
		if utils.Contains(t.CantSeeBehind(), obj.Name()) {
			CantSeeBehindObjects = append(CantSeeBehindObjects, obj)
		}
	}
	// Filter out positions to avoid regarding the agent position
	// If a CantSeeBehindObjects object is in the vision square, agent can't see behind it,
	// depending on the angle between the agent and the position to avoid, the agent can't see every positions behind it following the perspective logic

	// Get the positions behind the CantSeeBehindObjects objects
	objectsBehindPositions := []types.Position{}

	for _, object := range perceivedObjects {
		if !utils.Contains(CantSeeBehindObjects, object) {
			continue
		} else {
			// If a objCantSeeBehindObject is in the vision square, the agent can't see behind it
			// Get the angle between the agent and the position to avoid
			objectCenter := object.Center()
			angle := utils.GetAngle(t.Pos(), objectCenter)

			// Get the perceivedObjects behind the current position to avoid
			objectsBehindCurrentObject := utils.GetObjectsBehindPositions(t.Pos(), angle, topLeft, bottomRight)

			for _, position := range objectsBehindCurrentObject {
				objectsBehindPositions = append(objectsBehindPositions, position)
			}
		}
	}

	// Checks if the perceivedObjects are in a objectsBehindPositions position and remove them if they are
	perceivedObjects = utils.RemoveObjectsBehindPositions(perceivedObjects, objectsBehindPositions)
	perceivedAgents = removeAgentsBehindPositions(perceivedAgents, objectsBehindPositions)

	return perceivedObjects, perceivedAgents
}
