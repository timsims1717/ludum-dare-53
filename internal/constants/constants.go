package constants

import (
	"image/color"
	"timsims1717/ludum-dare-53/pkg/world"
)

const (
	Title   = "LD53"
	Release = 0
	Version = 1
	Build   = 20230428

	// Batches
	BlockKey = "blocks"

	// Tetris
	TileSize         = 32.
	TetrisWidth      = 10
	TetrisHeight     = 20
	DefaultSpeed     = 1
	ScoreCheckPoint  = 5
	SpeedModifier    = 0.05
	SpeedMax         = 2
	SpeedMin         = 0.1
	HighSpeedModifer = 0.01
	HighSpeedMark    = 0.3

	// Factory
	FactoryTile   = 48.
	FactoryWidth  = 5
	FactoryHeight = 7
)

var (
	BlackColor = color.RGBA{
		R: 19,
		G: 19,
		B: 19,
		A: 255,
	}
	TetrisStart          = world.Coords{X: 4, Y: 19}
	NormalizedTetronimos = map[[4]world.Coords]TetronimoType{
		[4]world.Coords{{0, 0}, {1, 0}, {2, 0}, {3, 0}}: I,
		[4]world.Coords{{0, 0}, {0, 1}, {0, 2}, {0, 3}}: I,
		[4]world.Coords{{0, 0}, {1, 0}, {0, 1}, {1, 1}}: O,
		[4]world.Coords{{0, 0}, {1, 0}, {2, 0}, {1, 1}}: T, //Point up, flat down
		[4]world.Coords{{0, 0}, {0, 1}, {0, 2}, {1, 1}}: T, //Point Right, Flat Left
		[4]world.Coords{{1, 0}, {0, 1}, {1, 1}, {2, 1}}: T, //Point Left, Flat Right
		[4]world.Coords{{1, 0}, {1, 1}, {1, 2}, {0, 1}}: T, //Point Down, Flat Up
		[4]world.Coords{{0, 1}, {1, 1}, {1, 0}, {0, 2}}: Z,
		[4]world.Coords{{0, 0}, {0, 1}, {1, 1}, {1, 2}}: Z,
		[4]world.Coords{{0, 0}, {1, 0}, {1, 1}, {2, 1}}: S,
		[4]world.Coords{{1, 0}, {2, 0}, {0, 1}, {1, 1}}: S,
		[4]world.Coords{{0, 0}, {1, 0}, {2, 0}, {2, 1}}: L,
		[4]world.Coords{{0, 0}, {1, 0}, {0, 1}, {0, 2}}: L,
		[4]world.Coords{{0, 2}, {1, 0}, {1, 1}, {1, 2}}: L,
		[4]world.Coords{{1, 0}, {1, 1}, {2, 1}, {0, 2}}: L,
		[4]world.Coords{{0, 0}, {1, 0}, {2, 0}, {0, 1}}: J,
		[4]world.Coords{{0, 0}, {0, 1}, {0, 2}, {1, 2}}: J,
		[4]world.Coords{{0, 1}, {1, 1}, {2, 1}, {0, 2}}: J,
		[4]world.Coords{{0, 0}, {1, 0}, {1, 1}, {1, 2}}: J,
	}
)

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
