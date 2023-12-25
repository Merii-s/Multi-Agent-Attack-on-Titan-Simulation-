package pkg

type Agent struct {
	id               Id
	reach            int
	strength         int
	speed            int
	maxHP            int
	vision           int
	perceivedObjects []Object
	perceivedAgents  []AgentI
	cantSeeBehind    []ObjectName
	object           *Object
}

func NewAgent(id Id, tl Position, life int, reach int, strength int, speed int, vision int, name ObjectName) *Agent {
	return &Agent{
		id:               id,
		reach:            reach,
		strength:         strength,
		speed:            speed,
		vision:           vision,
		perceivedObjects: []Object{},
		perceivedAgents:  []AgentI{},
		cantSeeBehind:    []ObjectName{},
		object:           NewObject(name, tl, life),
	}
}

func (t *Agent) Id() Id {
	return t.id
}

func (t *Agent) Pos() Position { return t.object.TL() }

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

func (t *Agent) Hp() int { return t.object.life }

func (t *Agent) SetPos(pos Position) { t.object.tl = pos }

func (t *Agent) SetHp(hp int) { t.object.life = hp }

func (t *Agent) Vision() int { return t.vision }

func (t *Agent) PerceivedObjects() []Object { return t.perceivedObjects }

func (t *Agent) PerceivedAgents() []AgentI { return t.perceivedAgents }

func (t *Agent) CantSeeBehind() []ObjectName { return t.cantSeeBehind }

func (t *Agent) Object() Object { return *t.object }

func (t *Agent) AddPerceivedObject(obj Object) { t.perceivedObjects = append(t.perceivedObjects, obj) }

func (t *Agent) AddPerceivedAgent(agt AgentI) { t.perceivedAgents = append(t.perceivedAgents, agt) }

// returns a list of objects that the agent can see
// the vision is a square centered on the agent position
func (t *Agent) GetVision(e *Environment) ([]Object, []AgentI) {
	// Get the top left and bottom right positions of the vision square
	topLeft := Position{t.Pos().X - t.Vision(), t.Pos().Y + t.Vision()}
	bottomRight := Position{t.Pos().X + t.Vision(), t.Pos().Y - t.Vision()}

	// Get the positions inside the vision square from the environment
	perceivedObjects := e.PerceivedObjects(topLeft, bottomRight)
	perceivedAgents := e.PerceivedAgents(topLeft, bottomRight, t.id)

	// Get the objects not seen by the agent
	CantSeeBehindObjects := []Object{}
	for _, obj := range perceivedObjects {
		if contains(t.CantSeeBehind(), obj.Name()) {
			CantSeeBehindObjects = append(CantSeeBehindObjects, obj)
		}
	}
	// Filter out positions to avoid regarding the agent position
	// If a CantSeeBehindObjects object is in the vision square, agent can't see behind it,
	// depending on the angle between the agent and the position to avoid, the agent can't see every positions behind it following the perspective logic

	// Get the positions behind the CantSeeBehindObjects objects
	objectsBehindPositions := []Position{}

	for _, object := range perceivedObjects {
		if !contains(CantSeeBehindObjects, object) {
			continue
		} else {
			// If a objCantSeeBehindObject is in the vision square, the agent can't see behind it
			// Get the angle between the agent and the position to avoid
			objectCenter := object.Center()
			angle := getAngle(t.Pos(), objectCenter)

			// Get the perceivedObjects behind the current position to avoid
			objectsBehindCurrentObject := getObjectsBehindPositions(t.Pos(), angle, topLeft, bottomRight)

			for _, position := range objectsBehindCurrentObject {
				objectsBehindPositions = append(objectsBehindPositions, position)
			}
		}
	}

	// Checks if the perceivedObjects are in a objectsBehindPositions position and remove them if they are
	perceivedObjects = removeObjectsBehindPositions(perceivedObjects, objectsBehindPositions)
	perceivedAgents = removeAgentsBehindPositions(perceivedAgents, objectsBehindPositions)

	return perceivedObjects, perceivedAgents
}
