package gui

import (
	obj "AOT/pkg/obj"
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"

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
		"male_soldier_sprite", "female_soldier_sprite", "jaw_titan_human_sprite", "armored_titan_human_sprite",
		"colossal_titan_human_sprite", "female_titan_human_sprite", "beast_titan_human_sprite",
	}
)

func Load_Sprites() ([]error, map[string]**ebiten.Image) {

	var errs []error
	imageVariables := make(map[string]**ebiten.Image)

	//Lecture des images et stockage dans imageVariables
	for _, file := range imgFiles {
		img, _, err := ebitenutil.NewImageFromFile(utils.GetImagePath(file))
		if err != nil {
			errs = append(errs, err)
		}
		imageVariables[file] = &img
	}

	return errs, imageVariables
}

func DrawSprite(screen *ebiten.Image, o obj.Object, imageVariables map[string]**ebiten.Image) {
	var (
		img *ebiten.Image
		op  ebiten.DrawImageOptions
	)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(o.TL().X), float64(o.TL().Y))

	switch o.Name() {
	case types.Field:
		img = *imageVariables["wheat_V2"]
	case types.BigHouse1:
		img = *imageVariables["big_house_sprite"]
	case types.BigHouse2:
		img = *imageVariables["big_house_spriteV2"]
	case types.Dungeon:
		img = *imageVariables["dungeon_sprite"]
	case types.Grass:
		img = *imageVariables["grass_spriteV4"]
	case types.Wall:
		img = *imageVariables["wall_sprite"]
	case types.Eren:
		img = *imageVariables["eren_small_sprite"]
	case types.Mikasa:
		img = *imageVariables["mikasa_sprite"]
	case types.MaleVillager:
		img = *imageVariables["male_villager_sprite"]
	case types.FemaleVillager:
		img = *imageVariables["female_villager_sprite"]
	case types.BasicTitan1:
		img = *imageVariables["basic_titan1_sprite"]
	case types.BasicTitan2:
		img = *imageVariables["basic_titan2_sprite"]
	case types.BeastTitan:
		img = *imageVariables["beast_titan_sprite_V2"]
	case types.BeastTitanHuman:
		img = *imageVariables["beast_titan_human_sprite"]
	case types.ArmoredTitan:
		img = *imageVariables["armored_titan_sprite"]
	case types.ArmoredTitanHuman:
		img = *imageVariables["armored_titan_human_sprite"]
	case types.FemaleTitan:
		img = *imageVariables["female_titan_sprite"]
	case types.FemaleTitanHuman:
		img = *imageVariables["female_titan_human_sprite"]
	case types.ColossalTitan:
		img = *imageVariables["colossal_titan_sprite"]
	case types.ColossalTitanHuman:
		img = *imageVariables["colossal_titan_human_sprite"]
	case types.ErenTitanS:
		img = *imageVariables["eren_titan_sprite"]
	case types.JawTitan:
		img = *imageVariables["jaw_titan_sprite"]
	case types.JawTitanHuman:
		img = *imageVariables["jaw_titan_human_sprite"]
	case types.MaleSoldier:
		img = *imageVariables["male_soldier_sprite"]
	case types.FemaleSoldier:
		img = *imageVariables["female_soldier_sprite"]
	default:
		img = *imageVariables["small_house_sprite"]
	}

	screen.DrawImage(img, &op)
}
