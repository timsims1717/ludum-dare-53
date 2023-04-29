package debug

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/pkg/timing"
	"timsims1717/ludum-dare-53/pkg/typeface"
)

var (
	fpsText     *typeface.Text
	versionText *typeface.Text
)

func InitializeFPS(v *pixel.Vec) {
	col := colornames.Aliceblue
	col.A = 90
	fpsText = typeface.New(v, "basic", typeface.NewAlign(typeface.Left, typeface.Bottom), 1.0, 0.5, 0., 0.)
	versionText = typeface.New(v, "basic", typeface.NewAlign(typeface.Right, typeface.Bottom), 1.0, 0.5, 0., 0.)
}

func DrawFPS(win *pixelgl.Window) {
	fpsText.SetText(fmt.Sprintf("FPS: %s", timing.FPS))
	fpsText.Obj.Pos = pixel.V(win.Bounds().W()*-0.5+2., win.Bounds().H()*-0.5+2)
	fpsText.Obj.Update()
	fpsText.Draw(win)
	versionText.SetText(fmt.Sprintf("%d.%d.%d", constants.Release, constants.Version, constants.Build))
	versionText.Obj.Pos = pixel.V(win.Bounds().W()*0.5-2., win.Bounds().H()*-0.5+2)
	versionText.Obj.Update()
	versionText.Draw(win)
}
