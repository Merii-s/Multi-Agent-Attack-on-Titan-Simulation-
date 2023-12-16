package pkg

import (
	"sync"
)

type Environment struct {
	sync.RWMutex
	agents     []AgentI
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

	move()
	eat()
	sleep()
	attack_success(spd_atk int, reach_atk int, spd_def int) float64
}

type Id string

type Position struct {
	X int
	Y int
}

type Type string

const (
	Civilian     = "Civilian"
	Soldier      = "Soldier"
	Titan        = "Titan"
	SpecialTitan = "SpecialTitan"
	ErenTitan    = "ErenTitan"
)

type BehaviorI interface {
	Percept(*Environment)
	Deliberate()
	Act(*Environment)
}
