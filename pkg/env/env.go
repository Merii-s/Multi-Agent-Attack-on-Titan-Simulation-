package env

import (
	agt "aot/agt"
	"sync"
)

type Environment struct {
	sync.RWMutex
	agents     []agt.Agent
	agentCount int
	in         uint64
	out        uint64
	noopCount  uint64
}
