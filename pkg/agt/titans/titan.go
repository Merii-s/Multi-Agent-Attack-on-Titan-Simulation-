package agt

import (
	pkg "AOT/pkg"
)

type TitanI interface {
	pkg.AgentI
	Attack()
}

type Titan struct {
	height int
}
