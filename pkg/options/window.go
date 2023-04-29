package options

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	VSync           bool
	FullScreen      bool
	ResolutionIndex int
	Resolutions     []pixel.Vec

	fullscreen bool
	resIndex   int
)

func RegisterResolution(res pixel.Vec) {
	Resolutions = append(Resolutions, res)
}

func WindowUpdate(win *pixelgl.Window) {
	if win.Focused() {
		win.SetVSync(VSync)
		if FullScreen != fullscreen {
			// get window position (center)
			pos := win.GetPos()
			pos.X += win.Bounds().W() * 0.5
			pos.Y += win.Bounds().H() * 0.5

			// find current monitor
			var picked *pixelgl.Monitor
			if len(pixelgl.Monitors()) > 1 {
				for _, m := range pixelgl.Monitors() {
					x, y := m.Position()
					w, h := m.Size()
					if pos.X >= x && pos.X <= x+w && pos.Y >= y && pos.Y <= y+h {
						picked = m
						break
					}
				}
				if picked == nil {
					pos = win.GetPos()
					for _, m := range pixelgl.Monitors() {
						x, y := m.Position()
						w, h := m.Size()
						if pos.X >= x && pos.X <= x+w && pos.Y >= y && pos.Y <= y+h {
							picked = m
							break
						}
					}
				}
			}
			if picked == nil {
				picked = pixelgl.PrimaryMonitor()
			}
			if FullScreen {
				win.SetMonitor(picked)
			} else {
				win.SetMonitor(nil)
			}
			fullscreen = FullScreen
		}
	}
}
