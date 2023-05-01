package data

import (
	"github.com/bcvery1/tilepix"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/reanimator"
)

const MSize = 16.

var (
	FactoryMap *tilepix.Map
	FactoryMat pixel.Matrix
)

var (
	StickyObj     *object.Object
	StickyNote    *pixel.Sprite
	FloorSection  []*img.Sprite
	PadSection    []*img.Sprite
	WallSection   []*img.Sprite
	DoorSection   []*img.Sprite
	SideSection   []*img.Sprite
	CornerSection []*img.Sprite
	SideDSection  []*img.Sprite
	BlockSpot     []*img.Sprite

	ConveyorBase []*img.Sprite
	BeltSize     = 53

	ConvLeftEdge  []*reanimator.Tree
	ConvRightEdge []*reanimator.Tree
	ConvMiddle    []*reanimator.Tree

	TruckWidth  = 10
	TruckHeight = 10

	TopTruck []*img.Sprite
	MidTruck []*img.Sprite
	BotTruck []*img.Sprite

	TV   []*img.Sprite
	Bulb []*img.Sprite

	TVShapes []*img.Sprite
)
