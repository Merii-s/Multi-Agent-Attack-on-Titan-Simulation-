package agt

import (
	pkg "AOT/pkg"
	"sync"
)

type CivilianI interface {
	HumanI
	build()
	getFood()
}

type Civilian struct {
	attributes Human
	stopCh     chan struct{}
	syncChan   chan string
	mu         sync.Mutex
	pkg.BehaviorI
}

func NewCivilian(id pkg.Id, t pkg.Type, topLeft pkg.Position, hp int, reach int, strength int, speed int) *Civilian {
	atts := NewHuman(id, t, topLeft, hp, reach, strength, speed)
	return &Civilian{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
		BehaviorI:  &CivilianBehavior{},
	}
}

// Setter and getter methods for Civilian
func (c *Civilian) SyncChan() chan string {
	return c.syncChan
}

func (c *Civilian) StopCh() chan struct{} {
	return c.stopCh
}

// func (c *Civilian) Mu() sync.Mutex {
// 	return c.mu
// }

// methods for civilian
func (c *Civilian) Percept(e *pkg.Environment) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.BehaviorI.Percept(e)
}

func (c *Civilian) Deliberate() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.BehaviorI.Deliberate()

}

func (c *Civilian) Act(e *pkg.Environment) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.BehaviorI.Act(e)
}

func (c *Civilian) Id() pkg.Id {
	return c.attributes.agentAttributes.Id()
}

func (*Civilian) Start() {

}

func (c *Civilian) move() {
	// TODO : Move randomly or towards a target --> not only in a straight line (top right here)
	new_X_pos := c.attributes.agentAttributes.Pos().X + c.attributes.agentAttributes.Speed()
	new_Y_pos := c.attributes.agentAttributes.Pos().Y + c.attributes.agentAttributes.Speed()
	new_pos := pkg.Position{X: new_X_pos, Y: new_Y_pos}
	c.attributes.agentAttributes.SetPos(new_pos)
}

func (c *Civilian) eat() {

}

func (*Civilian) sleep() {
	//time.Sleep(?)
}

func (c *Civilian) build() {

}

func (c *Civilian) getFood()

// Define the behavior struct of the Soldier :
type CivilianBehavior struct {
}

func (cb *CivilianBehavior) Percept(e *pkg.Environment) {

}

func (cb *CivilianBehavior) Deliberate() {

}

func (cb *CivilianBehavior) Act(e *pkg.Environment) {

}
