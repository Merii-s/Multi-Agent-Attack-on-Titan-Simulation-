package types

import "math"

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

func (Position) Distance(pos Position) float64 {
	return math.Sqrt(math.Pow(float64(pos.X), 2) + math.Pow(float64(pos.Y), 2))
}
