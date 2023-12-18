package pkg

import (
	"sync"
)

const (
	CWall = 20

	WField = 40
	HField = 34

	CGrass = 50

	WBHouse1 = 55
	HBhouse1 = 46

	WBHouse2 = 42
	HBhouse2 = 55

	WSHouse = 43
	HSHouse = 40
)

type Environment struct {
	sync.RWMutex
	agents  []AgentI
	village *Village
	/*
		agentCount int

		//A modifier
		in        uint64
		out       uint64
		noopCount uint64*/
}

type Village struct {
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

// Behavior ?
