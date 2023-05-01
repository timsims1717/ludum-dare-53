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

type Tetromino struct {
	Blocks  []*TetrisBlock
	NoRot   bool
	TetType constants.TetronimoType
}

func FacTetIsanI(f *Factromino) bool {
	xchange := false
	ychange := false
	x := f.Blocks[0].Coords.X
	y := f.Blocks[0].Coords.Y
	for _, block := range f.Blocks {
		xchange = x != block.Coords.X
		ychange = y != block.Coords.Y
	}
	return !xchange || !ychange
}

func Normalize(coords [4]world.Coords) [4]world.Coords {
	minX, minY := coords[0].X, coords[0].Y

	for _, coord := range coords {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
	}

	var normalized [4]world.Coords
	for i, coord := range coords {
		normalized[i] = world.Coords{coord.X - minX, coord.Y - minY}
	}
	return normalized
}
func TetronimoCoordsEqual(a, b [4]world.Coords) bool {
	for _, i := range a {
		if !CoordsIn(i, b) {
			return false
		}
	}
	return true
}

func CoordsIn(c world.Coords, list [4]world.Coords) bool {
	for _, l := range list {
		if c.Eq(l) {
			return true
		}
	}
	return false
}
func TetronimoCoordsRotate(coords []world.Coords) []world.Coords {
	rotated := make([]world.Coords, len(coords))
	for i, coord := range coords {
		rotated[i] = world.Coords{coord.X, -coord.Y}
	}
	return rotated
}

func (t Tetromino) IsValid() bool {
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

var TetrisBoard *tetrisBoard

type tetrisBoard struct {
	Board     [constants.TetrisHeight][constants.TetrisWidth]*TetrisBlock
	Shape     *Tetromino
	NextShape *Tetromino
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
		Board: [constants.TetrisHeight][constants.TetrisWidth]*TetrisBlock{},
		Shape: nil,
		Timer: timing.New(spd),
		Speed: spd,
		Stats: newTetrisStats(),
	}
}
