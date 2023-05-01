package data

import (
	"timsims1717/ludum-dare-53/internal/constants"
)

type TetrisStats struct {
	Score      int
	Tetrominos int
	Streak     int
	Checkpoint int
	MyFibScore *FibScore
}
type FactoryStats struct {
	Score         int
	Factrominos   int
	ColorStreak   int
	CurrentColor  TColor
	ShapeStreak   int
	LastTetromino constants.TetronimoType
}

func newTetrisStats() *TetrisStats {
	tScore := &TetrisStats{Score: 0, Streak: 0, MyFibScore: newFibScore(), Tetrominos: 0}
	return tScore
}

func (ts *TetrisStats) AddToScore() {
	ts.Streak++
	ts.Score = ts.Score + ts.MyFibScore.fibIter(1)
	ts.IncrementCheckpointAndSpeed()
}

func (ts *TetrisStats) IncrementCheckpointAndSpeed() {
	var checkpointtarget int
	checkpointtarget = ts.Score / constants.ScoreCheckPoint
	for i := ts.Checkpoint; i < checkpointtarget; i++ {
		ts.Checkpoint++
		TetrisBoard.SpeedUp()
	}
}

func (ts *TetrisStats) ResetStreak() {
	ts.Streak = 0
	ts.MyFibScore.reset()
}

func (ts *TetrisStats) FullReset() {
	ts.ResetStreak()
	ts.Tetrominos = 0
	ts.Score = 0
}

type FibScore struct {
	FibNMinus int
	FibN      int
}

func newFibScore() *FibScore {
	tFibScore := &FibScore{FibNMinus: 0, FibN: 1}
	return tFibScore
}

func (f *FibScore) fibIter(cycles int) int {
	score := 0
	for i := 0; i < cycles; i++ {
		fibf := f.FibNMinus + f.FibN
		f.FibNMinus = f.FibN
		f.FibN = fibf
		score = score + fibf
	}
	return score
}

func (f *FibScore) reset() {
	f.FibN = 1
	f.FibNMinus = 0
}