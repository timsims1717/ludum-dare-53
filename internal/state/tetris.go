package state

import (
	"github.com/bytearena/ecs"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/world"
)

var (
	TetrisBGEntities []*ecs.Entity
)

func BuildTetrisBG() {
	for y := 0; y < constants.TetrisHeight; y++ {
		for x := 0; x < constants.TetrisWidth; x++ {
			obj := object.New()
			obj.Pos = world.MapToWorld(world.Coords{X: x, Y: y})
			obj.Layer = 1
			spr := img.NewSprite("t_bg", constants.BlockKey)
			e := myecs.Manager.NewEntity()
			e.AddComponent(myecs.Object, obj).
				AddComponent(myecs.Drawable, spr)
			TetrisBGEntities = append(TetrisBGEntities, e)
		}
	}
}
