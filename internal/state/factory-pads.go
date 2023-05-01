package state

import (
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/internal/systems"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/world"
)

var (
	PadX1 = data.MSize * 9.
	PadX2 = data.MSize * 25.
	PadY1 = data.MSize * 21.
	PadY2 = data.MSize * 7.
	PadY3 = data.MSize * -7.
)

func BuildFactoryPads() {
	// Import Pads
	for i := 0; i < 5; i++ {
		pad := &data.FactoryPad{}
		obj := object.New()
		if i < 2 {
			obj.Pos.Y = PadY1
		} else if i == 2 {
			obj.Pos.Y = PadY2
		} else {
			obj.Pos.Y = PadY3
		}
		if i > 0 && i < 4 {
			obj.Pos.X = PadX2
		} else {
			obj.Pos.X = PadX1
		}
		obj.Layer = 10
		obj.Rect = pixel.R(0., 0., constants.FactoryTile*2.8, world.TileSize*2.)
		pad.Object = obj
		e := myecs.Manager.NewEntity()
		e.AddComponent(myecs.Object, pad.Object).
			AddComponent(myecs.Drawable, data.PadSection).
			AddComponent(myecs.Input, factoryInput).
			AddComponent(myecs.ViewPort, data.FactoryViewport).
			AddComponent(myecs.Click, data.NewFn(func() {
				if pad.Tet != nil && data.DraggingPiece == nil {
					data.DraggingPiece = pad.Tet
					data.DraggingPiece.Entity.AddComponent(myecs.Drag, &factoryInput.World)
					data.DraggingPiece.Object.Layer = 20
					pad.Tet = nil
				} else if pad.Tet == nil && data.DraggingPiece != nil {
					pad.Tet = data.DraggingPiece
					data.DraggingPiece.Entity.RemoveComponent(myecs.Drag)
					data.DraggingPiece.Object.Layer = 12
					data.DraggingPiece.Object.Pos = pad.Object.Pos
					data.DraggingPiece = nil
				}
			})).
			AddComponent(myecs.Update, data.NewFn(func() {
				if ((data.DraggingPiece != nil && pad.Tet == nil) ||
					(data.DraggingPiece == nil && pad.Tet != nil)) &&
					obj.PointInside(data.FactoryViewport.Projected(factoryInput.World)) {
					PadHighlight(obj.Pos)
				}
			}))
		pad.Entity = e
		FactoryBGEntities = append(FactoryBGEntities, e)
		data.FactoryPads = append(data.FactoryPads, pad)
		switch i {
		case 0:
			data.NorthPad = pad
		case 1:
			data.NorthEastPad = pad
		case 2:
			data.EastPad = pad
		case 3:
			data.SouthEastPad = pad
		case 4:
			data.SouthPad = pad
		}
	}
	// garbage pad
	data.GarbagePad = &data.FactoryPad{}
	obj := object.New()
	obj.Pos.Y = data.MSize * -9.
	obj.Pos.X = data.MSize * -6.5
	obj.Layer = 11
	obj.Rect = pixel.R(0., 0., constants.FactoryTile*2.8, world.TileSize*2.8)
	data.GarbagePad.Object = obj
	spr := img.NewSprite("bin", constants.FactoryKey)
	e := myecs.Manager.NewEntity()
	e.AddComponent(myecs.Object, data.GarbagePad.Object).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Input, factoryInput).
		AddComponent(myecs.ViewPort, data.FactoryViewport).
		AddComponent(myecs.Click, data.NewFn(func() {
			if data.DraggingPiece != nil {
				for _, block := range data.DraggingPiece.Blocks {
					myecs.Manager.DisposeEntity(block.Entity)
				}
				myecs.Manager.DisposeEntity(data.DraggingPiece.Entity)
				data.DraggingPiece = nil
			}
		})).
		AddComponent(myecs.Update, data.NewFn(func() {
			if data.DraggingPiece != nil &&
				obj.PointInside(data.FactoryViewport.Projected(factoryInput.World)) {
				PadHighlight(obj.Pos)
			}
		}))
	data.GarbagePad.Entity = e
	FactoryBGEntities = append(FactoryBGEntities, e)
	// queue pad
	data.QueuePad = &data.FactoryPad{}
	objQ := object.New()
	objQ.Pos.Y = data.ConveyorHeight
	objQ.Pos.X = -5. * data.MSize
	objQ.Layer = 11
	objQ.Rect = pixel.R(0., 0., constants.FactoryTile*2.8, world.TileSize*2.8)
	data.QueuePad.Object = objQ
	eQ := myecs.Manager.NewEntity()
	eQ.AddComponent(myecs.Object, data.QueuePad.Object).
		AddComponent(myecs.Input, factoryInput).
		AddComponent(myecs.ViewPort, data.FactoryViewport).
		AddComponent(myecs.Click, data.NewFn(AddToQueuePad)).
		AddComponent(myecs.Update, data.NewFn(func() {
			if data.DraggingPiece != nil && len(data.DraggingPiece.Blocks) == 4 &&
				objQ.PointInside(data.FactoryViewport.Projected(factoryInput.World)) {
				PadHighlight(objQ.Pos)
			}
		}))
	data.QueuePad.Entity = eQ
	FactoryBGEntities = append(FactoryBGEntities, eQ)
}

func AddToQueuePad() {
	if data.DraggingPiece != nil {
		if len(data.DraggingPiece.Blocks) == 4 {
			if data.QueuePad.Tet == nil {
				data.QueuePad.Tet = data.DraggingPiece
				data.DraggingPiece.Entity.RemoveComponent(myecs.Drag)
				data.DraggingPiece.Object.Layer = 12
				data.DraggingPiece.Object.Pos = data.QueuePad.Object.Pos
				data.Conveyor.Tets[data.ConveyorLength-1] = data.DraggingPiece
				data.DraggingPiece = nil
			}
		}
	}
}

func QueuePadUpdate() {

}

func AddToQueue() {
	if data.DraggingPiece != nil {
		if len(data.DraggingPiece.Blocks) == 4 {
			systems.FactoTet(data.DraggingPiece)

			for _, block := range data.DraggingPiece.Blocks {
				myecs.Manager.DisposeEntity(block.Entity)
			}
			myecs.Manager.DisposeEntity(data.DraggingPiece.Entity)
			data.DraggingPiece = nil
		}
	}
}

func PadHighlight(pos pixel.Vec) {
	obj := object.New()
	obj.Pos = pos
	obj.Layer = 19
	spr := img.NewSprite("highlight", constants.BlockKey)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Temp, myecs.ClearFlag(true))
}
