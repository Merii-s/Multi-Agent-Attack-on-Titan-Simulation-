package pkg

func NewEnvironement(H int, W int) *Environment {
	//env := &pkg.Environment{objects:new(pkg.Objects), agents:new(pkg.Agents), screenH:H, screenW:W}

	objects := make([]Object, 0)

	//Grass Sprites
	nW := W / CGrass
	nH := H / CGrass
	for i := 0; i < nH; i++ {
		for j := 0; j < nW; j++ {
			obj := NewObject(Grass, Position{X: j * CGrass, Y: i * CGrass}, 1000000000)
			objects = append(objects, *obj)
		}
	}

	//Wall Sprites

	wall_Tl_X := int(0.2 * float32(W))
	wall_Tl_Y := int(0.2 * float32(H))

	wWall := int(0.6 * float32(W))
	hWall := int((6. / 7.) * float32(H))

	nW = wWall / CWall
	nH = hWall / CWall

	//Mur du Nord, Horizontal
	for i := 0; i < nW+1; i++ {
		obj := NewObject(Wall, Position{X: wall_Tl_X + i*CWall, Y: wall_Tl_Y}, 1000000000)
		objects = append(objects, *obj)
	}

	//Murs cotes, Verticaux
	for i := 1; i < nH; i++ {
		obj1 := NewObject(Wall, Position{X: wall_Tl_X, Y: wall_Tl_Y + i*CWall}, 1000000000)
		objects = append(objects, *obj1)
		obj2 := NewObject(Wall, Position{X: wall_Tl_X + wWall, Y: wall_Tl_Y + i*CWall}, 1000000000)
		objects = append(objects, *obj2)
	}

	return &Environment{objects: objects}
}

func (p *Environment) Objects() []Object {
	return p.objects
}

func (p *Environment) Agents() []AgentI {
	return p.agents
}
