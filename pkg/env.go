package pkg

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
	coefsCoords = [][]float32{{0.29, 0.7}, {0.5, 0.55}, {0.4, 0.55}, {0.62, 0.7}}
	for i, coords := range coefsCoords {
		if i < 2 {
			obj = NewObject(BigHouse1, Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, 1000000000)
		} else {
			obj = NewObject(BigHouse2, Position{X: int(coords[0] * float32(W)), Y: int(coords[1] * float32(H))}, 1000000000)
		}
		objects = append(objects, *obj)
	}

	//Donjons
	obj = NewObject(Dungeon, Position{X: int(0.2*float32(W) + CWall), Y: int(0.2*float32(H) + CWall)}, 1000000000)
	objects = append(objects, *obj)
	obj = NewObject(Dungeon, Position{X: int(0.8*float32(W) - CWall - WDungeon/2), Y: int(0.2*float32(H) + CWall)}, 1000000000)
	objects = append(objects, *obj)

	return objects
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
