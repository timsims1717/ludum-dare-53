package state

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/systems"
	"timsims1717/ludum-dare-53/pkg/debug"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/state"
	"timsims1717/ludum-dare-53/pkg/viewport"
	"timsims1717/ludum-dare-53/pkg/world"
)

type gameState struct {
	*state.AbstractState

	// tetris half
	tetrisViewport *viewport.ViewPort
	halfViewport   *viewport.ViewPort

	lastLeft bool

	// factory half
	floorViewport *viewport.ViewPort
}

func (s *gameState) Unload() {

}

func (s *gameState) Load(done chan struct{}) {
	s.tetrisViewport = viewport.New(nil)
	s.tetrisViewport.SetRect(pixel.R(0, 0, world.TileSize*constants.TetrisWidth, world.TileSize*constants.TetrisHeight))
	s.tetrisViewport.CamPos = pixel.V(world.TileSize*0.5*(constants.TetrisWidth-1), world.TileSize*0.5*(constants.TetrisHeight-1))
	data.NewTetrisBoard(0.4)
	systems.CreateTetronimo()
	BuildTetrisBG()

	s.floorViewport = viewport.New(nil)
	s.floorViewport.SetRect(pixel.R(0, 0, constants.FactoryTile*constants.FactoryWidth, world.TileSize*constants.FactoryHeight))
	s.floorViewport.CamPos = pixel.V(constants.FactoryTile*0.5*(constants.FactoryWidth-1), world.TileSize*0.5*(constants.FactoryHeight-1))

	BuildFactoryBG()

	s.UpdateViews()
	done <- struct{}{}
}

func (s *gameState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")

	tetrisInput.Update(win, viewport.MainCamera.Mat)
	factoryInput.Update(win, viewport.MainCamera.Mat)

	if tetrisInput.Get("moveDown").JustPressed() || tetrisInput.Get("moveDown").Repeated() {
		systems.MoveDown = true
	}
	if tetrisInput.Get("moveLeft").JustPressed() || (tetrisInput.Get("moveLeft").Pressed() && s.lastLeft) {
		s.lastLeft = true
		if tetrisInput.Get("moveRight").JustPressed() {
			s.lastLeft = false
			systems.MoveRight = true
		} else if tetrisInput.Get("moveLeft").JustPressed() || tetrisInput.Get("moveLeft").Repeated() {
			systems.MoveLeft = true
		}
	} else if tetrisInput.Get("moveRight").JustPressed() || tetrisInput.Get("moveRight").Pressed() && !s.lastLeft {
		s.lastLeft = false
		if tetrisInput.Get("moveLeft").JustPressed() {
			s.lastLeft = true
			systems.MoveLeft = true
		} else if tetrisInput.Get("moveRight").JustPressed() || tetrisInput.Get("moveRight").Repeated() {
			systems.MoveRight = true
		}
	}
	if tetrisInput.Get("rotate").JustPressed() {
		systems.Rotate = true
	}

	systems.BlockSystem()
	systems.TetrisSystem()
	systems.DragSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	s.tetrisViewport.Update()
	s.floorViewport.Update()
}

func (s *gameState) Draw(win *pixelgl.Window) {
	s.tetrisViewport.Canvas.Clear(colornames.Yellow)
	systems.DrawSystem(win, 1)
	systems.DrawSystem(win, 2)
	img.Batchers[constants.BlockKey].Draw(s.tetrisViewport.Canvas)
	img.Clear()
	s.tetrisViewport.Canvas.Draw(win, s.tetrisViewport.Mat)
	s.floorViewport.Canvas.Clear(colornames.Green)
	systems.DrawSystem(win, 11)
	systems.DrawSystem(win, 12)
	img.Batchers[constants.BlockKey].Draw(s.floorViewport.Canvas)
	img.Clear()
	s.floorViewport.Canvas.Draw(win, s.floorViewport.Mat)
}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *gameState) UpdateViews() {
	portPos := pixel.V(viewport.MainCamera.PostCamPos.X+viewport.MainCamera.Rect.W()*0.5, viewport.MainCamera.PostCamPos.Y+viewport.MainCamera.Rect.H()*0.5)
	s.tetrisViewport.PortPos = portPos
	hRatio := viewport.MainCamera.Rect.H() / s.tetrisViewport.Rect.H()
	hRatio *= 0.8
	s.tetrisViewport.PortSize = pixel.V(hRatio, hRatio)
	s.tetrisViewport.PortPos.X += 0.5 * hRatio * s.tetrisViewport.Canvas.Bounds().W()

	s.floorViewport.PortPos = portPos
	s.floorViewport.PortSize = pixel.V(hRatio, hRatio)
	s.floorViewport.PortPos.X -= 0.5 * hRatio * s.floorViewport.Canvas.Bounds().W()
	s.floorViewport.PortPos.X -= 96.
}
