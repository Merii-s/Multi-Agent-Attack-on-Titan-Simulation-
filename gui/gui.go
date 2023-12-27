package gui

import (
	pkg "AOT/pkg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	imgFiles = []string{
		"wall_sprite", "wheat_V2", "grass_spriteV4", "small_house_sprite",
		"big_house_sprite", "big_house_spriteV2", "dungeon_sprite", "dungeon_sprite",
		"eren_small_sprite", "mikasa_sprite", "male_villager_sprite", "female_villager_sprite",
		"basic_titan1_sprite", "basic_titan2_sprite", "beast_titan_sprite_V2", "armored_titan_sprite",
		"colossal_titan_sprite", "female_titan_sprite", "eren_titan_sprite", "jaw_titan_sprite",
		"male_soldier_sprite", "female_soldier_sprite",
	}
)

func Load_Sprites() ([]error, map[string]**ebiten.Image) {

	var errs []error
	imageVariables := make(map[string]**ebiten.Image)

	//Lecture des images et stockage dans imageVariables
	for _, file := range imgFiles {
		img, _, err := ebitenutil.NewImageFromFile(pkg.GetImagePath(file))
		if err != nil {
			errs = append(errs, err)
		}
		imageVariables[file] = &img
	}

	return errs, imageVariables
}

func DrawSprite(screen *ebiten.Image, o pkg.Object, imageVariables map[string]**ebiten.Image) {
	var (
		img *ebiten.Image
		op  ebiten.DrawImageOptions
	)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(o.TL().X), float64(o.TL().Y))

	switch o.Name() {
	case pkg.Field:
		img = *imageVariables["wheat_V2"]
	case pkg.BigHouse1:
		img = *imageVariables["big_house_sprite"]
	case pkg.BigHouse2:
		img = *imageVariables["big_house_spriteV2"]
	case pkg.Dungeon:
		img = *imageVariables["dungeon_sprite"]
	case pkg.Grass:
		img = *imageVariables["grass_spriteV4"]
	case pkg.Wall:
		img = *imageVariables["wall_sprite"]
	case pkg.Eren:
		img = *imageVariables["eren_small_sprite"]
	case pkg.Mikasa:
		img = *imageVariables["mikasa_sprite"]
	case pkg.MaleVillager:
		img = *imageVariables["male_villager_sprite"]
	case pkg.FemaleVillager:
		img = *imageVariables["female_villager_sprite"]
	case pkg.BasicTitan1:
		img = *imageVariables["basic_titan1_sprite"]
	case pkg.BasicTitan2:
		img = *imageVariables["basic_titan2_sprite"]
	case pkg.BeastTitan:
		img = *imageVariables["beast_titan_sprite_V2"]
	case pkg.ArmoredTitan:
		img = *imageVariables["armored_titan_sprite"]
	case pkg.FemaleTitan:
		img = *imageVariables["female_titan_sprite"]
	case pkg.ColossalTitan:
		img = *imageVariables["colossal_titan_sprite"]
	case pkg.ErenTitanS:
		img = *imageVariables["eren_titan_sprite"]
	case pkg.JawTitan:
		img = *imageVariables["jaw_titan_sprite"]
	case pkg.MaleSoldier:
		img = *imageVariables["male_soldier_sprite"]
	case pkg.FemaleSoldier:
		img = *imageVariables["female_soldier_sprite"]
	default:
		img = *imageVariables["small_house_sprite"]
	}

	screen.DrawImage(img, &op)
}
