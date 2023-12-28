package titans

import (
	env "AOT/agt/env"
	types "AOT/pkg/types"
)

type TitanI interface {
	env.AgentI
	attack()
	regenerate()
}

type Titan struct {
	agentAttributes env.Agent
	regenRate       int
}

func NewTitan(id types.Id, tl types.Position, life int, r int, s int, spd int, v int, o types.ObjectName, regen int) *Titan {
	return &Titan{
		agentAttributes: *env.NewAgent(id, tl, life, r, s, spd, v, o),
		regenRate:       regen}
}

func (t *Titan) RegenRate() int {
	return t.regenRate
}
