package systems

import (
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/pkg/viewport"
)

func MenuSystem(world pixel.Vec, click bool) {
	for _, item := range data.MenuItems {
		if !item.Text.Obj.Hide {
			item.Text.Obj.Update()
			if item.Text.Obj.PointInside(data.StickyViewport.Projected(world)) {
				item.Text.SetSize(item.OrigSize * 1.1)
				if click {
					item.Click()
				}
			} else {
				item.Text.SetSize(item.OrigSize)
			}
		}
	}
}

func DrawMenuSystem(vp *viewport.ViewPort) {
	for _, item := range data.MenuItems {
		if !item.Text.Obj.Hide {
			item.Text.Draw(vp.Canvas)
		}
	}
}
