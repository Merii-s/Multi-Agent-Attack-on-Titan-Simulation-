package pkg

import (
	"sync"
)

type Environment struct {
	sync.RWMutex
	agents  []AgentI
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
	Start()
	Id()

	move(Position)
	eat()
	sleep()
	attack_success(spd_atk int, reach_atk int, spd_def int) float64
	Pos()
}

type Id string

type Position struct {
	X int
	Y int
}

type Type string

const (
	Civilian     Type = "Civilian"
	Soldier      Type = "Soldier"
	Titan        Type = "Titan"
	SpecialTitan Type = "SpecialTitan"
	ErenTitan    Type = "ErenTitan"
)

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
)

type BehaviorI interface {
	Percept(*Environment)
	Deliberate()
	Act(*Environment)
}
