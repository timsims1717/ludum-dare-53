package data

import "timsims1717/ludum-dare-53/pkg/timing"

type Funky struct {
	Fn func()
}

func NewFn(fn func()) *Funky {
	return &Funky{Fn: fn}
}

type TimerFunc struct {
	Timer *timing.Timer
	Func  func() bool
}

func NewTimerFunc(fn func() bool, dur float64) *TimerFunc {
	return &TimerFunc{
		Timer: timing.New(dur),
		Func:  fn,
	}
}

type FrameFunc struct {
	Func func() bool
}

func NewFrameFunc(fn func() bool) *FrameFunc {
	return &FrameFunc{Func: fn}
}
