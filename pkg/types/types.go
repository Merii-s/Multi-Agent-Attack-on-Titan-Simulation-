package types

type Id string

type Position struct {
	X int
	Y int
}

func (p *Position) Equals(p2 Position) bool {
	return p.X == p2.X && p.Y == p2.Y
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
