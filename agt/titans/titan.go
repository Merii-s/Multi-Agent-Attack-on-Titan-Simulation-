package titans

import (
	env "AOT/agt/env"
	"AOT/pkg/obj"
	types "AOT/pkg/types"
)

type TitanI interface {
	env.AgentI
	attack()
	regenerate()
}

type Titan struct {
	agentAttributes env.Agent
	regenRate       int
	attackObject    bool
	objectToAttack  *obj.Object
}

func NewTitan(id types.Id, tl types.Position, life int, r int, s int, spd int, v int, o types.ObjectName, regen int) *Titan {
	return &Titan{
		agentAttributes: *env.NewAgent(id, tl, life, r, s, spd, v, o),
		regenRate:       regen}
}

func (t *Titan) RegenRate() int {
	return t.regenRate
}

func (t *Titan) AttackObjectBool() bool { return t.attackObject }

func (t *Titan) SetAttackObject(b bool) { t.attackObject = b }

func (t *Titan) ObjectToAttack() obj.Object { return *t.objectToAttack }

func (t *Titan) SetObjectToAttack(o *obj.Object) { t.objectToAttack = o }

func (t *Titan) AttackObject(o obj.Object) { o.TakeDamage(t.agentAttributes.Strength()) }
