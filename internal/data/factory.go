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
	MyTetrominoType     constants.TetronimoType
}

func (f *Factromino) RefreshState() {
	if f.MyTetrominoType == constants.UndefinedTetronimoType {
		f.DetectTetrominoType()
	}
	if f.Color == 0 && len(f.Blocks) > 0 {
		f.Color = f.Blocks[0].Color
	}
}
func (f *Factromino) DetectTetrominoType() {
	if len(f.Blocks) == 4 {
		var originalCoords [4]world.Coords
		for i, block := range f.Blocks {
			originalCoords[i] = world.Coords{block.Coords.X, block.Coords.Y}
		}
		newCoords := Normalize(originalCoords)
		for i, kv := range constants.NormalizedTetronimos {
			if TetronimoCoordsEqual(i, newCoords) {
				f.MyTetrominoType = kv
				return
			}
		}
	}
	f.MyTetrominoType = constants.UndefinedTetronimoType
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
	ConveyorHeight = 20.5 * MSize
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
	Stats  *FactoryStats
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
		Stats:  newFactoryStats(),
	}
}

func FactoryLegal(coords world.Coords) bool {
	return coords.X >= 0 && coords.X < constants.FactoryWidth &&
		coords.Y >= 0 && coords.Y < constants.FactoryHeight
}
