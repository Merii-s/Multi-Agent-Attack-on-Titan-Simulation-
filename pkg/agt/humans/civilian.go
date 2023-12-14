package agt

import (
	pkg "AOT/pkg"
)

type CivilianI interface {
	HumanI
	build()
	getFood()
}

type Civilian struct {
	attributes Human
}

func NewCivilian(id pkg.Id, topLeft pkg.Position, bottomRight pkg.Position, hp int, reach int, speed int, strength int) *Civilian {
	x, y := pkg.GetRandomCoords(topLeft, bottomRight)
	pos := pkg.Position{X: x, Y: y}
	atts := NewHuman(id, pos, hp, reach, strength, speed)
	return &Civilian{attributes: *atts}
}

func (*Civilian) Percept(*pkg.Environment) {

}

func (*Civilian) Deliberate() {

}

func (*Civilian) Act(*pkg.Environment) {

}

func (*Civilian) Start() {

}

func (*Civilian) Id() {

}

func (*Civilian) move() {

}

func (*Civilian) eat() {

}

func (*Civilian) sleep() {

}

func (*Civilian) build() {

}

func (*Civilian) getFood()
