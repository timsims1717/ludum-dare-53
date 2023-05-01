package systems

import (
	pxginput "github.com/timsims1717/pixel-go-input"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/pkg/viewport"
)

func MenuSystem(in *pxginput.Input) {
	for _, item := range data.MenuItems {
		if !item.Text.Obj.Hide {
			item.Text.Obj.Update()
			if item.Text.Obj.PointInside(data.StickyViewport.Projected(in.World)) {
				item.Text.SetSize(item.OrigSize * 1.1)
				if in.Get("click").JustPressed() {
					in.Get("click").Consume()
					item.Click()
				}
			} else {
				item.Text.SetSize(item.OrigSize)
			}
		} else {
			item.Text.SetSize(item.OrigSize)
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
