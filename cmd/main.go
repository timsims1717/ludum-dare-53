package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/loading"
	"timsims1717/ludum-dare-53/internal/state"
	"timsims1717/ludum-dare-53/pkg/debug"
	"timsims1717/ludum-dare-53/pkg/options"
	"timsims1717/ludum-dare-53/pkg/timing"
	"timsims1717/ludum-dare-53/pkg/viewport"
	"timsims1717/ludum-dare-53/pkg/world"
)

func run() {
	world.SetTileSize(constants.TileSize)
	cfg := pixelgl.WindowConfig{
		Title:     constants.Title,
		Bounds:    pixel.R(0, 0, 1600, 900),
		VSync:     true,
		Invisible: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(false)

	viewport.MainCamera = viewport.New(win.Canvas())
	viewport.MainCamera.SetRect(pixel.R(0, 0, 1600, 900))
	viewport.MainCamera.PortPos = pixel.V(0., 0.)

	options.VSync = true

	loading.LoadImg()

	debug.Initialize(&viewport.MainCamera.PostCamPos, &viewport.MainCamera.PostCamPos)

	win.Show()

	timing.Reset()
	for !win.Closed() {
		state.Update(win)
		state.Draw(win)
	}
}

// fire the run function (the real main function)
func main() {
	pixelgl.Run(run)
}
