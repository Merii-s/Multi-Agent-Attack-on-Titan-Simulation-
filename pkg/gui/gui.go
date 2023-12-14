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

	//Recuperation du fichier png cotenant l'image du mur
	wallImg, _, err1 = ebitenutil.NewImageFromFile(utils.GetPath("wall_sprite"))
	//Reglages de la taille de l'image
	spriteX := (Width / 2)
	spriteY := (Height / 2)
	// imageBounds := wallImg.Bounds()
	// width := imageBounds.Dx()
	// height := imageBounds.Dy()
	//opWall.GeoM.Scale(0.005*float64(width), 0.005*float64(height))
	opWall.GeoM.Translate(float64(spriteX), float64(spriteY))

	fieldImg, _, err2 = ebitenutil.NewImageFromFile(utils.GetPath("wheat_sprite"))
	//Reglages de la taille de l'image
	//opField.GeoM.Scale(0.005, 0.005)
	//opField.GeoM.Translate(Width/2, Height/2)

	grassImg, _, err3 = ebitenutil.NewImageFromFile(utils.GetPath("grass_sprite"))
	//Reglages de la taille de l'image
	//spriteX = (Width - Width*0.005) / 2
	//spriteY = (Height - Height*0.005) / 2
	//opGrass.GeoM.Translate(float64(0), float64(0))
	//opGrass.GeoM.Scale(0.005, 0.005)

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

func drawCityBorders(sprite *ebiten.Image, screen *ebiten.Image) {
	imageBounds := sprite.Bounds()
	cSprite := imageBounds.Dx()
	xTL := 0.2 * Width
	yTL := 0.2 * Height
	xBR := 0.8 * Width
	yBR := 0.8 * Height
	// mur haut horizontal G --> D
	drawWall(int(xTL), int(yTL), int(xBR)+cSprite, int(yTL), true, sprite, screen)
	// mur gauche vertical H --> B
	drawWall(int(xTL), int(yTL+float64(cSprite)), int(xTL), int(yBR), false, sprite, screen)
	// mur bas horizontal G --> D
	drawWall(int(xTL), int(yBR), int(xBR)+cSprite, int(yBR), true, sprite, screen)
	// mur droit vertical H --> B
	drawWall(int(xBR), int(yTL+float64(cSprite)), int(xBR), int(yBR), false, sprite, screen)

	// screen.DrawImage(wallImg /*&opField*/, &opWall)
}

func drawGrass() {

}

func drawEnvironment() {

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(grassImg /*&opField*/, &opGrass)
	drawCityBorders(wallImg, screen)
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
