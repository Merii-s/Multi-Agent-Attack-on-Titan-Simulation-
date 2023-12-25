package pkg

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

	WEren = 10
	HEren = 19

	WMikasa = 9
	HMikasa = 19

	WMaleVillager = 10
	HMaleVillager = 18

	WFemaleVillager = 10
	HFemaleVillager = 17

	WBasicTitanF = 21
	HBasicTitanF = 40

	WBasicTitanM = 22
	HBasicTitanM = 40

	WArmoredTitan = 20
	HArmoredTitan = 49

	WBeastTitan = 31
	HBeastTitan = 64

	WColossalTitan = 28
	HColossalTitan = 65

	WErenTitan = 20
	HErenTitan = 50

	WFemaleTitan = 19
	HFemaleTitan = 50

	WJawTitan = 32
	HJawTitan = 34

	WSoldierM = 15
	HSoldierM = 22

	WSoldierF = 20
	HSoldierF = 22
)

type Object struct {
	name    ObjectName
	tl      Position
	life    int
	idAgent Id
}

func NewObject(name ObjectName, tl Position, life int) *Object {
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

func (o *Object) SetPosition(p Position) {
	o.tl = p
}

func (o *Object) Name() ObjectName {
	return o.name
}

func (o *Object) TL() Position {
	return o.tl
}

func (o *Object) IdAgent() Id {
	return o.idAgent
}

type FieldObject struct {
	attr    Object
	reserve int
}

func NewField(tl Position, life int, reserve int) *FieldObject {
	return &FieldObject{
		attr: Object{
			name: Field,
			tl:   tl,
			life: life,
		},
		reserve: reserve,
	}
}

func (f *Object) hitbox() []Position {
	var w, h int
	switch f.name {

	case Wall:
		h = CWall
		w = CWall
	case Grass:
		h = CGrass
		w = CGrass
	case BigHouse1:
		h = HBHouse1
		w = WBHouse1
	case SmallHouse:
		h = HSHouse
		w = WSHouse
	case Dungeon:
		h = HDungeon
		w = WDungeon
	case BigHouse2:
		h = HBHouse2
		w = WBHouse2
	case Eren:
		h = HEren
		w = WEren
	case Mikasa:
		h = HMikasa
		w = WMikasa
	case MaleVillager:
		h = HMaleVillager
		w = WMaleVillager
	case FemaleVillager:
		h = HFemaleVillager
		w = WFemaleVillager
	case BasicTitan1:
		h = HBasicTitanF
		w = WBasicTitanF
	case BasicTitan2:
		h = HBasicTitanM
		w = WBasicTitanM
	case ArmoredTitan:
		h = HArmoredTitan
		w = WArmoredTitan
	case BeastTitan:
		h = HBeastTitan
		w = WBeastTitan
	case ColossalTitan:
		h = HColossalTitan
		w = WColossalTitan
	case ErenTitanS:
		h = HErenTitan
		w = WErenTitan
	case FemaleTitan:
		h = HFemaleTitan
		w = WFemaleTitan
	case JawTitan:
		h = HJawTitan
		w = WJawTitan
	case FemaleSoldier:
		h = HSoldierF
		w = WSoldierF
	case MaleSoldier:
		h = HSoldierM
		w = WSoldierM
	default:
		h = HField
		w = WField
	}

	hb := make([]Position, 2)
	hb[0] = f.TL()
	hb[1] = Position{X: f.tl.X + w, Y: f.tl.Y + h}

	return hb
}

// getter for Object
