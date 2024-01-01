package main

import (
	agt_utils "AOT/agt/agt_utils"
	env "AOT/agt/env"
	gui "AOT/gui"
	obj "AOT/pkg/obj"
	params "AOT/pkg/parameters"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sync.Mutex
	c        chan *env.Environment
	elements []obj.Object
	newEnv   bool
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
	fmt.Println("DRAWING...")
	g.drawEnvironment(screen)
	if g.newEnv {
		g.newEnv = false
		g.Unlock()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return params.ScreenWidth, params.ScreenHeight
}

func (g *Game) ListenToSimu() {
	var e *env.Environment
	fmt.Println("UI is listening to simu...")
	for {
		if !g.newEnv {
			e = <-g.c
			fmt.Println("RECEIVED NEW ENV...")
			g.Lock()
			g.elements = make([]obj.Object, len(e.Objects())+len(e.Agents()))
			g.elements = append(g.elements, e.Objects()...)
			for _, agent := range e.Agents() {
				g.elements = append(g.elements, agent.Object())
			}
			g.newEnv = true
		}
	}
}

func NewEnvironement(H int, W int) *env.Environment {
	objects := env.CreateStaticObjects(H, W)
	agents := agt_utils.CreateAgents(H, W, objects)
	return &env.Environment{Agts: agents, Objs: objects}
}

var wgPercept sync.WaitGroup
var wgDeliberate sync.WaitGroup
var wgAct sync.WaitGroup

func main() {
	g := Game{c: make(chan *env.Environment), newEnv: false}
	e := NewEnvironement(params.ScreenHeight, params.ScreenWidth)
	// go env.Simu(e, &wgPercept, &wgDeliberate, &wgAct, g.c)
	//
	simu := env.NewSimulation(params.NB_AGENTS, params.MaxStep, params.MaxDuration, e, g.c)
	go g.ListenToSimu()
	go simu.Run()
	ebiten.SetWindowSize(params.ScreenWidth, params.ScreenHeight)
	ebiten.SetWindowTitle("AOT Simulation")

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
