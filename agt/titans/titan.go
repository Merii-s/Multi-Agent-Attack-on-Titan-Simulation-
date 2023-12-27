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
	regenRate       int
}

func NewTitan(id pkg.Id, tl pkg.Position, life int, r int, s int, spd int, v int, o pkg.ObjectName, regen int) *Titan {
	return &Titan{
		agentAttributes: *pkg.NewAgent(id, tl, life, r, s, spd, v, o),
		regenRate:       regen}
}

func (t *Titan) RegenRate() int {
	return t.regenRate
}
