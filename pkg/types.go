package pkg

import (
	"AOT/pkg"
	"sync"
)

type Environment struct {
	sync.RWMutex

	//A modifier quand le constructeur d'agent sera pret, C'est un tableau d'agents
	agents []Object

	objects []Object

	screenH int
	screenW int

	agentCount int

	// Day/Night cycle
	day bool
}

type AgentI interface {
	Percept(*Environment)
	Deliberate()
	Act(*Environment)
	Start(*Environment)
	Id() Id

	Move(pos Position)
	Eat()
	Sleep()
	AttackSuccess(spdAtk int, spdDef int) float64
	Attack(pkg.AgentI)

	Pos() Position
	Vision() int
	PerceivedObjects() []Object
	PerceivedAgents() []AgentI
	Object() *Object
	Agent() *Agent
}

type Id string

type Position struct {
	X int
	Y int
}

type ObjectName string

const (
	Grass          ObjectName = "Grass"
	Field          ObjectName = "Field"
	Wall           ObjectName = "Wall"
	BigHouse1      ObjectName = "BigHouse1"
	BigHouse2      ObjectName = "BigHouse2"
	SmallHouse     ObjectName = "SmallHouse"
	Dungeon        ObjectName = "Dungeon"
	Eren           ObjectName = "Eren"
	Mikasa         ObjectName = "Mikasa"
	MaleVillager   ObjectName = "MaleVillager"
	FemaleVillager ObjectName = "FemaleVillager"
	BasicTitan1    ObjectName = "BasicTitan1"
	BasicTitan2    ObjectName = "BasicTitan2"
	BeastTitan     ObjectName = "BeastTitan"
	ColossalTitan  ObjectName = "ColossalTitan"
	ArmoredTitan   ObjectName = "ArmoredTitan"
	ErenTitanS     ObjectName = "ErenTitan"
	FemaleTitan    ObjectName = "FemaleTitan"
	JawTitan       ObjectName = "JawTitan"
	MaleSoldier    ObjectName = "MaleSoldier"
	FemaleSoldier  ObjectName = "FemaleSoldier"
)

type BehaviorI interface {
	Percept(*Environment)
	Deliberate()
	Act(*Environment)
}
