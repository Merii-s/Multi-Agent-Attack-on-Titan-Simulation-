package agt

import (
	pkg "AOT/pkg"
)

type TitanI interface {
	pkg.AgentI
	attack()
}

type Titan struct {
	id       pkg.Id
	pos      pkg.Position
	heightp  int
	reach    int
	strength int
	speed    int
	height   int
}

func NewTitan(id pkg.Id, p pkg.Position, hp int, r int, s int, spd int, h int) *Titan {
	return &Titan{id: id, pos: p, heightp: hp, reach: r, strength: s, speed: spd, height: h}
}

func (t *Titan) Id() pkg.Id {
	return t.id
}

func (t *Titan) Pos() pkg.Position {
	return t.pos
}

func (t *Titan) Heightp() int {
	return t.heightp
}

func (t *Titan) Reach() int {
	return t.reach
}

func (t *Titan) Strength() int {
	return t.strength
}

func (t *Titan) Speed() int {
	return t.speed
}

func (t *Titan) Height() int {
	return t.height
}

func (t *Titan) SetPos(pos pkg.Position) {
	t.pos = pos
}
