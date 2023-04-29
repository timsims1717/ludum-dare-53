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
	TileSize     = 32.
	TetrisWidth  = 10
	TetrisHeight = 20
	DefaultSpeed = 1

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
	TetrisStart = world.Coords{X: 4, Y: 19}
)
