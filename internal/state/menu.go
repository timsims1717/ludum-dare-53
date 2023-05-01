package state

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/options"
	"timsims1717/ludum-dare-53/pkg/typeface"
)

var (
	resumeItem = &data.MenuItem{}
	quitItem   = &data.MenuItem{}
	optionItem = &data.MenuItem{}
	vsyncItem  = &data.MenuItem{}
	fullScreen = &data.MenuItem{}
	backItem   = &data.MenuItem{}
)

func InitMenu(win *pixelgl.Window) {
	resumeItem.Click = func() {
		if data.Paused {
			data.Paused = false
			data.StickyOpen = false
			closeMenu()
		}
	}
	resumeItem.Text = typeface.New(nil, "sticky", typeface.NewAlign(typeface.Center, typeface.Center), 1.5, 0.32, 0, 0)
	resumeItem.Text.SetPos(pixel.V(0., 0.))
	resumeItem.Text.SetSize(0.32)
	resumeItem.Text.SetColor(constants.BlackColor)
	resumeItem.Text.SetText("Resume")
	resumeItem.Text.Obj = object.New()
	resumeItem.Text.Obj.Pos.X = -120.
	resumeItem.Text.Obj.Pos.Y = 136.
	resumeItem.Text.Obj.Hide = true
	resumeItem.OrigSize = 0.32
	quitItem.Click = func() {
		win.SetClosed(true)
	}
	quitItem.Text = typeface.New(nil, "sticky", typeface.NewAlign(typeface.Center, typeface.Center), 1.5, 0.32, 0, 0)
	quitItem.Text.SetPos(pixel.V(0., 0.))
	quitItem.Text.SetSize(0.32)
	quitItem.Text.SetColor(constants.BlackColor)
	quitItem.Text.SetText("Quit")
	quitItem.Text.Obj = object.New()
	quitItem.Text.Obj.Pos.X = 65.
	quitItem.Text.Obj.Pos.Y = -168.
	quitItem.Text.Obj.Hide = true
	quitItem.OrigSize = 0.32
	optionItem.Click = func() {
		closeMenu()
		openOptionsMenu()
	}
	optionItem.Text = typeface.New(nil, "sticky", typeface.NewAlign(typeface.Center, typeface.Center), 1.5, 0.32, 0, 0)
	optionItem.Text.SetPos(pixel.V(0., 0.))
	optionItem.Text.SetSize(0.32)
	optionItem.Text.SetColor(constants.BlackColor)
	optionItem.Text.SetText("Options")
	optionItem.Text.Obj = object.New()
	optionItem.Text.Obj.Pos.X = 118.
	optionItem.Text.Obj.Pos.Y = 36.
	optionItem.Text.Obj.Hide = true
	optionItem.OrigSize = 0.32
	vsyncItem.Click = func() {
		options.VSync = !options.VSync
		if !options.VSync {
			vsyncItem.Text.SetText("VSync (It's off)")
		} else {
			vsyncItem.Text.SetText("VSync (It's on)")
		}
	}
	vsyncItem.Text = typeface.New(nil, "sticky", typeface.NewAlign(typeface.Center, typeface.Center), 1.5, 0.26, 0, 0)
	vsyncItem.Text.SetPos(pixel.V(0., 0.))
	vsyncItem.Text.SetSize(0.26)
	vsyncItem.Text.SetColor(constants.BlackColor)
	vsyncItem.Text.SetText("VSync (It's on)")
	vsyncItem.Text.Obj = object.New()
	vsyncItem.Text.Obj.Pos.X = 118.
	vsyncItem.Text.Obj.Pos.Y = 132.
	vsyncItem.Text.Obj.Hide = true
	vsyncItem.OrigSize = 0.32
	fullScreen.Click = func() {
		options.FullScreen = !options.FullScreen
		if !options.FullScreen {
			fullScreen.Text.SetText("Fullsrceen (It's off)")
		} else {
			fullScreen.Text.SetText("FullScreen (It's on)")
		}
	}
	fullScreen.Text = typeface.New(nil, "sticky", typeface.NewAlign(typeface.Center, typeface.Center), 1.5, 0.26, 0, 0)
	fullScreen.Text.SetPos(pixel.V(0., 0.))
	fullScreen.Text.SetSize(0.26)
	fullScreen.Text.SetColor(constants.BlackColor)
	fullScreen.Text.SetText("FullScreen (It's off)")
	fullScreen.Text.Obj = object.New()
	fullScreen.Text.Obj.Pos.X = -32.
	fullScreen.Text.Obj.Pos.Y = 45.
	fullScreen.Text.Obj.Hide = true
	fullScreen.OrigSize = 0.26
	backItem.Click = func() {
		closeMenu()
		openPauseMenu()
	}
	backItem.Text = typeface.New(nil, "sticky", typeface.NewAlign(typeface.Center, typeface.Center), 1.5, 0.32, 0, 0)
	backItem.Text.SetPos(pixel.V(0., 0.))
	backItem.Text.SetSize(0.32)
	backItem.Text.SetColor(constants.BlackColor)
	backItem.Text.SetText("Back")
	backItem.Text.Obj = object.New()
	backItem.Text.Obj.Pos.X = -145.
	backItem.Text.Obj.Pos.Y = -162.
	backItem.Text.Obj.Hide = true
	backItem.OrigSize = 0.32

	resumeItem.Text.Obj.Update()
	quitItem.Text.Obj.Update()
	optionItem.Text.Obj.Update()
	vsyncItem.Text.Obj.Update()
	fullScreen.Text.Obj.Update()
	backItem.Text.Obj.Update()

	data.MenuItems = append(data.MenuItems, resumeItem)
	data.MenuItems = append(data.MenuItems, quitItem)
	data.MenuItems = append(data.MenuItems, optionItem)
	data.MenuItems = append(data.MenuItems, vsyncItem)
	data.MenuItems = append(data.MenuItems, fullScreen)
	data.MenuItems = append(data.MenuItems, backItem)
}

func openPauseMenu() {
	resumeItem.Text.Obj.Hide = false
	quitItem.Text.Obj.Hide = false
	optionItem.Text.Obj.Hide = false
}

func openOptionsMenu() {
	vsyncItem.Text.Obj.Hide = false
	fullScreen.Text.Obj.Hide = false
	backItem.Text.Obj.Hide = false
}

func closeMenu() {
	resumeItem.Text.Obj.Hide = true
	quitItem.Text.Obj.Hide = true
	optionItem.Text.Obj.Hide = true
	vsyncItem.Text.Obj.Hide = true
	fullScreen.Text.Obj.Hide = true
	backItem.Text.Obj.Hide = true
}
