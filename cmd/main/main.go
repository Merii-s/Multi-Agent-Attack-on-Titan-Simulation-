package main

import (
	pkg "AOT/pkg"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	sync.Mutex
	c    chan *pkg.Environment
	objs []pkg.Object
}

var (
	imageVariables = make(map[string]**ebiten.Image)
	imgFiles       = []string{
		"wall_sprite", "wheat_V2", "grass_spriteV4", "small_house_sprite",
		"big_house_sprite", "big_house_spriteV2", "dungeon_sprite", "dungeon_sprite",
		"eren_small_sprite", "mikasa_sprite", "male_villager_sprite", "female_villager_sprite",
		"basic_titan1_sprite", "basic_titan2_sprite", "beast_titan_sprite_V2", "armored_titan_sprite",
		"colossal_titan_sprite", "female_titan_sprite", "eren_titan_sprite", "jaw_titan_sprite",
		"male_soldier_sprite", "female_soldier_sprite",
	}
)

func init() {
	var errs []error

	//Lecture des images et stockage dans imageVariables
	for _, file := range imgFiles {
		img, _, err := ebitenutil.NewImageFromFile(pkg.GetImagePath(file))
		if err != nil {
			errs = append(errs, err)
		}
		imageVariables[file] = &img
	}

	for _, err := range errs {
		if err != nil {
			log.Fatal(err)
		}
	}
}

func drawSprite(screen *ebiten.Image, o pkg.Object) {
	var (
		img *ebiten.Image
		err error
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

	if err != nil {
		log.Fatal(err)
	} else {
		screen.DrawImage(img, &op)
	}
}

func (g *Game) drawEnvironment(screen *ebiten.Image) {
	for _, o := range g.objs {
		if o.Life() > 0 {
			drawSprite(screen, o)
		}
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	time.Sleep(100 * time.Millisecond)
	g.drawEnvironment(screen)
	g.Unlock()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1000, 700
}

func (g *Game) ListenToSimu() {
	var e *pkg.Environment
	for {
		e = <-g.c
		g.Lock()
		g.objs = make([]pkg.Object, 10000)
		copy(g.objs, e.Objects())
	}
}

var wg1 sync.WaitGroup //Simulation waitgroup

func main() {
	g := Game{c: make(chan *pkg.Environment)}
	e := pkg.NewEnvironement(700, 1000)

	go pkg.MoveColossal(e, g.c, &wg1)
	go g.ListenToSimu()

	ebiten.SetWindowSize(1000, 700)
	ebiten.SetWindowTitle("AOT Simulation")

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
