package obj

import (
	types "AOT/pkg"
)

const (
	CWall = 20

	WField = 40
	HField = 34

	CGrass = 50

	WBHouse1 = 55
	HBhouse1 = 46

	WBHouse2 = 42
	HBhouse2 = 55

	WSHouse = 43
	HSHouse = 40

	WDungeon = 29
	HDungeon = 52
)

type Sprite struct {
	Tl   types.Position
	life int
}

func (f *Sprite) hitbox() (br *types.Position) {
	return types.NewPosition(f.Tl.X()+WField, f.Tl.Y()+HField)
}
