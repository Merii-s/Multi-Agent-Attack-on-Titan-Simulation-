package env

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
	Percept(*Environment /*, *sync.WaitGroup*/)
	Deliberate( /**sync.WaitGroup*/ )
	Act(*Environment /*, *sync.WaitGroup*/)
	Start(*Environment /*, *sync.WaitGroup, *sync.WaitGroup, *sync.WaitGroup, *sync.WaitGroup*/)
	Id() types.Id

	Move(pos types.Position)
	Eat()
	Sleep()
	AttackSuccess(spdAtk int, spdDef int) float64
	Attack(*AgentI)
	AgtSyncChan() chan int

	Pos() types.Position
	Vision() int
	PerceivedObjects() []*obj.Object
	PerceivedAgents() []*AgentI
	Object() obj.Object
	Agent() *Agent
	SetPos(types.Position)
}
