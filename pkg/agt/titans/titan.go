package agt

import (
	pkg "AOT/pkg"
)

type TitanI interface {
	pkg.AgentI
	attack()
	regenerate()
}

type Titan struct {
	agentAttributes pkg.Agent
	height          int
	regenRate       int
}

func NewTitan(id pkg.Id, t pkg.Type, p pkg.Position, hp int, r int, s int, spd int, h int, regen int) *Titan {
	return &Titan{agentAttributes: *pkg.NewAgent(id, t, p, hp, r, s, spd), height: h, regenRate: regen}
}

func (t *Titan) RegenRate() int {
	return t.regenRate
}

func (t *Titan) Height() int {
	return t.height
}
