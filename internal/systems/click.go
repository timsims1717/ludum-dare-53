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
		_, okD := result.Components[myecs.Drag].(*pixel.Vec)
		in, okI := result.Components[myecs.Input].(*pxginput.Input)
		if okO && ok && okD && okI {
			pos := in.World
			if vp != nil {
				pos = vp.Projected(pos)
				pos = vp.ConstrainR(pos, obj.Rect)
			}
			obj.Pos = pos
		}
	}
}

func ClickSystem() {
	for _, result := range myecs.Manager.Query(myecs.CanClick) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		vp, ok := result.Components[myecs.ViewPort].(*viewport.ViewPort)
		click, okC := result.Components[myecs.Click].(*data.Funky)
		in, okI := result.Components[myecs.Input].(*pxginput.Input)
		if okO && ok && okC && okI && !obj.Hide {
			pos := in.World
			if vp != nil {
				pos = vp.Projected(pos)
			}
			if in.Get("click").JustPressed() {
				if obj.PointInside(pos) && click.Fn != nil {
					in.Get("click").Consume()
					click.Fn()
				}
			}
		}
	}
}
