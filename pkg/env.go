package pkg

import (
	"sync"
)

func createStaticObjects(H int, W int) []Object {
	objects := make([]Object, 0)
	var obj *Object
	nb_objects := 0

	//Grass Sprites
	nW := W / CGrass
	nH := H / CGrass
	for i := 0; i < nH; i++ {
		for j := 0; j < nW; j++ {
			obj = NewObject(Grass, Position{X: j * CGrass, Y: i * CGrass}, 1000000000)
			objects = append(objects, *obj)
			nb_objects = nb_objects + 1
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
		nb_objects = nb_objects + 1
	}
	//Murs cotes, Verticaux
	for i := 1; i < nH; i++ {
		obj = NewObject(Wall, Position{X: wall_Tl_X, Y: wall_Tl_Y + i*CWall}, 1000000000)
		objects = append(objects, *obj)
		obj = NewObject(Wall, Position{X: wall_Tl_X + wWall, Y: wall_Tl_Y + i*CWall}, 1000000000)
		objects = append(objects, *obj)
		nb_objects = nb_objects + 2
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
		nb_objects = nb_objects + 1
	}

	//Petites maisons
	coefsCoords := [][]float32{{0.29, 0.4}, {1 - 0.29, 0.4}, {0.29, 0.85}, {1 - 0.29, 0.85}, {0.29, 0.55}, {1 - 0.29, 0.65}, {0.5, 0.85}}
	for _, coords := range coefsCoords {

		obj = NewObject(SmallHouse, Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, 1000000000)
		objects = append(objects, *obj)
		nb_objects = nb_objects + 1
	}

	//Grandes maisons 1 et 2
	coefsCoords = [][]float32{{0.29, 0.7}, {0.5, 0.55}, {0.60, 0.7}}
	for _, coords := range coefsCoords {
		obj = NewObject(BigHouse1, Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, 1000000000)
		objects = append(objects, *obj)
		nb_objects = nb_objects + 1
	}

	//Donjons
	obj = NewObject(Dungeon, Position{X: int(0.2*float32(W) + CWall), Y: int(0.2*float32(H) + CWall)}, 1000000000)
	objects = append(objects, *obj)
	obj = NewObject(Dungeon, Position{X: int(0.8*float32(W) - CWall - WDungeon/2), Y: int(0.2*float32(H) + CWall)}, 1000000000)
	objects = append(objects, *obj)

	obj = NewObject(ColossalTitan, Position{X: 640, Y: 350}, 1000000000)
	objects = append(objects, *obj)
	nb_objects = nb_objects + 3

	return objects
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
				wg.Done()
			}(pos)
			wg.Wait()
		}
	}
}

func createHumans(objs []Object, tl_village Position, br_village Position) []Object {
	var human *Object
	var end bool
	var counter int
	var nb_grass int

	//A modifier quand le construteur d'humain sera pret
	humans := make([]Object, 0)

	for i := 0; i < NB_HUMANS; i++ {
		if i < NB_VILLAGERS {
			x, y := GetRandomCoords(tl_village, br_village)
			if i < NB_VILLAGERS/2 {
				//A modifier quand le construteur d'humain sera pret
				human = NewObject(MaleVillager, Position{x, y}, 100)
			} else {
				//A modifier quand le construteur d'humain sera pret
				human = NewObject(FemaleVillager, Position{x, y}, 100)
			}
		}

		end = false
		for !end {
			counter = 0
			nb_grass = 0
			for _, obj := range objs {
				if obj.name != Grass {
					if DetectCollision(*human, obj) {
						x, y := GetRandomCoords(tl_village, br_village)
						human.SetPosition(Position{x, y})
						break
					} else {
						counter++
					}
				} else {
					nb_grass++
				}
			}

			for _, hu := range humans {
				if DetectCollision(*human, hu) {
					x, y := GetRandomCoords(tl_village, br_village)
					human.SetPosition(Position{x, y})
					break
				} else {
					counter++
				}
			}

			if counter == len(objs)+len(humans)-nb_grass {
				end = true
			}
		}

		humans = append(humans, *human)
	}
	return humans
}

func NewEnvironement(H int, W int) *Environment {
	objects := createStaticObjects(H, W)
	humans := createHumans(objects, Position{X: int(0.2*float32(W)) + CWall, Y: int(0.2*float32(H)) + CWall}, Position{X: int(0.8 * float32(W)), Y: 700 - HMaleVillager})

	return &Environment{agents: humans, objects: objects}
}

func (p *Environment) Objects() []Object {
	return p.objects
}

// A modifier quand le construteur d'agent sera pret
func (p *Environment) Agents() []Object {
	return p.agents
}
