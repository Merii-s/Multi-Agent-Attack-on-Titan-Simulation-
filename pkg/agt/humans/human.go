package agt

import (
	pkg "AOT/pkg"
)

type HumanI interface {
	pkg.AgentI
	escape()
}

type Human struct {
	Id       pkg.Id
	Pos      pkg.Position
	Hp       int
	Reach    int
	Strength int
	Speed    int
}

func NewHuman(id pkg.Id, p pkg.Position, hp int, r int, s int, spd int) *Human {
	return &Human{Id: id, Pos: p, Hp: hp, Reach: r, Strength: s, Speed: spd}
}
