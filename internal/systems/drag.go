package systems

import (
	pxginput "github.com/timsims1717/pixel-go-input"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/object"
)

func DragSystem(in *pxginput.Input) {
	for _, result := range myecs.Manager.Query(myecs.IsDrag) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		drag, ok := result.Components[myecs.Drag].(*data.Dragger)
		if okO && ok {
			pos := in.World
			if drag.ViewPort != nil {
				pos = drag.ViewPort.Projected(pos)
			}
			obj.Pos = pos
		}
	}
}
