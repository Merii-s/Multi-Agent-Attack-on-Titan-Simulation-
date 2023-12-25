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

func NewHuman(id pkg.Id, t pkg.Type, p pkg.Position, hp int, r int, s int, spd int) *Human {
	return &Human{agentAttributes: *pkg.NewAgent(id, t, p, hp, r, s, spd)}
}
