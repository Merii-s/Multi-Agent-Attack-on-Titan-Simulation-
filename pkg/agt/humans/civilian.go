package agt

import (
	pkg "AOT/pkg"
	"sync"
	"time"
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
	behavior   pkg.BehaviorI
}

func NewCivilian(id pkg.Id, tl pkg.Position, life int, reach int, strength int, speed int, vision int, obj pkg.ObjectName) *Civilian {
	atts := NewHuman(id, tl, life, reach, strength, speed, vision, obj)
	c := &Civilian{
		attributes: *atts,
		stopCh:     make(chan struct{}),
		syncChan:   make(chan string),
		mu:         sync.Mutex{},
	}
	behavior := &CivilianBehavior{c: c}
	c.SetBehavior(behavior)
	return c
}

// Setter and getter methods for Civilian
func (c *Civilian) SyncChan() chan string {
	return c.syncChan
}

func (c *Civilian) StopCh() chan struct{} {
	return c.stopCh
}

func (c *Civilian) Behavior() *pkg.BehaviorI {
	return &c.behavior
}

func (c *Civilian) SetBehavior(b pkg.BehaviorI) {
	c.behavior = b
}

// methods for civilian
func (c *Civilian) Percept(e *pkg.Environment) {
	c.behavior.Percept(e)
}

func (c *Civilian) Deliberate() {
	c.behavior.Deliberate()

}

func (c *Civilian) Act(e *pkg.Environment) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.behavior.Act(e)
}

func (c *Civilian) Id() pkg.Id {
	return c.attributes.agentAttributes.Id()
}

func (c *Civilian) Start(e *pkg.Environment) {
	// launch the agent goroutine Percept-Deliberate-Act cycle
	go func() {
		for {
			println("Civilian Start")
			c.behavior.Percept(e)
			time.Sleep(100 * time.Millisecond)
			c.behavior.Deliberate()
			c.behavior.Act(e)
		}
	}()
}

func (c *Civilian) move(pos pkg.Position) {
	// TODO : Move randomly or towards a target --> not only in a straight line (top right here)
	c.attributes.agentAttributes.SetPos(pos)
}

func (c *Civilian) eat() {

}

func (*Civilian) sleep() {
	//time.Sleep(?)
}

func (c *Civilian) build() {

}

func (c *Civilian) getFood()

// Define the behavior struct of the Civilian :
type CivilianBehavior struct {
	c *Civilian
}

func (cb *CivilianBehavior) Percept(e *pkg.Environment) {

}

func (cb *CivilianBehavior) Deliberate() {

}

func (cb *CivilianBehavior) Act(e *pkg.Environment) {

}
