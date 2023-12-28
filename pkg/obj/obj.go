package obj

import (
	gui "AOT/pkg/gui"
	types "AOT/pkg/types"
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

func (o *Object) Life() int {
	return o.life
}

func (o *Object) SetPosition(p types.Position) {
	o.tl = p
}

func (o *Object) Name() types.ObjectName {
	return o.name
}

func (o Object) GetName() types.ObjectName {
	return o.name
}

func (o *Object) TL() types.Position {
	return o.tl
}

type FieldObject struct {
	attr    Object
	reserve int
}

func NewField(tl types.Position, life int, reserve int) *FieldObject {
	return &FieldObject{
		attr: Object{
			name: types.Field,
			tl:   tl,
			life: life,
		},
		reserve: reserve,
	}
}

func (f *Object) Hitbox() (hb []types.Position) {
	var w, h int
	switch f.name {

	case types.Wall:
		h = gui.CWall
		w = gui.CWall
	case types.Grass:
		h = gui.CGrass
		w = gui.CGrass
	case types.BigHouse1:
		h = gui.HBHouse1
		w = gui.WBHouse1
	case types.SmallHouse:
		h = gui.HSHouse
		w = gui.WSHouse
	case types.Dungeon:
		h = gui.HDungeon
		w = gui.WDungeon
	case types.BigHouse2:
		h = gui.HBHouse2
		w = gui.WBHouse2
	case types.Eren:
		h = gui.HEren
		w = gui.WEren
	case types.Mikasa:
		h = gui.HMikasa
		w = gui.WMikasa
	case types.MaleVillager:
		h = gui.HMaleVillager
		w = gui.WMaleVillager
	case types.FemaleVillager:
		h = gui.HFemaleVillager
		w = gui.WFemaleVillager
	case types.BasicTitan1:
		h = gui.HBasicTitanF
		w = gui.WBasicTitanF
	case types.BasicTitan2:
		h = gui.HBasicTitanM
		w = gui.WBasicTitanM
	case types.ArmoredTitan:
		h = gui.HArmoredTitan
		w = gui.WArmoredTitan
	case types.BeastTitan:
		h = gui.HBeastTitan
		w = gui.WBeastTitan
	case types.ColossalTitan:
		h = gui.HColossalTitan
		w = gui.WColossalTitan
	case types.ErenTitanS:
		h = gui.HErenTitan
		w = gui.WErenTitan
	case types.FemaleTitan:
		h = gui.HFemaleTitan
		w = gui.WFemaleTitan
	case types.JawTitan:
		h = gui.HJawTitan
		w = gui.WJawTitan
	case types.FemaleSoldier:
		h = gui.HSoldierF
		w = gui.WSoldierF
	case types.MaleSoldier:
		h = gui.HSoldierM
		w = gui.WSoldierM
	default:
		h = gui.HField
		w = gui.WField
	}
	hb = make([]types.Position, 0)
	hb = append(hb, f.TL())
	hb = append(hb, types.Position{X: f.tl.X + w, Y: f.tl.Y + h})
	return hb
}

// return the center position of the object
func (o *Object) Center() types.Position {
	var w, h int
	switch o.name {
	case types.Wall:
		h = gui.CWall
		w = gui.CWall
	case types.Grass:
		h = gui.CGrass
		w = gui.CGrass
	case types.BigHouse1:
		h = gui.HBHouse1
		w = gui.WBHouse1
	case types.SmallHouse:
		h = gui.HSHouse
		w = gui.WSHouse
	case types.Dungeon:
		h = gui.HDungeon
		w = gui.WDungeon
	case types.BigHouse2:
		h = gui.HBHouse2
		w = gui.WBHouse2
	case types.Eren:
		h = gui.HEren
		w = gui.WEren
	case types.Mikasa:
		h = gui.HMikasa
		w = gui.WMikasa
	case types.MaleVillager:
		h = gui.HMaleVillager
		w = gui.WMaleVillager
	case types.FemaleVillager:
		h = gui.HFemaleVillager
		w = gui.WFemaleVillager
	default:
		h = gui.HField
		w = gui.WField
	}
	return types.Position{X: o.tl.X + w/2, Y: o.tl.Y + h}
}
