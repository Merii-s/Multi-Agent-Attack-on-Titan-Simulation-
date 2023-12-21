package gui

import (
	pkg "AOT/pkg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

var (
	//var qui vont acceuillir les sprites
	wallImg           *ebiten.Image
	fieldImg          *ebiten.Image
	grassImg          *ebiten.Image
	sHouseImg         *ebiten.Image
	bHouse1Img        *ebiten.Image
	bHouse2Img        *ebiten.Image
	dungeonImg        *ebiten.Image
	cannonImg         *ebiten.Image
	erenImg           *ebiten.Image
	mikasaImg         *ebiten.Image
	maleVillagerImg   *ebiten.Image
	femaleVillagerImg *ebiten.Image
	basicTitan1Img    *ebiten.Image
	basicTitan2Img    *ebiten.Image
	beastTitanImg     *ebiten.Image

	op ebiten.DrawImageOptions

	e *pkg.Environment
)

func init() {
	var (
		err1, err2, err3, err4, err5, err6, err7, err8, err9, err10, err11, err12, err13, err14, err15 error
	)

	//Lecture des fichiers png dans des variables
	wallImg, _, err1 = ebitenutil.NewImageFromFile(pkg.GetImagePath("wall_sprite"))
	fieldImg, _, err2 = ebitenutil.NewImageFromFile(pkg.GetImagePath("wheat_V2"))
	grassImg, _, err3 = ebitenutil.NewImageFromFile(pkg.GetImagePath("grass_spriteV4"))
	sHouseImg, _, err4 = ebitenutil.NewImageFromFile(pkg.GetImagePath("small_house_sprite"))
	bHouse1Img, _, err5 = ebitenutil.NewImageFromFile(pkg.GetImagePath("big_house_sprite"))
	bHouse2Img, _, err6 = ebitenutil.NewImageFromFile(pkg.GetImagePath("big_house_spriteV2"))
	dungeonImg, _, err7 = ebitenutil.NewImageFromFile(pkg.GetImagePath("dungeon_sprite"))
	cannonImg, _, err8 = ebitenutil.NewImageFromFile(pkg.GetImagePath("dungeon_sprite"))
	erenImg, _, err9 = ebitenutil.NewImageFromFile(pkg.GetImagePath("eren_small_sprite"))
	mikasaImg, _, err10 = ebitenutil.NewImageFromFile(pkg.GetImagePath("mikasa_sprite"))
	maleVillagerImg, _, err11 = ebitenutil.NewImageFromFile(pkg.GetImagePath("male_villager_sprite"))
	femaleVillagerImg, _, err12 = ebitenutil.NewImageFromFile(pkg.GetImagePath("female_villager_sprite"))
	basicTitan1Img, _, err13 = ebitenutil.NewImageFromFile(pkg.GetImagePath("basic_titan1_sprite"))
	basicTitan2Img, _, err14 = ebitenutil.NewImageFromFile(pkg.GetImagePath("basic_titan2_sprite"))
	beastTitanImg, _, err15 = ebitenutil.NewImageFromFile(pkg.GetImagePath("beast_titan_sprite"))

	if err1 != nil {
		log.Fatal(err1)
	} else if err2 != nil {
		log.Fatal(err2)
	} else if err3 != nil {
		log.Fatal(err3)
	} else if err4 != nil {
		log.Fatal(err4)
	} else if err5 != nil {
		log.Fatal(err5)
	} else if err6 != nil {
		log.Fatal(err6)
	} else if err7 != nil {
		log.Fatal(err7)
	} else if err8 != nil {
		log.Fatal(err8)
	} else if err9 != nil {
		log.Fatal(err9)
	} else if err10 != nil {
		log.Fatal(err10)
	} else if err11 != nil {
		log.Fatal(err11)
	} else if err12 != nil {
		log.Fatal(err12)
	} else if err13 != nil {
		log.Fatal(err13)
	} else if err14 != nil {
		log.Fatal(err14)
	} else if err15 != nil {
		log.Fatal(err15)
	}

	e = pkg.NewEnvironement(700, 1000)
}

func drawSprite(screen *ebiten.Image, o pkg.Object) {
	var (
		img *ebiten.Image
		err error
	)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(o.TL().X), float64(o.TL().Y))

	switch o.Name() {
	case pkg.Field:
		img = fieldImg
	case pkg.BigHouse1:
		img = bHouse1Img
	case pkg.BigHouse2:
		img = bHouse2Img
	case pkg.Dungeon:
		img = dungeonImg
	case pkg.Grass:
		img = grassImg
	case pkg.Wall:
		img = wallImg
	case pkg.Eren:
		img = erenImg
	case pkg.Mikasa:
		img = mikasaImg
	case pkg.MaleVillager:
		img = maleVillagerImg
	case pkg.FemaleVillager:
		img = femaleVillagerImg
	case pkg.BasicTitan1:
		img = basicTitan1Img
	case pkg.BasicTitan2:
		img = basicTitan2Img
	case pkg.BeastTitan:
		img = beastTitanImg
	default:
		img = sHouseImg
	}

	if err != nil {
		log.Fatal(err)
	} else {
		screen.DrawImage(img, &op)
	}
}

func eraseSprite(screen *ebiten.Image, o pkg.Object) {
	var (
		img *ebiten.Image
		err error
	)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(o.TL().X), float64(o.TL().Y))

	switch o.Name() {
	case pkg.Field:
		img = fieldImg
	case pkg.BigHouse1:
		img = bHouse1Img
	case pkg.BigHouse2:
		img = bHouse2Img
	case pkg.Dungeon:
		img = dungeonImg
	case pkg.Grass:
		img = grassImg
	case pkg.Wall:
		img = wallImg
	default:
		img = sHouseImg
	}

	if err != nil {
		log.Fatal(err)
	} else {
		screen.DrawImage(img, &op)
	}
}

func drawEnvironment(screen *ebiten.Image, env *pkg.Environment) {
	for _, o := range env.Objects() {
		if o.Life() > 0 {
			drawSprite(screen, o)
		}
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawEnvironment(screen, e)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1000, 700
}

func RunDisplay() {
	ebiten.SetWindowSize(1000, 700)
	ebiten.SetWindowTitle("AOT Simulation")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
