package obj

import (
	params "AOT/pkg/parameters"
	types "AOT/pkg/types"
)

type Object struct {
	name    types.ObjectName
	tl      types.Position
	life    int
	reserve int
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

func (o *Object) SetPosition(p types.Position) { o.tl = p }

func (o *Object) Name() types.ObjectName {
	return o.name
}

func (o Object) GetName() types.ObjectName {
	return o.name
}

func (o *Object) TL() types.Position {
	return o.tl
}

func (o *Object) Reserve() int { return o.reserve }

func (o *Object) SetReserve(r int) { o.reserve = r }

// TODO : fonction qui remplie la reserve de la field tous les x pas de temps

func (f *Object) Hitbox() (hb []types.Position) {
	var w, h int
	switch f.name {

	case types.Wall:
		h = params.CWall
		w = params.CWall
	case types.Grass:
		h = params.CGrass
		w = params.CGrass
	case types.BigHouse:
		h = params.HBHouse1
		w = params.WBHouse1
	case types.SmallHouse:
		h = params.HSHouse
		w = params.WSHouse
	case types.Dungeon:
		h = params.HDungeon
		w = params.WDungeon
	case types.Eren:
		h = params.HEren
		w = params.WEren
	case types.Mikasa:
		h = params.HMikasa
		w = params.WMikasa
	case types.MaleCivilian:
		h = params.HMaleCivilian
		w = params.WMaleCivilian
	case types.FemaleCivilian:
		h = params.HFemaleCivilian
		w = params.WFemaleCivilian
	case types.BasicTitan1:
		h = params.HBasicTitanF
		w = params.WBasicTitanF
	case types.BasicTitan2:
		h = params.HBasicTitanM
		w = params.WBasicTitanM
	case types.ArmoredTitan:
		h = params.HArmoredTitan
		w = params.WArmoredTitan
	case types.BeastTitan:
		h = params.HBeastTitan
		w = params.WBeastTitan
	case types.ColossalTitan:
		h = params.HColossalTitan
		w = params.WColossalTitan
	case types.ErenTitanS:
		h = params.HErenTitan
		w = params.WErenTitan
	case types.FemaleTitan:
		h = params.HFemaleTitan
		w = params.WFemaleTitan
	case types.JawTitan:
		h = params.HJawTitan
		w = params.WJawTitan
	case types.FemaleSoldier:
		h = params.HSoldierF
		w = params.WSoldierF
	case types.MaleSoldier:
		h = params.HSoldierM
		w = params.WSoldierM
	default:
		h = params.HField
		w = params.WField
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
		h = params.CWall
		w = params.CWall
	case types.Grass:
		h = params.CGrass
		w = params.CGrass
	case types.BigHouse:
		h = params.HBHouse1
		w = params.WBHouse1
	case types.SmallHouse:
		h = params.HSHouse
		w = params.WSHouse
	case types.Dungeon:
		h = params.HDungeon
		w = params.WDungeon
	case types.Eren:
		h = params.HEren
		w = params.WEren
	case types.Mikasa:
		h = params.HMikasa
		w = params.WMikasa
	case types.MaleCivilian:
		h = params.HMaleCivilian
		w = params.WMaleCivilian
	case types.FemaleCivilian:
		h = params.HFemaleCivilian
		w = params.WFemaleCivilian
	default:
		h = params.HField
		w = params.WField
	}
	return types.Position{X: o.tl.X + w/2, Y: o.tl.Y + h}
}

func (o *Object) TakeDamage(dmg int) {
	o.life -= dmg
}
