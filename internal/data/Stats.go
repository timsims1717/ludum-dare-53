package data

import (
	"timsims1717/ludum-dare-53/internal/constants"
)

type TetrisStats struct {
	Score         int
	LinesCleared  int
	Tetrominos    int
	Streak        int
	LongestStreak int
	Checkpoint    int
	MyFibScore    *FibScore
}

func (ts *TetrisStats) GlobalScore() int {
	return TetrisBoard.Stats.Score + FactoryFloor.Stats.Score
}

type FactoryStats struct {
	Score                int
	Factrominos          int
	ColorStreak          int
	CurrentColor         TColor
	ShapeStreak          int
	LastTetromino        constants.TetronimoType
	LongestColorStreak   int
	LongestShapeStreak   int
	UnoStreak            int
	BalanceStreak        int
	LongestBalanceStreak int
	LongestUnoSteak      int
	ShapesTrashed        int
	LargestShape         int
	MyFibScore           FibScore
	TimesSinceLastShape  map[constants.TetronimoType]int
}

func newFactoryStats() *FactoryStats {
	tScore := &FactoryStats{Score: 0, Factrominos: 0, ColorStreak: 0, CurrentColor: 0, ShapeStreak: 0, LastTetromino: constants.UndefinedTetronimoType}
	tScore.TimesSinceLastShape = map[constants.TetronimoType]int{
		constants.I: 0,
		constants.O: 0,
		constants.T: 0,
		constants.S: 0,
		constants.Z: 0,
		constants.J: 0,
		constants.L: 0,
	}
	tScore.MyFibScore = *newFibScore()
	return tScore
}

func (fs *FactoryStats) AddToFactoryStats(factromino Factromino) {
	fs.Factrominos++
	timeSinceLastShape := 0
	for key, value := range fs.TimesSinceLastShape {
		if key == factromino.MyTetronimoType {
			timeSinceLastShape = value
			fs.TimesSinceLastShape[key] = 0
		} else {
			fs.TimesSinceLastShape[key]++
		}
	}
	if timeSinceLastShape >= constants.BalanceStreakTarget {
		fs.BalanceStreak++
		fs.Score += fs.MyFibScore.fibIter()
	} else {
		fs.ResetFactoryBalanceStreak()
		fs.Score += 1
	}
	if factromino.Color == fs.CurrentColor || factromino.MyTetronimoType == fs.LastTetromino {
		fs.UnoStreak++
	} else {
		fs.UnoStreak = 0
	}
	if factromino.Color == fs.CurrentColor {
		fs.ColorStreak++
	} else {
		if fs.ColorStreak > fs.LongestColorStreak {
			fs.LongestColorStreak = fs.ColorStreak
		}
		fs.ColorStreak = 0
		fs.CurrentColor = factromino.Color
	}
	if factromino.MyTetronimoType == fs.LastTetromino {
		fs.ShapeStreak++
	} else {
		if fs.ShapeStreak > fs.LongestShapeStreak {
			fs.LongestShapeStreak = fs.ShapeStreak
		}
		fs.ShapeStreak = 0
		fs.LastTetromino = factromino.MyTetronimoType
	}
}
func (fs *FactoryStats) ResetFactoryBalanceStreak() {
	fs.BalanceStreak = 0
	fs.MyFibScore.reset()
}

func (fs *FactoryStats) FullFactoryStatReset() {
	fs.Score = 0
	fs.Factrominos = 0
	fs.ColorStreak = 0
	fs.CurrentColor = 0
	fs.ShapeStreak = 0
	fs.LastTetromino = constants.UndefinedTetronimoType
	fs.LongestColorStreak = 0
	fs.LongestShapeStreak = 0
	fs.ResetFactoryBalanceStreak()
	fs.UnoStreak = 0
	fs.LongestUnoSteak = 0
	fs.ShapesTrashed = 0
	fs.LargestShape = 0
	fs.TimesSinceLastShape = map[constants.TetronimoType]int{
		constants.I: 0,
		constants.O: 0,
		constants.T: 0,
		constants.S: 0,
		constants.Z: 0,
		constants.J: 0,
		constants.L: 0,
	}
}

func newTetrisStats() *TetrisStats {
	tScore := &TetrisStats{Score: 0, Streak: 0, MyFibScore: newFibScore(), Tetrominos: 0}
	return tScore
}

func (ts *TetrisStats) AddToTetrisStats(clearedRows int) {
	ts.Streak++
	ts.LinesCleared += clearedRows
	for itr := 0; itr < clearedRows; itr++ {
		ts.Score += ts.MyFibScore.fibIter()
	}
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

func (f *FibScore) fibIter() int {
	score := f.FibN
	fibf := f.FibNMinus + f.FibN
	f.FibNMinus = f.FibN
	f.FibN = fibf
	return score
}

func (f *FibScore) reset() {
	f.FibN = 1
	f.FibNMinus = 0
}
