package systems

import (
	"github.com/faiface/pixel"
	pxginput "github.com/timsims1717/pixel-go-input"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/viewport"
)

func DragSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsDrag) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		vp, ok := result.Components[myecs.ViewPort].(*viewport.ViewPort)
		vec, okV := result.Components[myecs.Drag].(*pixel.Vec)
		if okO && ok && okV {
			pos := *vec
			if vp != nil {
				pos = vp.Projected(pos)
			}
			obj.Pos = pos
		}
	}
}

func ClickSystem(in *pxginput.Input) {
	for _, result := range myecs.Manager.Query(myecs.IsDrag) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		vp, ok := result.Components[myecs.ViewPort].(*viewport.ViewPort)
		click, ok := result.Components[myecs.Click].(*data.Funky)
		if okO && ok {
			pos := in.World
			if vp != nil {
				pos = vp.Projected(pos)
			}
			if obj.PointInside(pos) && in.Get("click").JustPressed() && click.Fn != nil {
				click.Fn()
			}
		}
	}
}
