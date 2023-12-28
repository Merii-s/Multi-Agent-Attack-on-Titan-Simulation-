package env

import (
	obj "AOT/pkg/obj"
	params "AOT/pkg/parameters"
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"
	"sync"
)

type Environment struct {
	sync.RWMutex
	Agts []AgentI
	Objs []obj.Object
	// Day/Night cycle
	day bool
}

func (p *Environment) Objects() []obj.Object {
	return p.Objs
}

// A modifier quand le construteur d'agent sera pret
func (p *Environment) Agents() []AgentI {
	return p.Agts
}

func (e *Environment) PerceivedObjects(topLeft types.Position, bottomRight types.Position) []obj.Object {
	positions := make([]obj.Object, 0)
	for _, obj := range e.Objs {
		tl, br := obj.Hitbox()[0], obj.Hitbox()[1]
		if utils.IntersectSquare(tl, br, topLeft, bottomRight) {
			positions = append(positions, obj)
		}
	}
	return positions
}

func (e *Environment) PerceivedAgents(topLeft types.Position, bottomRight types.Position, agtId types.Id) []AgentI {
	positions := make([]AgentI, 0)
	for _, agt := range e.Agents() {
		object := agt.Object()
		tl, br := object.Hitbox()[0], object.Hitbox()[1]
		if utils.IntersectSquare(tl, br, topLeft, bottomRight) && agt.Id() != agtId {
			positions = append(positions, agt)
		}
	}
	return positions
}

func (e *Environment) Add(a AgentI) {
	e.Agts = append(e.Agts, a)
}

func CreateStaticObjects(H int, W int) []obj.Object {
	objects := make([]obj.Object, 0)
	var object *obj.Object
	nb_objects := 0

	//Grass Sprites
	nW := W / params.CGrass
	nH := H / params.CGrass
	for i := 0; i < nH; i++ {
		for j := 0; j < nW; j++ {
			object = obj.NewObject(types.Grass, types.Position{X: j * params.CGrass, Y: i * params.CGrass}, params.GRASS_LIFE)
			objects = append(objects, *object)
			nb_objects = nb_objects + 1
		}
	}

	//Wall Sprites
	wall_Tl_X := int(0.2 * float32(W))
	wall_Tl_Y := int(0.2 * float32(H))
	wWall := int(0.6 * float32(W))       //Largeur du mur
	hWall := int((6. / 7.) * float32(H)) //Hauteur du mur
	nW = wWall / params.CWall            //Nombre de sprites en largeur
	nH = hWall / params.CWall            //Nombre de sprites en hauteur

	//Mur du Nord, Horizontal
	for i := 0; i < nW+1; i++ {
		object = obj.NewObject(types.Wall, types.Position{X: wall_Tl_X + i*params.CWall, Y: wall_Tl_Y}, params.WALL_LIFE)
		objects = append(objects, *object)
		nb_objects = nb_objects + 1
	}
	//Murs cotes, Verticaux
	for i := 1; i < nH; i++ {
		object = obj.NewObject(types.Wall, types.Position{X: wall_Tl_X, Y: wall_Tl_Y + i*params.CWall}, params.WALL_LIFE)
		objects = append(objects, *object)
		object = obj.NewObject(types.Wall, types.Position{X: wall_Tl_X + wWall, Y: wall_Tl_Y + i*params.CWall}, params.WALL_LIFE)
		objects = append(objects, *object)
		nb_objects = nb_objects + 2
	}

	//Champs
	nW = ((wWall - 2*params.CWall) / 2) / params.WField //Nombre de sprites en largeur
	for i := 0; i < nW*2; i++ {
		var x, y int
		if i < 7 {
			x = wall_Tl_X + params.CWall + int(wWall/4) + i*params.WField
			y = wall_Tl_Y + params.CWall + int(hWall/10)
		} else {
			x = wall_Tl_X + params.CWall + int(wWall/4) + (i-7)*params.WField
			y = wall_Tl_Y + params.CWall + int(hWall/10) + params.HField*2
		}
		object = obj.NewObject(types.Field, types.Position{X: x, Y: y}, params.FIELD_LIFE)
		objects = append(objects, *object)
		nb_objects = nb_objects + 1
	}

	//Petites maisons
	coefsCoords := [][]float32{{0.29, 0.4}, {1 - 0.29, 0.4}, {0.29, 0.85}, {1 - 0.29, 0.85}, {0.29, 0.55}, {1 - 0.29, 0.65}, {0.5, 0.85}}
	for _, coords := range coefsCoords {

		object = obj.NewObject(types.SmallHouse, types.Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, params.SMALL_HOUSE_LIFE)
		objects = append(objects, *object)
		nb_objects = nb_objects + 1
	}

	//Grandes maisons 1 et 2
	coefsCoords = [][]float32{{0.29, 0.7}, {0.5, 0.55}, {0.60, 0.7}}
	for _, coords := range coefsCoords {
		object = obj.NewObject(types.BigHouse1, types.Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, params.BIG_HOUSE_LIFE)
		objects = append(objects, *object)
		nb_objects = nb_objects + 1
	}

	//Donjons
	object = obj.NewObject(types.Dungeon, types.Position{X: int(0.2*float32(W) + params.CWall), Y: int(0.2*float32(H) + params.CWall)}, params.DUNGEON_LIFE)
	objects = append(objects, *object)
	object = obj.NewObject(types.Dungeon, types.Position{X: int(0.8*float32(W) - params.CWall - params.WDungeon/2), Y: int(0.2*float32(H) + params.CWall)}, params.DUNGEON_LIFE)
	objects = append(objects, *object)

	object = obj.NewObject(types.ColossalTitan, types.Position{X: 640, Y: 350}, params.COLOSSAL_TITAN_LIFE)
	objects = append(objects, *object)
	nb_objects = nb_objects + 3

	return objects
}

func MoveColossal(e *Environment, c chan *Environment, wg *sync.WaitGroup) {
	var ind int
	coords := [][]int{{250, 250}, {750, 250}, {750, 450}, {250, 450}}
	for i, _ := range e.Objs {
		if e.Objs[i].Name() == types.ColossalTitan {
			ind = i
			break
		}
	}

	for {
		for _, pos := range coords {
			wg.Add(1)
			go func(pos []int) {
				e.Objs[ind].SetPosition(types.Position{X: pos[0], Y: pos[1]})
				c <- e
				wg.Done()
			}(pos)
			wg.Wait()
		}
	}
}

// function that gets the Agent with the specified selected id
// func (e *pkg.Environment) GetAgent(id pkg.Id) pkg.Agent {
// 	// For example, if you have a map of agents in the environment:
// 	agent, exists := e.agents[id]
// 	if !exists {
// 		fmt.Println("Agent not found in the environment")
// 	}
// 	return agent
// }
