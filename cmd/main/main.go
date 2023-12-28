package main

import (
	env "AOT/agt/env"
	gui "AOT/gui"
	obj "AOT/pkg/obj"
	params "AOT/pkg/parameters"

	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sync.Mutex
	c        chan *env.Environment
	elements []obj.Object
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
	var e *env.Environment
	for {
		e = <-g.c
		g.Lock()
		g.elements = make([]obj.Object, len(e.Objects())+len(e.Agents()))
		g.elements = append(g.elements, e.Objects()...)
		for _, agent := range e.Agents() {
			g.elements = append(g.elements, agent.Object())
		}
	}
}

func NewEnvironement(H int, W int) *Environment {
	objects := createStaticObjects(H, W)
	humans := createHumans(objects, types.Position{X: int(0.2*float32(W)) + params.CWall, Y: int(0.2*float32(H)) + params.CWall}, types.Position{X: int(0.8 * float32(W)), Y: H})
	titans := createTitans(H, W)
	// for _, titan := range titans {
	// 	fmt.Println(titan)
	// }
	merged_agents := make([]AgentI, len(titans)+len(humans))
	merged_agents = append(merged_agents, humans...)
	merged_agents = append(merged_agents, titans...)
	return &Environment{agents: merged_agents, objects: objects}
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
