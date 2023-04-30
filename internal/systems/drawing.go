package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/reanimator"
)

func AnimationSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasAnimation) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		anim, ok := result.Components[myecs.Animation].(*reanimator.Tree)
		if okO && ok && !obj.Hide {
			anim.Update()
		}
	}
}

func DrawSystem(win *pixelgl.Window, layer int) {
	count := 0
	for _, result := range myecs.Manager.Query(myecs.IsDrawable) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		if okO && obj.Layer == layer {
			draw := result.Components[myecs.Drawable]
			if draw == nil {
				continue
			} else if draws, okD := draw.([]*img.Sprite); okD {
				for _, d := range draws {
					DrawThing(d, obj, win)
					count++
				}
			} else {
				DrawThing(draw, obj, win)
				count++
			}
		}
	}
	//debug.AddText(fmt.Sprintf("Layer %d: %d entities", layer, count))
}

func DrawThing(draw interface{}, obj *object.Object, win *pixelgl.Window) {
	if spr, ok0 := draw.(*pixel.Sprite); ok0 {
		spr.Draw(win, obj.Mat)
	} else if sprH, ok1 := draw.(*img.Sprite); ok1 {
		if batch, okB := img.Batchers[sprH.Batch]; okB {
			batch.DrawSpriteColor(sprH.Key, obj.Mat.Moved(sprH.Offset), sprH.Color)
		}
	} else if anim, ok2 := draw.(*reanimator.Tree); ok2 {
		sprA := anim.CurrentSprite()
		if batch, okB := img.Batchers[sprA.Batch]; okB {
			batch.DrawSpriteColor(sprA.Key, obj.Mat.Moved(sprA.Offset), sprA.Color)
		}
	}
}
