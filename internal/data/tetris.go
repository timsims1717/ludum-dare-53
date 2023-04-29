package data

import (
	"github.com/bytearena/ecs"
	"math/rand"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/pkg/timing"
	"timsims1717/ludum-dare-53/pkg/world"
)

type TColor int

const (
	Green = iota
	Orange
	Indigo
	Red
	Teal
	Yellow
	Pink
	Black
)

func (c TColor) String() string {
	switch c {
	case Green:
		return "green"
	case Orange:
		return "orange"
	case Indigo:
		return "indigo"
	case Red:
		return "red"
	case Teal:
		return "teal"
	case Yellow:
		return "yellow"
	case Pink:
		return "pink"
	case Black:
		return "black"
	}
	return ""
}

func RandColor() TColor {
	return TColor(rand.Intn(Black))
}

type Tetronimo struct {
	Blocks []*TetrisBlock
	NoRot  bool
}

type TetrisBlock struct {
	Coords world.Coords
	Color  TColor
	Moving bool
	Entity *ecs.Entity
}

var TetrisBoard *tetrisBoard

type tetrisBoard struct {
	Board [constants.TetrisHeight][constants.TetrisWidth]*TetrisBlock
	Shape *Tetronimo
	Timer *timing.Timer
	Speed float64
}

func (t *tetrisBoard) Get(c world.Coords) *TetrisBlock {
	return t.Board[c.Y][c.X]
}

func (t *tetrisBoard) Set(c world.Coords, b *TetrisBlock) {
	t.Board[c.Y][c.X] = b
}

func (t *tetrisBoard) ResetTimer() {
	t.Timer = timing.New(t.Speed)
}

func NewTetrisBoard(spd float64) {
	TetrisBoard = &tetrisBoard{
		Board: [constants.TetrisHeight][constants.TetrisWidth]*TetrisBlock{},
		Shape: nil,
		Timer: timing.New(spd),
		Speed: spd,
	}
}
