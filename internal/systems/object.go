package systems

import (
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/object"
)

func ObjectSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsObject) {
		if obj, ok := result.Components[myecs.Object].(*object.Object); ok {
			if obj.Kill {
				myecs.Manager.DisposeEntity(result)
			} else {
				obj.Update()
			}
		}
	}
}

func ParentSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasParent) {
		tran, okT := result.Components[myecs.Object].(*object.Object)
		parent, okP := result.Components[myecs.Parent].(*object.Object)
		if okT && okP {
			if parent.Kill {
				myecs.Manager.DisposeEntity(result)
			} else {
				tran.Pos = parent.Pos
				tran.Hide = parent.Hide
			}
		}
	}
}
