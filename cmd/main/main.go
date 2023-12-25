package main

import (
	pkg "AOT/pkg"
	gui "AOT/pkg/gui"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sync.Mutex
	c        chan *pkg.Environment
	elements []pkg.Object
}

var (
	imageVariables = make(map[string]**ebiten.Image)
)

func init() {
	var errs []error
	errs, imageVariables = gui.Load_Sprites()
	for _, err := range errs {
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) drawEnvironment(screen *ebiten.Image) {
	for _, o := range g.elements {
		if o.Life() > 0 {
			gui.DrawSprite(screen, o, imageVariables)
		}
	}
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
		g.elements = make([]pkg.Object, len(e.Objects())+len(e.Agents()))
		mergedSlice := append(e.Objects(), e.Agents()...)
		copy(g.elements, mergedSlice)
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
