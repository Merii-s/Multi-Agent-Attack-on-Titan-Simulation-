package humans

import (
	env "AOT/agt/env"
	types "AOT/pkg/types"
)

type HumanI interface {
	env.AgentI
	escape()
}

type Human struct {
	agentAttributes env.Agent
}

func NewHuman(id types.Id, tl types.Position, life int, reach int, strength int, speed int, vision int, object types.ObjectName) *Human {
	return &Human{agentAttributes: *env.NewAgent(id, tl, life, reach, strength, speed, vision, object)}
}
