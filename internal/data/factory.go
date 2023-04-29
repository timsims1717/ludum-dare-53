package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/world"
)

var (
	DraggingPiece *FacTetronimo
)

type FacTetronimo struct {
	Blocks  []*FactoryBlock
	NoRot   bool
	Object  *object.Object
	Entity  *ecs.Entity
	LastPos pixel.Vec
}

type FactoryBlock struct {
	Coords world.Coords
	Color  TColor
	Object *object.Object
	Entity *ecs.Entity
}
