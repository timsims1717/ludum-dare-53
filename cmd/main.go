package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/loading"
	"timsims1717/ludum-dare-53/internal/state"
	"timsims1717/ludum-dare-53/pkg/debug"
	"timsims1717/ludum-dare-53/pkg/options"
	"timsims1717/ludum-dare-53/pkg/sfx"
	"timsims1717/ludum-dare-53/pkg/timing"
	"timsims1717/ludum-dare-53/pkg/typeface"
	"timsims1717/ludum-dare-53/pkg/viewport"
	"timsims1717/ludum-dare-53/pkg/world"
)

func run() {
	world.SetTileSize(constants.TileSize)
	cfg := pixelgl.WindowConfig{
		Title:     constants.RandomTitle(),
		Bounds:    pixel.R(0, 0, 1600, 900),
		VSync:     true,
		Invisible: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(false)
	win.SetCursorVisible(false)

	viewport.MainCamera = viewport.New(win.Canvas())
	viewport.MainCamera.SetRect(pixel.R(0, 0, 1600, 900))
	viewport.MainCamera.PortPos = pixel.V(0., 0.)

	options.VSync = true

	mainFont, err := typeface.LoadTTF("assets/FR73PixD.ttf", constants.TypeFaceSize)
	if err != nil {
		panic(err)
	}
	typeface.Atlases["main"] = text.NewAtlas(mainFont, text.ASCII)
	stickyFont, err := typeface.LoadTTF("assets/HomemadeApple-Regular.ttf", constants.TypeFaceSize)
	if err != nil {
		panic(err)
	}
	typeface.Atlases["sticky"] = text.NewAtlas(stickyFont, text.ASCII)

	loading.LoadImg()

	sfx.MusicPlayer.RegisterMusicTrack("assets/thesong.wav", "song")
	sfx.MusicPlayer.NewSet("song", []string{"song"}, sfx.Repeat, 0., 2.)

	sfx.SoundPlayer.RegisterSound("assets/alarm.wav", "alarm")
	sfx.SoundPlayer.RegisterSound("assets/buttonpress.wav", "buttonpress")
	sfx.SoundPlayer.RegisterSound("assets/conveyor.wav", "conveyor")
	sfx.SoundPlayer.RegisterSound("assets/pickup.wav", "pickup")
	sfx.SoundPlayer.RegisterSound("assets/place.wav", "place")
	sfx.SoundPlayer.RegisterSound("assets/place2.wav", "place2")
	sfx.SoundPlayer.RegisterSound("assets/trash.wav", "trash")

	debug.Initialize(&viewport.MainCamera.PostCamPos, &viewport.MainCamera.PostCamPos)

	state.InitMenu(win)
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
