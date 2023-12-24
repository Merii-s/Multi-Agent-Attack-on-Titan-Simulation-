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

	//Il faut  ajouter les durrees de vie des objets
)

type Object struct {
	name ObjectName
	tl   Position
	life int
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
	default:
		h = HField
		w = WField
	}

	hb := make([]Position, 2)
	hb[0] = f.TL()
	hb[1] = Position{X: f.tl.X + w, Y: f.tl.Y + h}
	return hb
}
