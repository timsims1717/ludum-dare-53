package systems

import (
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/util"
	"timsims1717/ludum-dare-53/pkg/world"
)

var (
	PieceDone     bool
	FailCondition bool
	FailReason    constants.FailCondition
)

func TetrisSystem() {
	if FailCondition {
		return
	}
	if PieceDone {
		// check for rows
		var fullRows []int
		for y, row := range data.TetrisBoard.Board {
			full := true
			for _, block := range row {
				if block == nil {
					full = false
					break
				}
			}
			if full {
				fullRows = append(fullRows, y)
				data.TetrisBoard.Stats.AddToTetrisStats(1)
			}
		}
		if len(fullRows) == 0 {
			data.TetrisBoard.Stats.ResetStreak()
		}
		down := 0
		for y, row := range data.TetrisBoard.Board {
			if util.Contains(y, fullRows) {
				for x, block := range row {
					myecs.Manager.DisposeEntity(block.Entity)
					data.TetrisBoard.Board[y][x] = nil
				}
				down++
			} else if down > 0 {
				for _, block := range row {
					if block != nil {
						data.TetrisBoard.Set(block.Coords, nil)
						block.Coords.Y -= down
						data.TetrisBoard.Set(block.Coords, block)
					}
				}
			}
		}
		// create new piece
		if HasTetromino() {
			if PlaceTetromino() {
				PieceDone = false
			} else {
				FailCondition = true
				FailReason = constants.BoardFull
			}
		} else {
			PieceDone = true
			if data.TetrisBoard.Stats.Tetrominos > constants.MinPiecesToFail && !constants.IgnoreEmptyConv {
				empty := true
				for _, tet := range data.Conveyor.Tets {
					if tet != nil {
						empty = false
					}
				}
				FailCondition = empty
				if FailCondition {
					FailReason = constants.ConveyorBeltEmpty
				}
			}
		}
	}
}

func ClearBoard() {
	for y, row := range data.TetrisBoard.Board {
		for x, block := range row {
			if block != nil {
				myecs.Manager.DisposeEntity(block.Entity)
				data.TetrisBoard.Board[y][x] = nil
			}
		}
	}
	data.TetrisBoard.Shape = nil
	data.TetrisBoard.Speed = constants.DefaultSpeed
	data.TetrisBoard.ConvSpd = constants.ConvSpdMin
	data.TetrisBoard.NextShape = NewTetromino()
	data.TetrisBoard.Stats.FullReset()
	PlaceTetromino()
}

func ClearFactory() {
	for _, pad := range data.FactoryPads {
		if pad.Tet != nil {
			for _, block := range pad.Tet.Blocks {
				myecs.Manager.DisposeEntity(block.Entity)
			}
			myecs.Manager.DisposeEntity(pad.Tet.Entity)
			pad.Tet = nil
		}
	}
	for y, row := range data.FactoryFloor.Blocks {
		for x, block := range row {
			if block != nil {
				myecs.Manager.DisposeEntity(block.Entity)
				data.FactoryFloor.Set(world.Coords{X: x, Y: y}, nil)
			}
		}
	}
	if data.DraggingPiece != nil {
		for _, block := range data.DraggingPiece.Blocks {
			myecs.Manager.DisposeEntity(block.Entity)
		}
		myecs.Manager.DisposeEntity(data.DraggingPiece.Entity)
		data.DraggingPiece = nil
	}
	for i, tet := range data.Conveyor.Tets {
		if tet != nil {
			for _, block := range tet.Blocks {
				myecs.Manager.DisposeEntity(block.Entity)
			}
			myecs.Manager.DisposeEntity(tet.Entity)
			data.Conveyor.Tets[i] = nil
		}
	}
}
