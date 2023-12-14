package agt

import (
	pkg "AOT/pkg"
)

type BasicTitanI interface {
	TitanI
}

type BasicTitan struct {
	attributes Titan
	type_agt   pkg.Type
}

func NewBasicTitan(id pkg.Id, topLeft pkg.Position, bottomRight pkg.Position, hp int, reach int, speed int, strength int, height int) *BasicTitan {
	x, y := pkg.GetRandomCoords(topLeft, bottomRight)
	pos := pkg.Position{X: x, Y: y}
	atts := NewTitan(id, pos, hp, reach, strength, speed, height)
	return &BasicTitan{attributes: *atts, type_agt: pkg.Titan}
}

func (*BasicTitan) Percept(e *pkg.Environment) {

}

func (*BasicTitan) Deliberate() {

}

func (*BasicTitan) Act(e *pkg.Environment) {

}

func (*BasicTitan) Start() {

}

func (bt *BasicTitan) Id() pkg.Id {
	return bt.attributes.Id()
}

func (bt *BasicTitan) Move(x_variation int, y_variation int) {
	new_X_pos := bt.attributes.Pos().X + x_variation
	new_Y_pos := bt.attributes.Pos().Y + y_variation
	new_pos := pkg.Position{X: new_X_pos, Y: new_Y_pos}
	bt.attributes.SetPos(new_pos)
}

func (*BasicTitan) Eat() {

}

func (*BasicTitan) Sleep() {

}

func (*BasicTitan) Attack() {

}
