package data

import (
	"github.com/bytearena/ecs"
	"timsims1717/ludum-dare-53/pkg/world"
)

var (
	DraggingPiece *ecs.Entity
)

type FacTetronimo struct {
	Blocks []*FactoryBlock
	NoRot  bool
}

type FactoryBlock struct {
	Coords world.Coords
	Color  TColor
	Entity *ecs.Entity
}
