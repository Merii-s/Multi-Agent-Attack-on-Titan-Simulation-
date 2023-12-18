package obj

import (
	types "AOT/pkg"
)

const (
	//Screen Dimensions
	Height = 700
	Width  = 1000

	//Sprite Dimensions
	CWall = 20

	WField = 40
	HField = 34

	CGrass = 50

	WBHouse1 = 55
	HBHouse1 = 46

	WBHouse2 = 42
	HBHouse2 = 55

	WSHouse = 43
	HSHouse = 40

	WDungeon = 29
	HDungeon = 52

	//Il faut  ajouter les durrees de vie des objets
)

type Object struct {
	name types.ObjectName
	tl   types.Position
	life int
}

func NewObject(name types.ObjectName, tl types.Position, life int) *Object {
	return &Object{
		name: name,
		tl:   tl,
		life: life,
	}
}

func (o *Object) SetLife(l int) {
	o.life = l
}

func (o *Object) SetPosition(p types.Position) {
	o.tl = p
}

type Field struct {
	attr    Object
	reserve int
}

func NewField(tl types.Position, life int, reserve int) *Field {

	return &Field{
		attr: Object{
			name: types.Field,
			tl:   tl,
			life: life,
		},
		reserve: reserve,
	}
}

func (f *Object) hitbox() (br *types.Position) {
	var w, h int
	switch f.name {
	case types.Wall:
		h = CWall
		w = CWall
	case types.Grass:
		h = CGrass
		w = CGrass
	case types.BigHouse:
		h = HBHouse1
		w = WBHouse1
	case types.SmallHouse:
		h = HSHouse
		w = WSHouse
	case types.Dungeon:
		h = HDungeon
		w = WDungeon
	default:
		h = HField
		w = WField
	}
	return &types.Position{X: f.tl.X + w, Y: f.tl.Y + h}
}
