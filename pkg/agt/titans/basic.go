package agt

import (
	pkg "AOT/pkg"
)

type BasicTitanI interface {
	TitanI
}

type BasicTitan struct {
	attributes Titan
}

func (*BasicTitan) Percept(*pkg.Environment) {

}

func (*BasicTitan) Deliberate() {

}

func (*BasicTitan) Act(*pkg.Environment) {

}

func (*BasicTitan) Start() {

}

func (*BasicTitan) Id() {

}

func (*BasicTitan) Move() {

}

func (*BasicTitan) Eat() {

}

func (*BasicTitan) Sleep() {

}

func (*BasicTitan) Attack() {

}
