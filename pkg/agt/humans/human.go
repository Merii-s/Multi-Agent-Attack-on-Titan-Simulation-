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

func NewHuman(id pkg.Id, tl pkg.Position, life int, r int, s int, spd int, v int, o pkg.ObjectName) *Human {
	return &Human{agentAttributes: *pkg.NewAgent(id, tl, life, r, s, spd, v, o)}
}
