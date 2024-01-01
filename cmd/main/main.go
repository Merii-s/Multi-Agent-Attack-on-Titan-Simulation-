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
	if g.newEnv {
		time.Sleep(100 * time.Millisecond)
		g.drawEnvironment(screen)
		wgSimu.Done()
		//g.Unlock()
		g.newEnv = false
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return params.ScreenWidth, params.ScreenHeight
}

func (g *Game) ListenToSimu() {
	var e *env.Environment
	fmt.Println("UI is listening to simu...")
	for {
		e = <-g.c
		wgSimu.Wait()
		wgSimu.Add(1)
		//g.Lock()
		g.elements = make([]obj.Object, len(e.Objects())+len(e.Agents()))
		g.elements = append(g.elements, e.Objects()...)
		for _, agent := range e.Agents() {
			g.elements = append(g.elements, agent.Object())
		}
		g.newEnv = true
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
var wgSimu sync.WaitGroup

func main() {
	g := Game{c: make(chan *env.Environment)}
	e := NewEnvironement(params.ScreenHeight, params.ScreenWidth)
	go env.Simu(e, &wgPercept, &wgDeliberate, &wgAct, g.c)
	go g.ListenToSimu()
	ebiten.SetWindowSize(params.ScreenWidth, params.ScreenHeight)
	ebiten.SetWindowTitle("AOT Simulation")

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
	// for i, o := range e.Objects() {
	// 	if o.Name() == types.Wall {
	// 		e.Objects()[i].SetLife(150)
	// 		fmt.Println(e.Objects()[i].Life())
	// 	}
	// }
	// o := obj.NewObject(types.Wall, types.Position{X: 5, Y: 3}, 0)
	// fmt.Println(o.Life())
	// o.SetLife(150)
	// fmt.Println(o.Life())
}
