package state

import (
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/viewport"
	"timsims1717/ludum-dare-53/pkg/world"
)

func BuildFactoryPads(vp *viewport.ViewPort) {
	// Import Pads
	for i := 0; i < 5; i++ {
		pad := &data.FactoryPad{}
		obj := object.New()
		if i < 3 {
			obj.Pos.Y = 400.
		} else if i == 3 {
			obj.Pos.Y = 100.
		} else if i == 4 {
			obj.Pos.Y = -200.
		}
		if i > 1 {
			obj.Pos.X = -150.
		} else if i == 1 {
			obj.Pos.X = 75.
		} else if i == 0 {
			obj.Pos.X = 300.
		}
		obj.Layer = 11
		obj.Rect = pixel.R(0., 0., constants.FactoryTile*2.8, world.TileSize*2.8)
		pad.Object = obj
		spr := img.NewSprite("yellow_circle", constants.BlockKey)
		e := myecs.Manager.NewEntity()
		e.AddComponent(myecs.Object, pad.Object).
			AddComponent(myecs.Drawable, spr).
			AddComponent(myecs.Input, factoryInput).
			AddComponent(myecs.ViewPort, vp).
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
		FactoryBGEntities = append(FactoryBGEntities, e)
		data.FactoryPads = append(data.FactoryPads, pad)
	}
}
