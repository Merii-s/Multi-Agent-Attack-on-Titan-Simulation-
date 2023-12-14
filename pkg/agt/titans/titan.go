package agt

import (
	pkg "aot/pkg"
)

type TitanI interface {
	pkg.AgentI
	Attack()
}

type Titan struct {
	height int
}
