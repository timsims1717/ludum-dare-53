package data

type TetrisScore struct {
	Score      int
	Streak     int
	MyFibScore *FibScore
}

func newTetrisScore() *TetrisScore {
	tScore := &TetrisScore{Score: 0, Streak: 0, MyFibScore: newFibScore()}
	return tScore
}

func (ts *TetrisScore) AddToScore() {
	ts.Streak++
	ts.Score = ts.Score + ts.MyFibScore.fibIter(1)
}

func (ts *TetrisScore) ResetStreak() {
	ts.Streak = 0
	ts.MyFibScore.reset()
}

func (ts *TetrisScore) FullReset() {
	ts.ResetStreak()
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
