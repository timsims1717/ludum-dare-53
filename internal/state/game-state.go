package state

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thoas/go-funk"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"strconv"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/internal/systems"
	"timsims1717/ludum-dare-53/pkg/debug"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/options"
	"timsims1717/ludum-dare-53/pkg/reanimator"
	"timsims1717/ludum-dare-53/pkg/sfx"
	"timsims1717/ludum-dare-53/pkg/state"
	"timsims1717/ludum-dare-53/pkg/timing"
	"timsims1717/ludum-dare-53/pkg/typeface"
	"timsims1717/ludum-dare-53/pkg/viewport"
	"timsims1717/ludum-dare-53/pkg/world"
)

type gameState struct {
	*state.AbstractState

	sfxTimer *timing.Timer
	start    bool
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

	data.ShCounts = typeface.New(nil, "main", typeface.NewAlign(typeface.Center, typeface.Top), 1.5, 1., 0, 0)
	data.ShCounts.Obj.Layer = 12
	data.ShCounts.SetPos(pixel.V(-13.*data.MSize, 1.8*data.MSize))
	data.ShCounts.SetColor(constants.TVTextColor)
	data.ShCounts.SetSize(0.08)
	data.ShCounts.SetText("Scores")

	data.StickyText = typeface.New(nil, "sticky", typeface.NewAlign(typeface.Center, typeface.Center), 1.5, 0.28, 3.5*constants.TypeFaceSize, 0)
	data.StickyText.SetPos(pixel.V(0., 0.))
	data.StickyText.SetColor(constants.BlackColor)
	data.StickyText.SetText("Paused")

	for i, achFam := range constants.AchievementFamilies {
		achFam.StickyNote = object.New().WithID(achFam.Name)
		achFam.StickyNote.Pos = achFam.StickyNotePosition
		achFam.StickyNote.Layer = 12
		achFam.StickyNote.Hide = true
		achFam.StickyNote.Rect = pixel.R(0, 0, 32, 32)
		constants.AchievementFamilies[i] = achFam
		newachFam := constants.AchievementFamilies[i]
		myecs.Manager.NewEntity().AddComponent(myecs.Object, constants.AchievementFamilies[i].StickyNote).
			AddComponent(myecs.Drawable, data.TinyNote).
			AddComponent(myecs.ViewPort, data.FactoryViewport).
			AddComponent(myecs.Input, gameInput).
			AddComponent(myecs.Click, data.NewFn(ClickAchievement(&newachFam))).
			AddComponent(myecs.Update, data.NewFn(func() {
				if newachFam.StickyNote.PointInside(data.FactoryViewport.Projected(gameInput.World)) &&
					newachFam.Achieved() {
					data.HandState = 1
				}
			}))
	}

	pauseBtnObj := object.New()
	pauseBtnObj.Pos = pixel.V(-295., -180.)
	pauseBtnObj.Rect = pixel.R(0, 0, 96, 64)
	pauseBtnObj.Layer = 11
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, pauseBtnObj).
		AddComponent(myecs.Drawable, data.PauseButSprs).
		AddComponent(myecs.Update, data.NewFn(func() {
			if pauseBtnObj.PointInside(data.FactoryViewport.Projected(gameInput.World)) {
				data.HandState = 1
				if gameInput.Get("click").Pressed() {
					data.PauseButSprs[1].Offset.Y = -3.
				} else {
					data.PauseButSprs[1].Offset.Y = 0
				}
				if gameInput.Get("click").JustReleased() {
					data.Paused = true
					openPauseMenu()
					sfx.SoundPlayer.PlaySound("buttonpress", 0.)
				}
			} else {
				data.PauseButSprs[1].Offset.Y = 0
			}
		}))

	restartBtnObj := object.New()
	restartBtnObj.Pos = pixel.V(-820., -180.)
	restartBtnObj.Rect = pixel.R(0, 0, 96, 64)
	restartBtnObj.Layer = 11
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, restartBtnObj).
		AddComponent(myecs.Drawable, data.RestartButSprs).
		AddComponent(myecs.Update, data.NewFn(func() {
			if restartBtnObj.PointInside(data.FactoryViewport.Projected(gameInput.World)) {
				data.HandState = 1
				if gameInput.Get("click").Pressed() {
					data.RestartButSprs[1].Offset.Y = -3.
				} else {
					data.RestartButSprs[1].Offset.Y = 0
				}
				if gameInput.Get("click").JustReleased() {
					if systems.FailCondition {
						systems.FailCondition = false
						systems.WasFail = false
						systems.ClearBoard()
						systems.ClearFactory()
					}
					sfx.SoundPlayer.PlaySound("buttonpress", 0.)
				}
			} else {
				data.RestartButSprs[1].Offset.Y = 0
			}
		}))

	stickyTutObj := object.New().WithID("sticky-tutorial")
	stickyTutObj.Pos = pixel.V(540., 480)
	stickyTutObj.Layer = 12
	stickyTutObj.Rect = pixel.R(0, 0, 32, 32)
	myecs.Manager.NewEntity().AddComponent(myecs.Object, stickyTutObj).
		AddComponent(myecs.Drawable, data.TinyNote).
		AddComponent(myecs.ViewPort, data.FactoryViewport).
		AddComponent(myecs.Input, gameInput).
		AddComponent(myecs.Click, data.NewFn(func() {
			OpenSticky(data.Instructions)
		})).
		AddComponent(myecs.Update, data.NewFn(func() {
			if stickyTutObj.PointInside(data.FactoryViewport.Projected(gameInput.World)) {
				data.HandState = 1
			}
		}))

	s.UpdateViews()
	sfx.MusicPlayer.PlayMusic("song")
	data.HandObj = object.New()
	reanimator.SetFrameRate(16)
	reanimator.Reset()
	s.sfxTimer = timing.New(rand.Float64()*20. + 5.)
	s.start = true
	done <- struct{}{}
}

func (s *gameState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")

	if options.Updated {
		s.UpdateViews()
	}
	gameInput.Update(win, viewport.MainCamera.Mat)
	obj := object.New()
	obj.Pos = gameInput.World
	spr := data.HandPoint

	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Temp, myecs.ClearFlag(true))
	if !data.PauseMenu && data.StickyOpen && gameInput.Get("click").JustPressed() {
		if !data.StickyObj.PointInside(data.StickyViewport.Projected(gameInput.World)) {
			CloseSticky()
		}
	}
	if gameInput.Get("pause").JustPressed() {
		data.Paused = !data.Paused
		if data.Paused {
			openPauseMenu()
		} else {
			closeMenu()
		}
	}
	if win.Focused() && !data.Paused {
		if s.start {
			s.start = false
			OpenSticky(data.Instructions)
		}
		reanimator.Update()
		debug.AddText(fmt.Sprintf("Mouse Input: (%d,%d)", int(gameInput.World.X), int(gameInput.World.Y)))
		debug.AddText(fmt.Sprintf("Factory Input: (%d,%d)", int(data.FactoryViewport.Projected(gameInput.World).X), int(data.FactoryViewport.Projected(gameInput.World).Y)))

		if systems.HoldDown {
			if gameInput.Get("moveDown").Pressed() {
				if systems.HoldDownT.UpdateDone() {
					systems.MoveDown = true
				}
			} else {
				systems.HoldDown = false
			}
		}
		if gameInput.Get("moveDown").JustPressed() {
			systems.MoveDown = true
			systems.HoldDown = true
			systems.HoldDownT = timing.New(0.25)
		}
		if gameInput.Get("moveLeft").JustPressed() {
			gameInput.Get("moveLeft").Consume()
			systems.MoveLeft = true
		} else if gameInput.Get("moveRight").JustPressed() {
			gameInput.Get("moveRight").Consume()
			systems.MoveRight = true
		}
		if gameInput.Get("rotate").JustPressed() {
			systems.Rotate = true
		}
		if gameInput.Get("reset").JustPressed() {
			if systems.FailCondition {
				systems.FailCondition = false
				systems.WasFail = false
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
		if gameInput.Get("showTitle").JustPressed() {
			OpenSticky(&data.StickyMsg{
				Message: constants.TitleText,
				Offset:  pixel.V(40., 55.),
			})
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
	}
	if data.Paused {
		systems.MenuSystem(gameInput)
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
		if !systems.WasFail {
			OpenSticky(&data.StickyMsg{
				Message: systems.FailReason.String(),
				Offset:  pixel.V(40., 55.),
			})
			systems.WasFail = true
		}
	}
	debug.AddText(fmt.Sprintf("PieceDone: %t", systems.PieceDone))

	data.SBLabels.SetText("Score:\nLevel:\nDeliveries:\nBalance Bonus:\nLines Cleared:\nClear Bonus:")
	data.SBLabels.Obj.Update()

	data.SBScores.SetText(fmt.Sprintf("%05d\n%d\n%03d\n+%d\n%03d\n+%d", data.TetrisBoard.Stats.GlobalScore(), data.TetrisBoard.Stats.Checkpoint, data.FactoryFloor.Stats.Factrominos, data.FactoryFloor.Stats.MyFibScore.FibN-1, data.TetrisBoard.Stats.LinesCleared, data.TetrisBoard.Stats.MyFibScore.FibN-1))
	data.SBScores.Obj.Update()

	bs := data.FactoryFloor.Stats.BuiltShapes
	data.ShCounts.SetText(fmt.Sprintf("%02d %02d %02d %02d %02d %02d %02d", bs[constants.O], bs[constants.S], bs[constants.Z], bs[constants.J], bs[constants.L], bs[constants.T], bs[constants.I]))
	data.ShCounts.Obj.Update()

	data.TetrisViewport.Update()
	data.FactoryViewport.Update()
	data.StickyText.Obj.Update()
	data.StickyObj.Update()
	data.StickyViewport.Update()
	if s.sfxTimer.UpdateDone() {
		switch rand.Intn(2) {
		case 0:
			sfx.SoundPlayer.PlaySound("conveyor", -2.)
		case 1:
			sfx.SoundPlayer.PlaySound("alarm", 0.)
		}
		s.sfxTimer = timing.New(rand.Float64()*20. + 5.)
	}
	data.HandObj.Pos = gameInput.World
	data.HandObj.Update()
	UpdateAchievements()
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
	data.ShCounts.Draw(data.FactoryViewport.Canvas)
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
	sprH := data.HandPoint
	if data.DraggingPiece != nil {
		sprH = data.HandGrab
	} else {
		switch data.HandState {
		case 0:
			sprH = data.HandOpen
		case 1:
			sprH = data.HandPoint
		case 2:
			sprH = data.HandGrab
		}
	}
	if sprH != nil {
		img.Batchers[constants.FactoryKey].DrawSpriteColor(sprH.Key, data.HandObj.Mat.Moved(sprH.Offset), sprH.Color)
		img.Batchers[constants.FactoryKey].Draw(win)
	}
	data.HandState = 0
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

func UpdateAchievements() {
	rawAchievementsForCalc := funk.Map(constants.Achievements, func(k string, value constants.Achievement) constants.Achievement {
		return value
	}).([]constants.Achievement)
	filteredTotalAchievements := funk.Filter(rawAchievementsForCalc, func(x constants.Achievement) bool {
		return x.MyFamily.Name != "AchievementProgress" && x.MyFamily.Name != "CompletedAchievements"
	}).([]constants.Achievement)
	Total := len(filteredTotalAchievements)
	filteredAchievementsCompleted := funk.Filter(rawAchievementsForCalc, func(x constants.Achievement) bool {
		return x.Achieved && x.MyFamily.Name != "AchievementProgress" && x.MyFamily.Name != "CompletedAchievements"
	}).([]constants.Achievement)
	Percent := float64(len(filteredAchievementsCompleted)) / float64(Total)

	filteredPercentAchievements := funk.Filter(rawAchievementsForCalc, func(x constants.Achievement) bool {
		return x.MyFamily.Name == "AchievementProgress"
	}).([]constants.Achievement)
	for _, value := range filteredPercentAchievements {
		if i, _ := strconv.ParseFloat(value.Properties["target"], 64); i <= Percent {
			temp := constants.Achievements[value.Name]
			if !temp.Achieved {
				temp.Achieved = true
			}
			constants.Achievements[value.Name] = temp
		}
	}
	if Total == len(filteredAchievementsCompleted) {
		temp := constants.Achievements["CompletedAchievements"]
		temp.Achieved = true
		constants.Achievements["CompletedAchievements"] = temp
	}

	for _, value := range constants.AchievementFamilies {
		if value.StickyNote != nil && value.Achieved() {
			rawAchievements := funk.Map(constants.Achievements, func(k string, value constants.Achievement) constants.Achievement {
				return value
			})
			filteredAchievements := funk.Filter(rawAchievements, func(x constants.Achievement) bool {
				return x.MyFamily.Name == value.Name
			}).([]constants.Achievement)

			for _, achievement := range filteredAchievements {
				if achievement.Achieved && !achievement.Presented {
					achievement.Presented = true
					constants.Achievements[achievement.Name] = achievement
					OpenSticky(&data.StickyMsg{
						Message: value.String(),
						Offset:  constants.NoteVec,
					})
					achievement.Presented = true
					constants.Achievements[achievement.Name] = achievement
				}
			}
			value.StickyNote.Hide = false
		}
	}
}

func ClickAchievement(a *constants.AchievementFamily) func() {
	return func() {
		OpenSticky(&data.StickyMsg{
			Message: a.String(),
			Offset:  constants.NoteVec,
		})
	}
}
