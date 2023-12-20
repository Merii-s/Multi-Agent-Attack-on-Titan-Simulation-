package gui

import (
	pkg "AOT/pkg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//types "AOT_local/AOT/pkg"
)

const (
	Height = 700
	Width  = 1000
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

	opWall    ebiten.DrawImageOptions
	opGrass   ebiten.DrawImageOptions
	opField   ebiten.DrawImageOptions
	opHouse   ebiten.DrawImageOptions
	opDungeon ebiten.DrawImageOptions
	op        ebiten.DrawImageOptions

	e *pkg.Environment
)

func init() {
	var (
		err1, err2, err3, err4, err5, err6, err7, err8 error
	)

	//Lecture des fichiers png dans des variables
	wallImg, _, err1 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("wall_sprite"))
	fieldImg, _, err2 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("wheat_V2"))
	grassImg, _, err3 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("grass_spriteV4"))
	sHouseImg, _, err4 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("small_house_sprite"))
	bHouse1Img, _, err5 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("big_house_sprite"))
	bHouse2Img, _, err6 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("big_house_spriteV2"))
	dungeonImg, _, err7 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("dungeon_sprite"))
	cannonImg, _, err8 = ebitenutil.NewImageFromFile(pkg.GetPath_Win("dungeon_sprite"))

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
		img, _, err = ebitenutil.NewImageFromFile(pkg.GetPath_Win("wheat_V2"))
	case pkg.BigHouse1:
		img, _, err = ebitenutil.NewImageFromFile(pkg.GetPath_Win("big_house_sprite"))
	case pkg.BigHouse2:
		img, _, err = ebitenutil.NewImageFromFile(pkg.GetPath_Win("big_house_spriteV2"))
	case pkg.Dungeon:
		img, _, err = ebitenutil.NewImageFromFile(pkg.GetPath_Win("dungeon_sprite"))
	case pkg.Grass:
		img, _, err = ebitenutil.NewImageFromFile(pkg.GetPath_Win("grass_spriteV4"))
	case pkg.Wall:
		img, _, err = ebitenutil.NewImageFromFile(pkg.GetPath_Win("wall_sprite"))
	default:
		img, _, err = ebitenutil.NewImageFromFile(pkg.GetPath_Win("small_house_sprite"))
	}

	if err != nil {
		log.Fatal(err)
	} else {
		screen.DrawImage(img, &op)
	}
}

func drawEnvironment(screen *ebiten.Image, env *pkg.Environment) {
	for _, o := range env.Objects() {
		drawSprite(screen, o)
	}
}

// dir : horizontal et sprite bien dimensionné
func drawWall(Xs int, Ys int, Xe int, Ye int, dir bool, sprite *ebiten.Image, screen *ebiten.Image) {
	imageBounds := sprite.Bounds()
	w := imageBounds.Dx()
	h := imageBounds.Dy()
	if dir {
		nbSprite := (Xe - Xs) / w
		for i := 0; i < nbSprite; i++ {
			opWall.GeoM.Reset()
			opWall.GeoM.Translate(float64(Xs+i*w), float64(Ys))
			screen.DrawImage(sprite, &opWall)
		}
	} else {
		nbSprite := (Ye - Ys) / h
		for i := 0; i < nbSprite; i++ {
			opWall.GeoM.Reset()
			opWall.GeoM.Translate(float64(Xs), float64(Ys+i*h))
			screen.DrawImage(sprite, &opWall)
		}
	}
}

func drawDungeons(screen *ebiten.Image, dungeonImg *ebiten.Image, cWall int) {
	imageBounds := dungeonImg.Bounds()
	w := imageBounds.Dx()
	opDungeon.GeoM.Reset()
	opDungeon.GeoM.Translate(float64(0.2*Width+cWall), float64(0.2*Height+cWall))
	screen.DrawImage(dungeonImg, &opDungeon)
	opDungeon.GeoM.Reset()
	opDungeon.GeoM.Translate(float64(0.8*Width-cWall-w/2), float64(0.2*Height+cWall))
	screen.DrawImage(dungeonImg, &opDungeon)
}

func drawCannons(screen *ebiten.Image, cannonImg *ebiten.Image, cWall int) {
	imageBounds := cannonImg.Bounds()
	w := imageBounds.Dx()
	//h := imageBounds.Dy()

	opDungeon.GeoM.Reset()
	opDungeon.GeoM.Translate(float64(0.2*Width+cWall), float64(0.2*Height+cWall))
	screen.DrawImage(dungeonImg, &opDungeon)
	opDungeon.GeoM.Reset()
	opDungeon.GeoM.Translate(float64(0.8*Width-cWall-w/2), float64(0.2*Height+cWall))
	screen.DrawImage(dungeonImg, &opDungeon)
}

func drawSmallHouses(screen *ebiten.Image, sHouseImg *ebiten.Image) {
	coefsCoords := [][]float32{{0.29, 0.4}, {1 - 0.29, 0.4}, {0.29, 0.85}, {1 - 0.29, 0.85}, {0.29, 0.55}, {1 - 0.29, 0.65}, {0.5, 0.85}}
	for _, coords := range coefsCoords {
		opHouse.GeoM.Reset()
		opHouse.GeoM.Translate(float64(coords[0]*Width), float64(coords[1]*Height))
		screen.DrawImage(sHouseImg, &opHouse)
	}
}

func drawBigHouses(screen *ebiten.Image, bHouse1Img *ebiten.Image, bHouse2Img *ebiten.Image) {
	coefsCoords := [][]float32{{0.29, 0.7}, {0.5, 0.55}, {0.4, 0.55}, {0.62, 0.7}}
	for i, coords := range coefsCoords {
		opHouse.GeoM.Reset()
		opHouse.GeoM.Translate(float64(coords[0]*Width), float64(coords[1]*Height))
		if i < 2 {
			screen.DrawImage(bHouse1Img, &opHouse)
		} else {
			screen.DrawImage(bHouse2Img, &opHouse)
		}
	}
}

func drawHouses(screen *ebiten.Image, sHouseImg *ebiten.Image, bHouse1Img *ebiten.Image, bHouse2Img *ebiten.Image) {
	drawSmallHouses(screen, sHouseImg)
	drawBigHouses(screen, bHouse1Img, bHouse2Img)
}

// Visuel dans lequel la ville est un rectangle au centre de la screen
func drawCityRectWalls(screen *ebiten.Image, wallSprite *ebiten.Image, fieldSprite *ebiten.Image) {
	// dimensions des sprites
	wallBounds := wallSprite.Bounds()
	fieldBounds := fieldSprite.Bounds()
	cSpriteWall := wallBounds.Dx()
	cSpriteField := fieldBounds.Dx()

	xTL := 0.2 * Width
	yTL := 0.2 * Height
	xBR := 0.8 * Width
	yBR := 0.8 * Height

	// ------- Draw walls -------
	// mur haut horizontal G --> D
	drawWall(int(xTL), int(yTL), int(xBR)+cSpriteWall, int(yTL), true, wallSprite, screen)
	// mur gauche vertical H --> B
	drawWall(int(xTL), int(yTL+float64(cSpriteWall)), int(xTL), int(yBR), false, wallSprite, screen)
	// mur bas horizontal G --> D
	drawWall(int(xTL), int(yBR), int(xBR)+cSpriteWall, int(yBR), true, wallSprite, screen)
	// mur droit vertical H --> B
	drawWall(int(xBR), int(yTL+float64(cSpriteWall)), int(xBR), int(yBR), false, wallSprite, screen)

	// ------- Draw fields -------
	// field haut horizontal 1
	drawWall(int(xTL)+cSpriteWall+0.2*1000, int(yTL)+cSpriteWall+0.18*Height, int(xBR)-cSpriteWall-0.2*Width, int(yTL)+cSpriteWall+0.18*Height, true, fieldSprite, screen)
	drawWall(int(xTL)+cSpriteWall+0.2*1000, int(yTL)+cSpriteWall+0.18*Height+cSpriteField, int(xBR)-cSpriteWall-0.2*Width, int(yTL)-cSpriteWall+0.18*Height+cSpriteField, true, fieldSprite, screen)
}

// Visuel des murs de la ville comme nous l'avait decrit Massil
func drawCityBorderWalls(screen *ebiten.Image, wallSprite *ebiten.Image, fieldSprite *ebiten.Image) {
	// dimensions des sprites
	wallBounds := wallSprite.Bounds()
	fieldBounds := fieldSprite.Bounds()
	cSpriteWall := wallBounds.Dx()
	cSpriteField := fieldBounds.Dx()

	xTL := 0.2 * Width
	yTL := 0.2 * Height
	xBR := 0.8 * Width
	yBR := Height

	// ------- Draw walls -------
	// mur haut horizontal G --> D
	drawWall(int(xTL), int(yTL), int(xBR)+cSpriteWall, int(yTL), true, wallSprite, screen)
	// mur gauche vertical H --> B
	drawWall(int(xTL), int(yTL+float64(cSpriteWall)), int(xTL), int(yBR), false, wallSprite, screen)
	// mur droit vertical H --> B
	drawWall(int(xBR), int(yTL+float64(cSpriteWall)), int(xBR), int(yBR), false, wallSprite, screen)

	// ------- Draw fields -------
	// field haut horizontal 1
	drawWall(int(xTL)+cSpriteWall+0.2*1000, int(yTL)+cSpriteWall+0.18*Height, int(xBR)-cSpriteWall-0.2*Width, int(yTL)+cSpriteWall+0.18*Height, true, fieldSprite, screen)
	drawWall(int(xTL)+cSpriteWall+0.2*1000, int(yTL)+cSpriteWall+0.18*Height+4*cSpriteField, int(xBR)-cSpriteWall-0.2*Width, int(yTL)-cSpriteWall+0.18*Height+4*cSpriteField, true, fieldSprite, screen)

	drawHouses(screen, sHouseImg, bHouse1Img, bHouse2Img)
	drawDungeons(screen, dungeonImg, cSpriteWall)

}

func drawGrass(screen *ebiten.Image, grassImg *ebiten.Image) {
	bounds := grassImg.Bounds()
	cSpriteGrass := bounds.Dx()

	nbSpritesWidth := Width / cSpriteGrass
	nbSpritesHeight := Height / cSpriteGrass

	for i := 0; i < nbSpritesHeight; i++ {
		for j := 0; j < nbSpritesWidth; j++ {
			opGrass.GeoM.Reset()
			opGrass.GeoM.Translate(float64(j*cSpriteGrass), float64(i*cSpriteGrass))
			screen.DrawImage(grassImg, &opGrass)
		}
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//drawGrass(screen, grassImg)
	//drawCityBorderWalls(screen, wallImg, fieldImg)

	drawEnvironment(screen, e)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func RunDisplay() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("AOT Simulation")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
