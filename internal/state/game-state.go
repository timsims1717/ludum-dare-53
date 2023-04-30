package state

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
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
	factoryViewPort *viewport.ViewPort
}

func (s *gameState) Unload() {

}

func (s *gameState) Load(done chan struct{}) {
	s.tetrisViewport = viewport.New(nil)
	s.tetrisViewport.SetRect(pixel.R(0, 0, world.TileSize*constants.TetrisWidth, world.TileSize*constants.TetrisHeight))
	s.tetrisViewport.CamPos = pixel.V(world.TileSize*0.5*(constants.TetrisWidth-1), world.TileSize*0.5*(constants.TetrisHeight-1))
	data.NewTetrisBoard(constants.DefaultSpeed)
	data.TetrisBoard.NextShape = systems.NewTetronimo()
	systems.PlaceTetronimo()
	BuildTetrisBG()

	s.factoryViewPort = viewport.New(nil)
	s.factoryViewPort.SetRect(pixel.R(0, 0, constants.FactoryTile*constants.FactoryWidth, world.TileSize*constants.FactoryHeight))
	s.factoryViewPort.CamPos = pixel.V(constants.FactoryTile*0.5*(constants.FactoryWidth-1), world.TileSize*0.5*(constants.FactoryHeight-1))

	BuildFactoryBG()

	s.UpdateViews()
	done <- struct{}{}
}

func (s *gameState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")

	tetrisInput.Update(win, viewport.MainCamera.Mat)
	factoryInput.Update(win, viewport.MainCamera.Mat)
	debug.AddText(fmt.Sprintf("Mouse Input: (%d,%d)", int(factoryInput.World.X), int(factoryInput.World.Y)))
	debug.AddText(fmt.Sprintf("Factory Input: (%d,%d)", int(s.factoryViewPort.Projected(factoryInput.World).X), int(s.factoryViewPort.Projected(factoryInput.World).Y)))

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
	if tetrisInput.Get("reset").JustPressed() {
		if systems.FailCondition {
			systems.FailCondition = false
			systems.ClearBoard()
			data.TetrisBoard.Stats.FullReset()
		}
	}
	if tetrisInput.Get("speedUp").JustPressed() {
		data.TetrisBoard.SpeedUp()
	}
	if tetrisInput.Get("speedDown").JustPressed() {
		data.TetrisBoard.SpeedDown()
	}

	if factoryInput.Get("generate").JustPressed() {
		tet := systems.CreateFactoryTet(s.factoryViewPort.Projected(factoryInput.World), data.RandColor())
		tet.Entity.AddComponent(myecs.ViewPort, s.factoryViewPort)
		tet.Entity.AddComponent(myecs.Input, factoryInput)
		tet.Entity.AddComponent(myecs.Click, data.NewFn(func() {
			if tet.Entity.HasComponent(myecs.Drag) {
				tet.Object.Pos = tet.LastPos
				tet.Entity.RemoveComponent(myecs.Drag)
			} else if data.DraggingPiece == nil {
				tet.Entity.AddComponent(myecs.Drag, &factoryInput.World)
			}
		}))
	}
	if factoryInput.Get("click").JustPressed() {

	}

	systems.BlockSystem()
	systems.TetrisSystem()
	systems.ClickSystem()
	systems.DragSystem()
	systems.ParentSystem()
	systems.ObjectSystem()
	debug.AddText(fmt.Sprintf("Tetris Score: %03d", data.TetrisBoard.Stats.Score))
	debug.AddText(fmt.Sprintf("Current Streak: %d", data.TetrisBoard.Stats.Streak))
	debug.AddText(fmt.Sprintf("Current Speed: %f", data.TetrisBoard.Speed))
	debug.AddText(fmt.Sprintf("Current Level: %d", data.TetrisBoard.Stats.Checkpoint))
	debug.AddText(fmt.Sprintf("Current Piece: %s", data.TetrisBoard.Shape.TetType.String()))
	debug.AddText(fmt.Sprintf("Next Piece: %s", data.TetrisBoard.NextShape.TetType.String()))
	if systems.FailCondition {
		debug.AddText("Game Over, dun dun dun")
	}
	s.tetrisViewport.Update()
	s.factoryViewPort.Update()
}

func (s *gameState) Draw(win *pixelgl.Window) {
	s.factoryViewPort.Canvas.Clear(colornames.Green)
	systems.DrawSystem(win, 11)
	systems.DrawSystem(win, 12)
	systems.DrawSystem(win, 13)
	img.Batchers[constants.BlockKey].Draw(s.factoryViewPort.Canvas)
	img.Clear()
	s.factoryViewPort.Canvas.Draw(win, s.factoryViewPort.Mat)
	s.tetrisViewport.Canvas.Clear(colornames.Yellow)
	systems.DrawSystem(win, 1)
	systems.DrawSystem(win, 2)
	img.Batchers[constants.BlockKey].Draw(s.tetrisViewport.Canvas)
	img.Clear()
	s.tetrisViewport.Canvas.Draw(win, s.tetrisViewport.Mat)
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

	s.factoryViewPort.PortPos = portPos
	s.factoryViewPort.PortPos.X -= 0.25 * viewport.MainCamera.Rect.W()
	s.factoryViewPort.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W()*0.5, viewport.MainCamera.Rect.H()))
}
