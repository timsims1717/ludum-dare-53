package data

import (
	"github.com/bcvery1/tilepix"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/reanimator"
)

const MSize = 16.

var (
	FactoryMap *tilepix.Map
	FactoryMat pixel.Matrix
)

var (
	FloorSection []*img.Sprite
	PadSection   []*img.Sprite
	WallSection  []*img.Sprite
	DoorSection  []*img.Sprite
	SideSection  []*img.Sprite
	SideDSection []*img.Sprite
	BlockSpot    []*img.Sprite

	ConveyorBase []*img.Sprite
	BeltSize     = 53

	ConvLeftEdge  []*reanimator.Tree
	ConvRightEdge []*reanimator.Tree
	ConvMiddle    []*reanimator.Tree
)