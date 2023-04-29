package menus

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/typeface"
)

type Item struct {
	Key  string
	Raw  string
	Text *typeface.Text
	Obj  *object.Object

	clickFn   func()
	leftFn    func()
	rightFn   func()
	hoverFn   func()
	unHoverFn func()

	Hovered  bool
	Disabled bool
	NoHover  bool
	Ignore   bool
	NoDraw   bool
	hovered  bool
	disabled bool
	noShowT  bool
}

func (menu *Menu) NewItem(key, raw string, pos pixel.Vec, align typeface.Alignment) *Item {
	tex := typeface.New(nil, menu.AtlasKey, align, 1.5, 1., 0., 0.)
	tex.SetColor(colornames.White)
	tex.SetText(raw)
	return &Item{
		Key:  key,
		Raw:  raw,
		Text: tex,
	}
}
