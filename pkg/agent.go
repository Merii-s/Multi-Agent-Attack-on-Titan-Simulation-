package pkg

type Agent struct {
	id            Id
	type_agt      Type
	pos           Position
	reach         int
	strength      int
	speed         int
	hp            int
	maxHP         int
	vision        int
	seenPositions map[Position]ObjectName
	cantSeeBehind []ObjectName
}

func NewAgent(id Id, t Type, p Position, hp int, r int, s int, spd int) *Agent {
	return &Agent{id: id, type_agt: t, pos: p, hp: hp, reach: r, strength: s, speed: spd, maxHP: hp}
}

func (t *Agent) Id() Id {
	return t.id
}

func (t *Agent) Pos() Position {
	return t.pos
}

func (t *Agent) Reach() int {
	return t.reach
}

func (t *Agent) Strength() int {
	return t.strength
}

func (t *Agent) Speed() int {
	return t.speed
}

func (t *Agent) MaxHP() int {
	return t.maxHP
}

func (t *Agent) Hp() int {
	return t.hp
}

func (t *Agent) SetPos(pos Position) { t.pos = pos }

func (t *Agent) SetHp(hp int) { t.hp = hp }

func (t *Agent) Vision() int { return t.vision }

func (t *Agent) SeenPositions() map[Position]ObjectName { return t.seenPositions }

func (t *Agent) CantSeeBehind() []ObjectName { return t.cantSeeBehind }
