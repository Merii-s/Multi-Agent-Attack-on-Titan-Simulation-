package pkg

import (
	"sync"
)

type Environment struct {
	sync.RWMutex
	agents     []AgentI
	abjects    []Object
	agentCount int
	// Day/Night cycle

	//A modifier
	in        uint64
	out       uint64
	noopCount uint64
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
	Grass      ObjectName = "Grass"
	Field      ObjectName = "Field"
	Wall       ObjectName = "Wall"
	BigHouse   ObjectName = "BigHouse"
	SmallHouse ObjectName = "SmallHouse"
	Dungeon    ObjectName = "Dungeon"
)

type BehaviorI interface {
	Percept(*Environment)
	Deliberate()
	Act(*Environment)
}
