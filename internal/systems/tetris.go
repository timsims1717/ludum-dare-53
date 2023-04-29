package systems

import (
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/util"
)

var (
	PieceDone     bool
	FailCondition bool
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
				data.TetrisBoard.Score.AddToScore()
			}
		}
		if len(fullRows) == 0 {
			data.TetrisBoard.Score.ResetStreak()
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
		FailCondition = !CreateTetronimo()
	}
	PieceDone = false
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
	CreateTetronimo()
}
