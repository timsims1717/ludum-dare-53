package state

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/internal/systems"
	"timsims1717/ludum-dare-53/pkg/debug"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/options"
	"timsims1717/ludum-dare-53/pkg/reanimator"
	"timsims1717/ludum-dare-53/pkg/state"
	"timsims1717/ludum-dare-53/pkg/typeface"
	"timsims1717/ludum-dare-53/pkg/viewport"
	"timsims1717/ludum-dare-53/pkg/world"
)

type gameState struct {
	*state.AbstractState

	lastLeft bool
}

func (s *gameState) Unload() {
	systems.ClearSystem()
}

func (s *gameState) Load(done chan struct{}) {

	data.StickyViewport = viewport.New(nil)
	data.StickyViewport.SetRect(pixel.R(0, 0, 1024, 1024))
	data.StickyViewport.CamPos = pixel.ZV

	data.TetrisViewport = viewport.New(nil)
	data.TetrisViewport.SetRect(pixel.R(0, 0, world.TileSize*constants.TetrisWidth, world.TileSize*constants.TetrisHeight))
	data.TetrisViewport.CamPos = pixel.V(world.TileSize*0.5*(constants.TetrisWidth-1), world.TileSize*0.5*(constants.TetrisHeight-1))
	data.NewTetrisBoard()
	BuildTetrisBG()

	data.FactoryViewport = viewport.New(nil)
	data.FactoryViewport.SetRect(pixel.R(0, 0, constants.FactoryTile*constants.FactoryWidth, world.TileSize*constants.FactoryHeight))
	data.FactoryViewport.CamPos = pixel.V(constants.FactoryTile*0.5*(constants.FactoryWidth-1), world.TileSize*0.5*(constants.FactoryHeight-1))
	data.NewFactoryFloor()
	LoadTileMaps()
	BuildFactoryFloor()
	BuildFactoryPads()
	CreateConveyor()
	CreateTrucks()

	data.TetrisBoard.NextShape = systems.NewTetromino()
	systems.PlaceTetromino()

	data.SBLabels = typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1.5, 1., 0, 0)
	data.SBLabels.Obj.Layer = 12
	data.SBLabels.SetPos(pixel.V(-21.5*data.MSize, 15.5*data.MSize))
	data.SBLabels.SetColor(constants.TVTextColor)
	data.SBLabels.SetSize(0.12)
	data.SBLabels.SetText("Scores")

	data.SBScores = typeface.New(nil, "main", typeface.NewAlign(typeface.Right, typeface.Top), 1.5, 1., 0, 0)
	data.SBScores.Obj.Layer = 12
	data.SBScores.SetPos(pixel.V(-3.5*data.MSize, 15.5*data.MSize))
	data.SBScores.SetColor(constants.TVTextColor)
	data.SBScores.SetSize(0.12)
	data.SBScores.SetText("Scores")

	data.StickyText = typeface.New(nil, "sticky", typeface.NewAlign(typeface.Center, typeface.Center), 1.5, 0.32, 0, 0)
	data.StickyText.SetPos(pixel.V(0., 0.))
	data.StickyText.SetColor(constants.BlackColor)
	data.StickyText.SetText("Paused")

	s.UpdateViews()
	reanimator.SetFrameRate(16)
	reanimator.Reset()
	done <- struct{}{}
}

func (s *gameState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")

	if options.Updated {
		s.UpdateViews()
	}
	gameInput.Update(win, viewport.MainCamera.Mat)
	if gameInput.Get("pause").JustPressed() {
		data.Paused = !data.Paused
		data.StickyOpen = data.Paused
		if data.Paused {
			data.SetStickyMsg(data.PauseMsg)
			openPauseMenu()
		} else {
			closeMenu()
		}
	}
	if win.Focused() && !data.Paused {
		reanimator.Update()
		debug.AddText(fmt.Sprintf("Mouse Input: (%d,%d)", int(gameInput.World.X), int(gameInput.World.Y)))
		debug.AddText(fmt.Sprintf("Factory Input: (%d,%d)", int(data.FactoryViewport.Projected(gameInput.World).X), int(data.FactoryViewport.Projected(gameInput.World).Y)))

		if gameInput.Get("moveDown").JustPressed() || gameInput.Get("moveDown").Repeated() {
			systems.MoveDown = true
		}
		if gameInput.Get("moveLeft").JustPressed() || (gameInput.Get("moveLeft").Pressed() && s.lastLeft) {
			s.lastLeft = true
			if gameInput.Get("moveRight").JustPressed() {
				s.lastLeft = false
				systems.MoveRight = true
			} else if gameInput.Get("moveLeft").JustPressed() || gameInput.Get("moveLeft").Repeated() {
				systems.MoveLeft = true
			}
		} else if gameInput.Get("moveRight").JustPressed() || gameInput.Get("moveRight").Pressed() && !s.lastLeft {
			s.lastLeft = false
			if gameInput.Get("moveLeft").JustPressed() {
				s.lastLeft = true
				systems.MoveLeft = true
			} else if gameInput.Get("moveRight").JustPressed() || gameInput.Get("moveRight").Repeated() {
				systems.MoveRight = true
			}
		}
		if gameInput.Get("rotate").JustPressed() {
			systems.Rotate = true
		}
		if gameInput.Get("reset").JustPressed() {
			if systems.FailCondition {
				systems.FailCondition = false
				systems.ClearBoard()
				systems.ClearFactory()
			}
		}
		if gameInput.Get("speedUp").JustPressed() {
			data.TetrisBoard.SpeedUp()
		}
		if gameInput.Get("speedDown").JustPressed() {
			data.TetrisBoard.SpeedDown()
		}

		if gameInput.Get("generate").JustPressed() {
			gameInput.Get("generate").Consume()
			r1 := rand.Intn(len(data.FactoryPads))
			pad := data.FactoryPads[r1]
			r := r1
			for pad.Tet != nil {
				r++
				r %= len(data.FactoryPads)
				if r == r1 {
					break
				}
				pad = data.FactoryPads[r]
			}
			if pad.Tet == nil {

				tet := systems.CreateFactoryTet(pad.Object.Pos, data.RandColor(), constants.FacUndefined)
				tet.Object.Hide = false
				tet.Entity.AddComponent(myecs.ViewPort, data.FactoryViewport)

				tet.Entity.AddComponent(myecs.Input, gameInput)
				pad.Tet = tet
				//tet.Entity.AddComponent(myecs.Click, data.NewFn(func() {
				//	if tet.Entity.HasComponent(myecs.Drag) {
				//		tet.Object.Pos = tet.LastPos
				//		tet.Entity.RemoveComponent(myecs.Drag)
				//	} else if data.DraggingPiece == nil {
				//		tet.Entity.AddComponent(myecs.Drag, &gameInput.World)
				//	}
				//}))
			}
		}

		systems.FunctionSystem()
		systems.BlockSystem()
		systems.TetrisSystem()
		systems.FactoryBlockSystem()
		systems.ClickSystem()
		systems.DragSystem()
		systems.ParentSystem()
		systems.ObjectSystem()
		systems.AnimationSystem()
	} else if data.Paused {
		systems.MenuSystem(gameInput.World, gameInput.Get("click").JustPressed())
	}
	debug.AddText(fmt.Sprintf("Global Score: %03d", data.TetrisBoard.Stats.GlobalScore()))
	debug.AddText(fmt.Sprintf("Tetris Score: %03d", data.TetrisBoard.Stats.Score))
	debug.AddText(fmt.Sprintf("Line Clearing Points: +%d", data.TetrisBoard.Stats.MyFibScore.FibN-1))
	debug.AddText(fmt.Sprintf("Factory Score: %03d", data.FactoryFloor.Stats.Score))
	debug.AddText(fmt.Sprintf("Factory Balance Bonus: +%d", data.FactoryFloor.Stats.MyFibScore.FibN-1))
	debug.AddText(fmt.Sprintf("Current Speed: %f", data.TetrisBoard.Speed))
	debug.AddText(fmt.Sprintf("Current Level: %d", data.TetrisBoard.Stats.Checkpoint))
	if data.TetrisBoard.Shape != nil {
		debug.AddText(fmt.Sprintf("Current Piece: %s", data.TetrisBoard.Shape.TetType.String()))
	}
	if data.TetrisBoard.NextShape != nil {
		debug.AddText(fmt.Sprintf("Next Piece: %s", data.TetrisBoard.NextShape.TetType.String()))
	}
	debug.AddText(fmt.Sprintf("Tetrominos: %d", data.TetrisBoard.Stats.Tetrominos))
	debug.AddText(fmt.Sprintf("Factrominos: %d", data.FactoryFloor.Stats.Factrominos))
	debug.AddText(fmt.Sprintf("Total Trashed Shapes: %d", data.FactoryFloor.Stats.TotalTrashedShapes()))
	if data.Conveyor != nil {
		count := 0
		for _, block := range data.Conveyor.Tets {
			if block != nil {
				count++
			}
		}
		debug.AddText(fmt.Sprintf("Num of Blocks on Conveyor: %d", count))
	}
	if systems.FailCondition {
		debug.AddText("Game Over, done done done")
	}
	debug.AddText(fmt.Sprintf("PieceDone: %t", systems.PieceDone))

	data.SBLabels.SetText("Score:\nBalance Bonus:\nClear Bonus:")
	data.SBLabels.Obj.Update()

	data.SBScores.SetText(fmt.Sprintf("%05d\n+%d\n+%d", data.TetrisBoard.Stats.GlobalScore(), data.FactoryFloor.Stats.MyFibScore.FibN-1, data.TetrisBoard.Stats.MyFibScore.FibN-1))
	data.SBScores.Obj.Update()

	data.TetrisViewport.Update()
	data.FactoryViewport.Update()
	data.StickyText.Obj.Update()
	data.StickyObj.Update()
	data.StickyViewport.Update()
}

func (s *gameState) Draw(win *pixelgl.Window) {
	data.FactoryViewport.Canvas.Clear(colornames.Green)
	systems.DrawSystem(win, 9)  // floor
	systems.DrawSystem(win, 10) // trucks
	systems.DrawSystem(win, 11) // walls
	systems.DrawSystem(win, 12) // tiles
	img.Batchers[constants.FactoryKey].Draw(data.FactoryViewport.Canvas)
	img.Batchers[constants.BlockKey].Draw(data.FactoryViewport.Canvas)
	img.Clear()
	data.SBLabels.Draw(data.FactoryViewport.Canvas)
	data.SBScores.Draw(data.FactoryViewport.Canvas)
	systems.DrawSystem(win, 13)
	systems.DrawSystem(win, 14)
	systems.DrawSystem(win, 15)
	systems.DrawSystem(win, 16)
	systems.DrawSystem(win, 17)
	systems.DrawSystem(win, 18)
	systems.DrawSystem(win, 19)
	systems.DrawSystem(win, 20) // dragged tile
	systems.DrawSystem(win, 21)
	systems.DrawSystem(win, 22)
	systems.DrawSystem(win, 23)
	systems.DrawSystem(win, 24)
	systems.DrawSystem(win, 25)
	systems.DrawSystem(win, 26)
	img.Batchers[constants.FactoryKey].Draw(data.FactoryViewport.Canvas)
	img.Batchers[constants.BlockKey].Draw(data.FactoryViewport.Canvas)
	img.Clear()
	data.FactoryViewport.Canvas.Draw(win, data.FactoryViewport.Mat)
	data.TetrisViewport.Canvas.Clear(colornames.Yellow)
	systems.DrawSystem(win, 1)
	systems.DrawSystem(win, 2)
	img.Batchers[constants.BlockKey].Draw(data.TetrisViewport.Canvas)
	img.Clear()
	data.TetrisViewport.Canvas.Draw(win, data.TetrisViewport.Mat)
	if data.StickyOpen {
		data.StickyViewport.Canvas.Clear(color.RGBA{})
		data.StickyNote.Draw(data.StickyViewport.Canvas, data.StickyObj.Mat)
		data.StickyText.Draw(data.StickyViewport.Canvas)
		systems.DrawMenuSystem(data.StickyViewport)
		data.StickyViewport.Canvas.Draw(win, data.StickyViewport.Mat)
	}
	systems.TemporarySystem()
}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *gameState) UpdateViews() {
	portPos := pixel.V(viewport.MainCamera.PostCamPos.X+viewport.MainCamera.Rect.W()*0.5, viewport.MainCamera.PostCamPos.Y+viewport.MainCamera.Rect.H()*0.5)
	hRatio := viewport.MainCamera.Rect.H() / (world.TileSize * 20)
	hRatio *= 0.8
	data.TetrisViewport.PortPos = portPos
	//s.tetrisViewport.PortSize = pixel.V(hRatio, hRatio)
	data.TetrisViewport.PortPos.X -= 25 * data.MSize
	data.TetrisViewport.PortPos.Y -= 3.5 * data.MSize
	//s.tetrisViewport.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W()*0.5, viewport.MainCamera.Rect.H()))

	data.FactoryViewport.PortPos = portPos
	//s.factoryViewPort.PortSize = pixel.V(hRatio, hRatio)
	//s.factoryViewPort.PortPos.X += 0.25 * viewport.MainCamera.Rect.W()
	data.FactoryViewport.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W(), viewport.MainCamera.Rect.H()))
	data.FactoryViewport.CamPos.Y = data.MSize * 10.
	data.FactoryViewport.CamPos.X = data.MSize * -10.

	data.StickyViewport.PortPos = portPos
	data.StickyViewport.PortSize = pixel.V(0.8, 0.8)
}
