package loading

import (
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/pkg/img"
)

func LoadImg() {
	// blocks
	blockSheet, err := img.LoadSpriteSheet("assets/blocks.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(constants.BlockKey, blockSheet, true, true)

}
