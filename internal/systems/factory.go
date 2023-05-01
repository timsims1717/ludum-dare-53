package systems

import (
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/object"
)

func FactoryBlockSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsBlock) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		tet, ok := result.Components[myecs.Block].(*data.FacTetromino)
		if okO && ok {
			for _, block := range tet.Blocks {
				block.Object.Layer = obj.Layer
			}
		}
	}
}
