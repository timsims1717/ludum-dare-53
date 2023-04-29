package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/world"
)

func CreateFactoryBlock(pos pixel.Vec, col data.TColor) *data.FactoryBlock {
	block := &data.FactoryBlock{
		Coords: world.Origin,
		Color:  col,
	}
	obj := object.New().WithID("factory-block")
	obj.Pos = pos
	obj.Layer = 13
	spr := img.NewSprite(fmt.Sprintf("%s_f", col.String()), constants.BlockKey)
	block.Entity = myecs.Manager.NewEntity()
	block.Entity.
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr)
	return block
}
