package state

import (
	"github.com/faiface/pixel/pixelgl"
	pxginput "github.com/timsims1717/pixel-go-input"
)

var (
	debugInput = &pxginput.Input{
		Buttons: map[string]*pxginput.ButtonSet{
			//"debugConsole": pxginput.NewJoyless(pixelgl.KeyGraveAccent),
			//"debug":        pxginput.NewJoyless(pixelgl.KeyF3),
			//"debugText":    pxginput.NewJoyless(pixelgl.KeyF4),
			"fullscreen": pxginput.NewJoyless(pixelgl.KeyF5),
			//"debugMenu":    pxginput.NewJoyless(pixelgl.KeyF7),
			//"debugTest":    pxginput.NewJoyless(pixelgl.KeyF8),
			//"debugPause":   pxginput.NewJoyless(pixelgl.KeyF9),
			//"debugResume":  pxginput.NewJoyless(pixelgl.KeyF10),
			//"debugSP":      pxginput.NewJoyless(pixelgl.KeyEqual),
			//"debugSM":      pxginput.NewJoyless(pixelgl.KeyMinus),
			//"camUp":        pxginput.NewJoyless(pixelgl.KeyP),
			//"camRight":     pxginput.NewJoyless(pixelgl.KeyApostrophe),
			//"camDown":      pxginput.NewJoyless(pixelgl.KeySemicolon),
			//"camLeft":      pxginput.NewJoyless(pixelgl.KeyL),
		},
		Mode: pxginput.KeyboardMouse,
	}
)
