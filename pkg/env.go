package pkg

import (
	"sync"
	"time"
)

func createStaticObjects(H int, W int) []Object {
	objects := make([]Object, 0)
	var obj *Object

	//Grass Sprites
	nW := W / CGrass
	nH := H / CGrass
	for i := 0; i < nH; i++ {
		for j := 0; j < nW; j++ {
			obj = NewObject(Grass, Position{X: j * CGrass, Y: i * CGrass}, 1000000000)
			objects = append(objects, *obj)
		}
	}

	//Wall Sprites
	wall_Tl_X := int(0.2 * float32(W))
	wall_Tl_Y := int(0.2 * float32(H))
	wWall := int(0.6 * float32(W))       //Largeur du mur
	hWall := int((6. / 7.) * float32(H)) //Hauteur du mur
	nW = wWall / CWall                   //Nombre de sprites en largeur
	nH = hWall / CWall                   //Nombre de sprites en hauteur
	//Mur du Nord, Horizontal
	for i := 0; i < nW+1; i++ {
		obj = NewObject(Wall, Position{X: wall_Tl_X + i*CWall, Y: wall_Tl_Y}, 1000000000)
		objects = append(objects, *obj)
	}
	//Murs cotes, Verticaux
	for i := 1; i < nH; i++ {
		obj = NewObject(Wall, Position{X: wall_Tl_X, Y: wall_Tl_Y + i*CWall}, 1000000000)
		objects = append(objects, *obj)
		obj = NewObject(Wall, Position{X: wall_Tl_X + wWall, Y: wall_Tl_Y + i*CWall}, 1000000000)
		objects = append(objects, *obj)
	}

	//Champs
	nW = ((wWall - 2*CWall) / 2) / WField //Nombre de sprites en largeur
	for i := 0; i < nW*2; i++ {
		var x, y int
		if i < 7 {
			x = wall_Tl_X + CWall + int(wWall/4) + i*WField
			y = wall_Tl_Y + CWall + int(hWall/10)
		} else {
			x = wall_Tl_X + CWall + int(wWall/4) + (i-7)*WField
			y = wall_Tl_Y + CWall + int(hWall/10) + HField*2
		}
		obj = NewObject(Field, Position{X: x, Y: y}, 1000000000)
		objects = append(objects, *obj)
	}

	//Petites maisons
	coefsCoords := [][]float32{{0.29, 0.4}, {1 - 0.29, 0.4}, {0.29, 0.85}, {1 - 0.29, 0.85}, {0.29, 0.55}, {1 - 0.29, 0.65}, {0.5, 0.85}}
	for _, coords := range coefsCoords {

		obj = NewObject(SmallHouse, Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, 1000000000)
		objects = append(objects, *obj)
	}

	//Grandes maisons 1 et 2
	coefsCoords = [][]float32{{0.29, 0.7}, {0.5, 0.55}, {0.60, 0.7}}
	for _, coords := range coefsCoords {
		obj = NewObject(BigHouse1, Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, 1000000000)
		objects = append(objects, *obj)
	}

	//Donjons
	obj = NewObject(Dungeon, Position{X: int(0.2*float32(W) + CWall), Y: int(0.2*float32(H) + CWall)}, 1000000000)
	objects = append(objects, *obj)
	obj = NewObject(Dungeon, Position{X: int(0.8*float32(W) - CWall - WDungeon/2), Y: int(0.2*float32(H) + CWall)}, 1000000000)
	objects = append(objects, *obj)

	//Adding characters
	k := 0.6 * float32(H)
	names := []ObjectName{Eren, Mikasa, MaleVillager, FemaleVillager, BasicTitan1, BasicTitan2, BeastTitan, ColossalTitan, ArmoredTitan, FemaleTitan, ErenTitanS, JawTitan, MaleSoldier, FemaleSoldier}
	coords := [][]int{{500, 350}, {510, 350}, {520, 350}, {530, 350}, {560, 400}, {590, 400}, {610, 350}, {610 + 30, 350}, {wall_Tl_X + CWall + int(wWall/4) + 25, int(k)}, {wall_Tl_X + CWall + int(wWall/4) + 50, int(k)}, {wall_Tl_X + CWall + int(wWall/4) + 75, int(k)}, {wall_Tl_X + CWall + int(wWall/4) + 100, int(k) + 20}, {450, 350}, {460, 350}}
	for i := 0; i < len(names); i++ {
		obj = NewObject(names[i], Position{X: coords[i][0], Y: coords[i][1]}, 1000000000)
		objects = append(objects, *obj)
	}

	return objects
}

func SendEnvToUI(e *Environment, c chan *Environment) {

}

func MoveColossal(e *Environment, c chan *Environment, wg *sync.WaitGroup) {
	var ind int
	coords := [][]int{{250, 250}, {750, 250}, {750, 450}, {250, 450}}
	for i, _ := range e.objects {
		if e.objects[i].name == ColossalTitan {
			ind = i
			break
		}
	}

	for {
		for _, pos := range coords {
			wg.Add(1)
			go func(pos []int) {
				e.objects[ind].SetPosition(Position{pos[0], pos[1]})
				c <- e
				time.Sleep(10 * time.Millisecond)
				wg.Done()
			}(pos)
			wg.Wait()
		}
	}
}

func NewEnvironement(H int, W int) *Environment {
	objects := createStaticObjects(H, W)
	return &Environment{objects: objects}
}

func (p *Environment) Objects() []Object {
	return p.objects
}

func (p *Environment) Agents() []AgentI {
	return p.agents
}
