package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
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

type FactoryPad struct {
	Tet    *FacTetronimo
	Object *object.Object
	Entity *ecs.Entity
}

var (
	FactoryPads []*FactoryPad
)

var FactoryFloor *factoryFloor

type factoryFloor struct {
	Blocks [constants.FactoryHeight][constants.FactoryWidth]*FactoryBlock
	Object *object.Object
	Entity *ecs.Entity
}

func (f *factoryFloor) Get(c world.Coords) *FactoryBlock {
	return f.Blocks[c.Y][c.X]
}

func (f *factoryFloor) Set(c world.Coords, b *FactoryBlock) {
	f.Blocks[c.Y][c.X] = b
}

func NewFactoryFloor() {
	FactoryFloor = &factoryFloor{
		Blocks: [constants.FactoryHeight][constants.FactoryWidth]*FactoryBlock{},
	}
}

func FactoryLegal(coords world.Coords) bool {
	return coords.X >= 0 && coords.X < constants.FactoryWidth &&
		coords.Y >= 0 && coords.Y < constants.FactoryHeight
}
