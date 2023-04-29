package world

import (
	"github.com/faiface/pixel"
)

var (
	TileSize float64
	Origin   = Coords{
		X: 0,
		Y: 0,
	}
	TileRect pixel.Rect
)

func SetTileSize(s float64) {
	TileSize = s
	TileRect = pixel.R(0, 0, s, s)
}

func MapToWorld(a Coords) pixel.Vec {
	return MapToWorldC(a, pixel.V(TileSize, TileSize))
}

func WorldToMap(x, y float64) (int, int) {
	return WorldToMapC(x, y, pixel.V(TileSize, TileSize))
}

func MapToWorldC(a Coords, size pixel.Vec) pixel.Vec {
	return pixel.V(float64(a.X)*size.X, float64(a.Y)*size.Y)
}

func WorldToMapC(x, y float64, size pixel.Vec) (int, int) {
	return int(x / size.X), int(y / size.Y)
}
