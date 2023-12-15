package gui

import (
	utils "AOT/pkg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//types "AOT_local/AOT/pkg"
)

const (
	Height = 500
	Width  = 1000
)

type Game struct{}

var (
	//var qui vont acceuillir le Sprite du mur
	wallImg  *ebiten.Image
	fieldImg *ebiten.Image
	grassImg *ebiten.Image

	opWall  ebiten.DrawImageOptions
	opGrass ebiten.DrawImageOptions
	opField ebiten.DrawImageOptions
)

func init() {
	var (
		err1, err2, err3 error
	)

	//Lecture des fichiers png dans des variables
	wallImg, _, err1 = ebitenutil.NewImageFromFile(utils.GetPath("wall_sprite"))
	fieldImg, _, err2 = ebitenutil.NewImageFromFile(utils.GetPath("wheat_sprite"))
	grassImg, _, err3 = ebitenutil.NewImageFromFile(utils.GetPath("grass_sprite"))

	if err1 != nil {
		log.Fatal(err1)
	} else if err2 != nil {
		log.Fatal(err1)
	} else if err3 != nil {
		log.Fatal(err1)
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

func drawCity(screen *ebiten.Image, wallSprite *ebiten.Image, fieldSprite *ebiten.Image) {
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

func drawGrass(screen *ebiten.Image) {
	screen.DrawImage(grassImg, &opGrass)
}

func drawEnvironment() {

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawGrass(screen)
	drawCity(screen, wallImg, fieldImg)
	//screen.DrawImage(fieldImg, &opField)
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
