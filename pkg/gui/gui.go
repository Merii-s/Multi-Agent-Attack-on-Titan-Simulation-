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
	wallImg    *ebiten.Image
	fieldImg   *ebiten.Image
	grassImg   *ebiten.Image
	sHouseImg  *ebiten.Image
	bHouse1Img *ebiten.Image
	bHouse2Img *ebiten.Image
	dungeonImg *ebiten.Image
	cannonImg  *ebiten.Image
	erenImg    *ebiten.Image

	op ebiten.DrawImageOptions

	e *pkg.Environment
)

func init() {
	var (
		err1, err2, err3, err4, err5, err6, err7, err8 error
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
	erenImg, _, _ = ebitenutil.NewImageFromFile(pkg.GetImagePath("eren_sprite"))

	// //Lecture des fichiers png dans des variables
	// wallImg, _, err1 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("wall_sprite"))
	// fieldImg, _, err2 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("wheat_V2"))
	// grassImg, _, err3 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("grass_spriteV4"))
	// sHouseImg, _, err4 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("small_house_sprite"))
	// bHouse1Img, _, err5 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("big_house_sprite"))
	// bHouse2Img, _, err6 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("big_house_spriteV2"))
	// dungeonImg, _, err7 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("dungeon_sprite"))
	// cannonImg, _, err8 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("dungeon_sprite"))

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
