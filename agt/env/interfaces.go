package agt

import (
	obj "AOT/pkg/obj"
	types "AOT/pkg/types"
)

type BehaviorI interface {
	Percept(*Environment)
	Deliberate()
	Act(*Environment)
}

type AgentI interface {
	Percept(*Environment)
	Deliberate()
	Act(*Environment)
	Start(*Environment)
	Id() types.Id

	Move(pos types.Position)
	Eat()
	Sleep()
	AttackSuccess(spdAtk int, spdDef int) float64
	Attack(AgentI)

	Pos() types.Position
	Vision() int
	PerceivedObjects() []obj.Object
	PerceivedAgents() []AgentI
	Object() obj.Object
	Agent() *Agent
}
