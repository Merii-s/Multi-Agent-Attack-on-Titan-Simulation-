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
	Grass              ObjectName = "Grass"
	Field              ObjectName = "Field"
	Wall               ObjectName = "Wall"
	BigHouse1          ObjectName = "BigHouse1"
	BigHouse2          ObjectName = "BigHouse2"
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
