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
			"debugText":  pxginput.NewJoyless(pixelgl.KeyF4),
			"fullscreen": pxginput.NewJoyless(pixelgl.KeyF5),
			//"debugMenu":    pxginput.NewJoyless(pixelgl.KeyF7),
			"debugClear":      pxginput.NewJoyless(pixelgl.KeyF8),
			"debugIgnoreConv": pxginput.NewJoyless(pixelgl.KeyF9),
			//"debugResume":  pxginput.NewJoyless(pixelgl.KeyF10),
			//"debugSP":      pxginput.NewJoyless(pixelgl.KeyEqual),
			//"debugSM":      pxginput.NewJoyless(pixelgl.KeyMinus),
			//"camUp":        pxginput.NewJoyless(pixelgl.KeyP),
			//"camRight":     pxginput.NewJoyless(pixelgl.KeyApostrophe),
			//"camDown":      pxginput.NewJoyless(pixelgl.KeySemicolon),
			//"camLeft":      pxginput.NewJoyless(pixelgl.KeyL),
			"debugAutoGenMode": pxginput.NewJoyless(pixelgl.KeyF10),
		},
		Mode: pxginput.KeyboardMouse,
	}
	gameInput = &pxginput.Input{
		Buttons: map[string]*pxginput.ButtonSet{
			"moveDown":   pxginput.NewJoyless(pixelgl.KeyS),
			"moveLeft":   pxginput.NewJoyless(pixelgl.KeyA),
			"moveRight":  pxginput.NewJoyless(pixelgl.KeyD),
			"rotate":     pxginput.NewJoyless(pixelgl.KeyW),
			"reset":      pxginput.NewJoyless(pixelgl.KeyR),
			"speedUp":    pxginput.NewJoyless(pixelgl.KeyPageUp),
			"speedDown":  pxginput.NewJoyless(pixelgl.KeyPageDown),
			"click":      pxginput.NewJoyless(pixelgl.MouseButtonLeft),
			"rightClick": pxginput.NewJoyless(pixelgl.MouseButtonRight),
			"generate":   pxginput.NewJoyless(pixelgl.KeyKPEnter),
			"pause":      pxginput.NewJoyless(pixelgl.KeyEscape),
			"showTitle":  pxginput.NewJoyless(pixelgl.KeyT),
		},
		Mode: pxginput.KeyboardMouse,
	}
)
