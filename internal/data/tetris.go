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
	Blocks  []*TetrisBlock
	NoRot   bool
	TetType TetronimoType
}

func (t Tetronimo) IsValid() bool {
	for _, b := range t.Blocks {
		if b == nil {
			return false
		}
	}
	return true
}

type TetrisBlock struct {
	Coords world.Coords
	Color  TColor
	Moving bool
	Entity *ecs.Entity
}

type TetronimoType int

const (
	I = iota
	O
	T
	S
	Z
	J
	L
)

func (t TetronimoType) String() string {
	switch t {
	case O:
		return "O"
	case I:
		return "I"
	case S:
		return "S"
	case Z:
		return "Z"
	case L:
		return "L"
	case J:
		return "J"
	case T:
		return "T"
	}
	return ""
}

var TetrisBoard *tetrisBoard

type tetrisBoard struct {
	Board     [constants.TetrisHeight][constants.TetrisWidth]*TetrisBlock
	Shape     *Tetronimo
	NextShape *Tetronimo
	Timer     *timing.Timer
	Speed     float64
	Stats     *TetrisStats
}

func (t *tetrisBoard) Get(c world.Coords) *TetrisBlock {
	return t.Board[c.Y][c.X]
}
func (t *tetrisBoard) SpeedUp() {
	if t.Speed > constants.SpeedMin {
		if t.Speed <= constants.HighSpeedMark {
			t.Speed = t.Speed - constants.HighSpeedModifer
		} else {
			t.Speed = t.Speed - constants.SpeedModifier
		}
	}
}
func (t *tetrisBoard) SpeedDown() {
	if t.Speed < constants.SpeedMax {
		if t.Speed <= constants.HighSpeedMark {
			t.Speed = t.Speed + constants.HighSpeedModifer
		} else {
			t.Speed = t.Speed + constants.SpeedModifier
		}
	}
}

func (t *tetrisBoard) Set(c world.Coords, b *TetrisBlock) {
	t.Board[c.Y][c.X] = b
}

func (t *tetrisBoard) ResetTimer() {
	t.Timer = timing.New(t.Speed)
}

func NewTetrisBoard(spd float64) {
	TetrisBoard = &tetrisBoard{
		Board: [20][10]*TetrisBlock{},
		Shape: nil,
		Timer: timing.New(spd),
		Speed: spd,
		Stats: newTetrisStats(),
	}
}
