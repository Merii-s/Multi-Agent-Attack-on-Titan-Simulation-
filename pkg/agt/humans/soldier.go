package agt

import (
	pkg "AOT/pkg"
)

type SoldierI interface {
	HumanI
	Gard()
	Attack()
}

type Soldier struct {
	attributes Human
}

func (*Soldier) Percept(*pkg.Environment) {

}

func (*Soldier) Deliberate() {

}

func (*Soldier) Act(*pkg.Environment) {

}

func (*Soldier) Start() {

}

func (*Soldier) Id() {

}

func (*Soldier) move() {

}

func (*Soldier) eat() {

}

func (*Soldier) sleep() {

}

func (*Soldier) Gard() {

}

func (*Soldier) Attack() {

}
