package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/world"
)

var (
	DraggingPiece *Factromino
)

type Factromino struct {
	Blocks              []*FactoryBlock
	NoRot               bool
	Object              *object.Object
	Entity              *ecs.Entity
	LastPos             pixel.Vec
	Moving              bool
	MyFactronimoType    constants.FactrominoType
	MyFactronimoVariant constants.FactrominoVariant
	Color               TColor
	MyTetronimoType     constants.TetronimoType
}

type FactoryBlock struct {
	Coords world.Coords
	Color  TColor
	Object *object.Object
	Entity *ecs.Entity
}

type FactoryPad struct {
	Tet    *Factromino
	Object *object.Object
	Entity *ecs.Entity
}

const (
	ConveyorLength = 5
	ConveyorSpeed  = 50.
	ConveyorHeight = 19.5 * MSize
)

type conveyor struct {
	Tets   [ConveyorLength]*Factromino
	Entity *ecs.Entity
	Slots  [ConveyorLength]pixel.Vec
}

func NewConveyor() {
	Conveyor = &conveyor{}
	Conveyor.Slots = [ConveyorLength]pixel.Vec{
		pixel.V(-53.*MSize, ConveyorHeight),
		pixel.V(-41.*MSize, ConveyorHeight),
		pixel.V(-29.*MSize, ConveyorHeight),
		pixel.V(-17.*MSize, ConveyorHeight),
		pixel.V(-5.*MSize, ConveyorHeight),
	}
}

var (
	FactoryPads  []*FactoryPad
	SouthPad     *FactoryPad
	SouthEastPad *FactoryPad
	EastPad      *FactoryPad
	NorthEastPad *FactoryPad
	NorthPad     *FactoryPad
	GarbagePad   *FactoryPad
	QueuePad     *FactoryPad
	FactoryFloor *factoryFloor
	Conveyor     *conveyor
)

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
