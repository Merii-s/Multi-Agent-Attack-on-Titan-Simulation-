package agt

import (
	pkg "AOT/pkg"
)

type HumanI interface {
	pkg.AgentI
	escape()
}

type Human struct {
	agentAttributes pkg.Agent
}

func NewHuman(id pkg.Id, tl pkg.Position, life int, reach int, strength int, speed int, vision int, object pkg.ObjectName) *Human {
	return &Human{agentAttributes: *pkg.NewAgent(id, tl, life, reach, strength, speed, vision, object)}
}
