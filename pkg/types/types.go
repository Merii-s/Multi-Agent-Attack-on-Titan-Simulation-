package types

import (
	"AOT/agt/env"
	"AOT/pkg/obj"
	pkg "AOT/pkg/utilitaries"
	"math"
)

type Id string

type Position struct {
	X int
	Y int
}

func (position *Position) Equals(p2 Position) bool {
	return position.X == p2.X && position.Y == p2.Y
}

type ObjectName string

const (
	Grass              ObjectName = "Grass"
	Field              ObjectName = "Field"
	Wall               ObjectName = "Wall"
	BigHouse           ObjectName = "BigHouse"
	SmallHouse         ObjectName = "SmallHouse"
	Dungeon            ObjectName = "Dungeon"
	Eren               ObjectName = "Eren"
	Mikasa             ObjectName = "Mikasa"
	MaleCivilian       ObjectName = "MaleCivilian"
	FemaleCivilian     ObjectName = "FemaleCivilian"
	BasicTitan1        ObjectName = "BasicTitan1"
	BasicTitan2        ObjectName = "BasicTitan2"
	BeastTitan         ObjectName = "BeastTitan"
	BeastTitanHuman    ObjectName = "BeastTitanHuman"
	ColossalTitan      ObjectName = "ColossalTitan"
	ColossalTitanHuman ObjectName = "ColossalTitanHuman"
	ArmoredTitan       ObjectName = "ArmoredTitan"
	ArmoredTitanHuman  ObjectName = "ArmoredTitanHuman"
	ErenTitanS         ObjectName = "ErenTitan"
	FemaleTitan        ObjectName = "FemaleTitan"
	FemaleTitanHuman   ObjectName = "FemaleTitanHuman"
	JawTitan           ObjectName = "JawTitan"
	JawTitanHuman      ObjectName = "JawTitanHuman"
	MaleSoldier        ObjectName = "MaleSoldier"
	FemaleSoldier      ObjectName = "FemaleSoldier"
)

func (position Position) ClosestPosition(positions []Position) Position {
	// Get the closest position from the list
	closestPosition := positions[0]
	for _, pos := range positions {
		if position.Distance(pos) < position.Distance(closestPosition) {
			closestPosition = pos
		}
	}
	return closestPosition
}

func (position Position) ClosestAgent(agents []env.AgentI) (env.AgentI, Position) {
	// Get the closest position from the list
	closestAgent := agents[0]
	closestAgentPosition := agents[0].Agent().ObjectP().Hitbox()[0]
	for _, agt := range agents {
		for _, pos := range pkg.GetPositionsInHitbox(agt.Agent().ObjectP().Hitbox()[0], agt.Agent().ObjectP().Hitbox()[1]) {
			if position.Distance(pos) < position.Distance(closestAgentPosition) {
				closestAgent = agt
				closestAgentPosition = pos
			}
		}
	}
	return closestAgent, closestAgentPosition
}

func (position Position) ClosestObject(objects []obj.Object) (obj.Object, Position) {
	// Get the closest position from the list
	closestObject := objects[0]
	closestObjectPosition := objects[0].Hitbox()[0]
	for _, object := range objects {
		for _, pos := range pkg.GetPositionsInHitbox(object.Hitbox()[0], object.Hitbox()[1]) {
			if position.Distance(pos) < position.Distance(closestObjectPosition) {
				closestObject = object
				closestObjectPosition = pos
			}
		}
	}
	return closestObject, closestObjectPosition
}

func (Position) Distance(pos Position) float64 {
	return math.Sqrt(math.Pow(float64(pos.X), 2) + math.Pow(float64(pos.Y), 2))
}
