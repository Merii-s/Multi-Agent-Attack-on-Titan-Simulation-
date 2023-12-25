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
			obj = NewObject(Grass, Position{X: j * CGrass, Y: i * CGrass}, GRASS_LIFE)
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
		obj = NewObject(Wall, Position{X: wall_Tl_X + i*CWall, Y: wall_Tl_Y}, WALL_LIFE)
		objects = append(objects, *obj)
		nb_objects = nb_objects + 1
	}
	//Murs cotes, Verticaux
	for i := 1; i < nH; i++ {
		obj = NewObject(Wall, Position{X: wall_Tl_X, Y: wall_Tl_Y + i*CWall}, WALL_LIFE)
		objects = append(objects, *obj)
		obj = NewObject(Wall, Position{X: wall_Tl_X + wWall, Y: wall_Tl_Y + i*CWall}, WALL_LIFE)
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
		obj = NewObject(Field, Position{X: x, Y: y}, FIELD_LIFE)
		objects = append(objects, *obj)
		nb_objects = nb_objects + 1
	}

	//Petites maisons
	coefsCoords := [][]float32{{0.29, 0.4}, {1 - 0.29, 0.4}, {0.29, 0.85}, {1 - 0.29, 0.85}, {0.29, 0.55}, {1 - 0.29, 0.65}, {0.5, 0.85}}
	for _, coords := range coefsCoords {

		obj = NewObject(SmallHouse, Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, SMALL_HOUSE_LIFE)
		objects = append(objects, *obj)
		nb_objects = nb_objects + 1
	}

	//Grandes maisons 1 et 2
	coefsCoords = [][]float32{{0.29, 0.7}, {0.5, 0.55}, {0.60, 0.7}}
	for _, coords := range coefsCoords {
		obj = NewObject(BigHouse1, Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, BIG_HOUSE_LIFE)
		objects = append(objects, *obj)
		nb_objects = nb_objects + 1
	}

	//Donjons
	obj = NewObject(Dungeon, Position{X: int(0.2*float32(W) + CWall), Y: int(0.2*float32(H) + CWall)}, DUNGEON_LIFE)
	objects = append(objects, *obj)
	obj = NewObject(Dungeon, Position{X: int(0.8*float32(W) - CWall - WDungeon/2), Y: int(0.2*float32(H) + CWall)}, DUNGEON_LIFE)
	objects = append(objects, *obj)

	obj = NewObject(ColossalTitan, Position{X: 640, Y: 350}, COLOSSAL_TITAN_LIFE)
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

	//A modifier quand le construteur d'humain sera pret
	humans := make([]Object, 0)

	for i := 0; i < NB_HUMANS; i++ {
		x, y := GetRandomCoords(tl_village, br_village)
		if i < NB_VILLAGERS {
			if i < NB_VILLAGERS/2 {
				//A modifier quand le construteur d'humain sera pret
				human = NewObject(MaleVillager, Position{x, y}, VILLAGER_LIFE)
			} else {
				//A modifier quand le construteur d'humain sera pret
				human = NewObject(FemaleVillager, Position{x, y}, VILLAGER_LIFE)
			}
		} else if i < NB_VILLAGERS+NB_SOLDIERS {
			if i < NB_VILLAGERS+NB_SOLDIERS/2 {
				//A modifier quand le construteur d'humain sera pret
				human = NewObject(MaleSoldier, Position{x, y}, SOLDIER_LIFE)
			} else {
				//A modifier quand le construteur d'humain sera pret
				human = NewObject(FemaleSoldier, Position{x, y}, SOLDIER_LIFE)
			}
		} else if i < NB_VILLAGERS+NB_SOLDIERS+1 {
			//A modifier quand le construteur d'humain sera pret
			human = NewObject(Eren, Position{X: x, Y: y}, EREN_LIFE)
		} else {
			//A modifier quand le construteur d'humain sera pret
			human = NewObject(Mikasa, Position{X: x, Y: y}, MIKASA_LIFE)
		}

		//Place l'humain dans des coordonnees aleatoires valides (i.e sans collisions) dans le village
		humans = PlaceHuman(objs, humans, human, tl_village, br_village)
	}
	return humans
}

func createTitans(H int, W int) []Object {
	var titan *Object
	tl_screen := Position{X: 0, Y: 0}
	br_screen := Position{X: W, Y: H}

	dir := 0

	//A modifier quand le construteur de titan sera pret
	titans := make([]Object, 0)

	for i := 0; i < NB_TITANS; i++ {
		x, y := GetRandomCoords(tl_screen, br_screen)

		if dir == 0 {
			y = y - H
			dir = 1
		} else if dir == 1 {
			x = x - W
			dir = 2
		} else {
			x = x + W
			dir = 0
		}

		if i < NB_BASIC_TITANS {
			if i < NB_BASIC_TITANS/2 {
				//A modifier quand le construteur d'humain sera pret
				titan = NewObject(BasicTitan1, Position{x, y}, BASIC_TITAN_LIFE)
			} else {
				//A modifier quand le construteur d'humain sera pret
				titan = NewObject(BasicTitan2, Position{x, y}, BASIC_TITAN_LIFE)
			}
		} else if i < NB_BASIC_TITANS+NB_SPECIAL_TITANS {
			if i < NB_BASIC_TITANS+1 {
				//A modifier quand le construteur d'humain sera pret
				titan = NewObject(ColossalTitan, Position{x, y}, COLOSSAL_TITAN_LIFE)
			} else if i < NB_BASIC_TITANS+2 {
				//A modifier quand le construteur d'humain sera pret
				titan = NewObject(BeastTitan, Position{x, y}, BEAST_TITAN_LIFE)
			} else if i < NB_BASIC_TITANS+3 {
				titan = NewObject(FemaleTitan, Position{x, y}, FEMALE_TITAN_LIFE)
			} else if i < NB_BASIC_TITANS+4 {
				titan = NewObject(JawTitan, Position{x, y}, JAW_TITAN_LIFE)
			} else if i < NB_BASIC_TITANS+5 {
				titan = NewObject(ArmoredTitan, Position{x, y}, ARMORED_TITAN_LIFE)
			}
		}

		//Place l'humain dans des coordonnees aleatoires valides (i.e sans collisions) dans le village
		titans = PlaceTitan(titan, titans, H, W, dir)
	}
	return titans

}

func NewEnvironement(H int, W int) *Environment {
	objects := createStaticObjects(H, W)
	humans := createHumans(objects, Position{X: int(0.2*float32(W)) + CWall, Y: int(0.2*float32(H)) + CWall}, Position{X: int(0.8 * float32(W)), Y: 700 - HMaleVillager})
	titans := createTitans(H, W)
	// for _, titan := range titans {
	// 	fmt.Println(titan)
	// }
	merged_agents := make([]Object, len(titans)+len(humans))
	merged_agents = append(merged_agents, humans...)
	merged_agents = append(merged_agents, titans...)
	return &Environment{agents: merged_agents, objects: objects}
}

func (p *Environment) Objects() []Object {
	return p.objects
}

// A modifier quand le construteur d'agent sera pret
func (p *Environment) Agents() []Object {
	return p.agents
}

func (e *Environment) PerceivedObjects(topLeft Position, bottomRight Position) []Object {
	positions := make([]Object, 0)
	for _, obj := range e.objects {
		tl, br := obj.Hitbox()[0], obj.Hitbox()[1]
		if IntersectSquare(tl, br, topLeft, bottomRight) {
			positions = append(positions, obj)
		}
	}
	return positions
}

func (e *Environment) PerceivedAgents(topLeft Position, bottomRight Position, agtId Id) []AgentI {
	positions := make([]AgentI, 0)
	for _, agt := range e.Agents() {
		object := agt.Object()
		tl, br := object.Hitbox()[0], object.Hitbox()[1]
		if IntersectSquare(tl, br, topLeft, bottomRight) && agt.Id() != agtId {
			positions = append(positions, agt)
		}
	}
	return positions
}

func (e *Environment) Add(a AgentI) {
	e.agents = append(e.agents, a)
}
