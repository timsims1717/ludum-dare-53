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

func BuildFactoryPads() {
	// Import Pads
	for i := 0; i < 5; i++ {
		pad := &data.FactoryPad{}
		obj := object.New()
		if i < 2 {
			obj.Pos.Y = data.MSize * 21.
		} else if i == 2 {
			obj.Pos.Y = data.MSize * 7.
		} else {
			obj.Pos.Y = data.MSize * -7.
		}
		if i > 0 && i < 4 {
			obj.Pos.X = data.MSize * 25.
		} else {
			obj.Pos.X = data.MSize * 9.
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
				// todo: add hover shine
			}))
		pad.Entity = e
		FactoryBGEntities = append(FactoryBGEntities, e)
		data.FactoryPads = append(data.FactoryPads, pad)
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
			// todo: add hover shine
		}))
	data.GarbagePad.Entity = e
	FactoryBGEntities = append(FactoryBGEntities, e)
	// queue pad
	data.QueuePad = &data.FactoryPad{}
	objQ := object.New()
	objQ.Pos.Y = 19.5 * data.MSize
	objQ.Pos.X = -5. * data.MSize
	objQ.Layer = 11
	objQ.Rect = pixel.R(0., 0., constants.FactoryTile*2.8, world.TileSize*2.8)
	data.QueuePad.Object = objQ
	sprQ := img.NewSprite("green_circle", constants.BlockKey)
	eQ := myecs.Manager.NewEntity()
	eQ.AddComponent(myecs.Object, data.QueuePad.Object).
		AddComponent(myecs.Drawable, sprQ).
		AddComponent(myecs.Input, factoryInput).
		AddComponent(myecs.ViewPort, data.FactoryViewport).
		AddComponent(myecs.Click, data.NewFn(AddToQueuePad)).
		AddComponent(myecs.Update, data.NewFn(func() {
			// todo: add hover shine
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
