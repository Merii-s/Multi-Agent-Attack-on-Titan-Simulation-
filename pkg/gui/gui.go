package gui

import (
	utils "AOT/pkg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//types "AOT_local/AOT/pkg"
)

const (
	Height = 700
	Width  = 1400
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
	opWall.GeoM.Translate(float64(spriteX), float64(spriteY))
	opWall.GeoM.Scale(0.005, 0.005)

	fieldImg, _, err2 = ebitenutil.NewImageFromFile(utils.GetPath("wheat_sprite"))
	//Reglages de la taille de l'image
	//opField.GeoM.Scale(0.005, 0.005)
	//opField.GeoM.Translate(Width/2, Height/2)

	grassImg, _, err3 = ebitenutil.NewImageFromFile(utils.GetPath("grass_sprite"))
	//Reglages de la taille de l'image
	//spriteX = (Width - Width*0.005) / 2
	//spriteY = (Height - Height*0.005) / 2
	//opGrass.GeoM.Translate(float64(spriteX), float64(spriteY))
	//opGrass.GeoM.Scale(0.005, 0.005)

	if err1 != nil {
		log.Fatal(err1)
	} else if err2 != nil {
		log.Fatal(err1)
	} else if err3 != nil {
		log.Fatal(err1)
	}
}

func drawWall() {

}

func drawGrass() {

}

func drawEnvironment() {

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(fieldImg /*&opField*/, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Height, Width
}

func RunDisplay() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("AOT Simulation")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
