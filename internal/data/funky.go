package data

type Funky struct {
	Fn func()
}

func NewFn(fn func()) *Funky {
	return &Funky{Fn: fn}
}
