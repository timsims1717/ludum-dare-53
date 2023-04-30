package loading

import (
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/pkg/img"
)

func LoadImg() {
	// blocks
	blockSheet, err := img.LoadSpriteSheet("assets/blocks.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(constants.BlockKey, blockSheet, true, true)

	// factory
	factorySheet, err := img.LoadSpriteSheet("assets/factory.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(constants.FactoryKey, factorySheet, true, true)

	// tilemaps
	// floor section
	for y := 0; y < 7; y++ {
		for x := 0; x < 11; x++ {
			str := "concrete_4"
			if x == 0 && y == 6 {
				str = "concrete_1"
			} else if x == 0 {
				str = "concrete_2"
			} else if y == 6 {
				str = "concrete_3"
			}
			spr := img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x)*data.MSize, float64(y)*data.MSize))
			data.FloorSection = append(data.FloorSection, spr)
		}
	}
	// tile pad (grate)
	for y := 0; y < 8; y++ {
		for x := 0; x < 12; x++ {
			str := "grate_m"
			if y == 0 {
				if x == 0 {
					str = "grate_bl"
				} else if x == 11 {
					str = "grate_br"
				} else {
					str = "grate_b"
				}
			} else if y == 7 {
				if x == 0 {
					str = "grate_tl"
				} else if x == 11 {
					str = "grate_tr"
				} else {
					str = "grate_t"
				}
			} else if y%2 == 0 {
				if x == 0 {
					str = "grate_cl"
				} else if x == 11 {
					str = "grate_cr"
				} else {
					str = "grate_c"
				}
			} else {
				if x == 0 {
					str = "grate_ml"
				} else if x == 11 {
					str = "grate_mr"
				} else {
					str = "grate_m"
				}
			}
			spr := img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x-6)*data.MSize, float64(y-4)*data.MSize))
			data.PadSection = append(data.PadSection, spr)
		}
	}
	// wall strip
	for y := 0; y < 28; y++ {
		for x := 0; x < 2; x++ {
			str := "wall_base"
			if y == 0 {
				str = "wall_base"
			} else if y == 14 {
				str = "wall_t"
			} else if y > 14 {
				str = "wall_in"
			} else if y%2 == 0 {
				str = "wall_2"
			} else {
				str = "wall_1"
			}
			spr := img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x)*data.MSize, float64(y)*data.MSize))
			data.WallSection = append(data.WallSection, spr)
		}
	}
	// back door
	for y := 0; y < 28; y++ {
		for x := 0; x < 16; x++ {
			str := ""
			if y > 10 || x > 13 {
				if y == 0 {
					str = "wall_base"
				} else if y == 14 {
					str = "wall_t"
				} else if y > 14 {
					str = "wall_in"
				} else if y%2 == 0 {
					str = "wall_2"
				} else {
					str = "wall_1"
				}
			} else if y == 10 {
				if x == 0 {
					str = "caution_1e"
				} else if x == 13 {
					str = "caution_2f"
				} else if x%2 == 0 {
					str = "caution_1d"
				} else {
					str = "caution_2d"
				}
			} else if x == 0 {
				if y%2 == 0 {
					str = "caution_1b"
				} else {
					str = "caution_2b"
				}
			} else if x == 13 {
				if y%2 == 0 {
					str = "caution_2c"
				} else {
					str = "caution_1c"
				}
			}
			if str != "" {
				spr := img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x)*data.MSize, float64(y)*data.MSize))
				data.DoorSection = append(data.DoorSection, spr)
			}
		}
	}
	// side wall
	spr := img.NewOffsetSprite("wall_side", constants.FactoryKey, pixel.V(0., 0.))
	data.SideSection = append(data.SideSection, spr)
	for x := 1; x < 8; x++ {
		spr = img.NewOffsetSprite("wall_in", constants.FactoryKey, pixel.V(float64(x)*data.MSize, 0.))
		data.SideSection = append(data.SideSection, spr)
	}
	// side doors
	for y := 0; y < 12; y++ {
		str := "caution_side_1"
		if y%2 == 0 {
			str = "caution_side_1"
		}
		spr = img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(0, float64(y)*data.MSize))
		data.SideDSection = append(data.SideDSection, spr)
	}
	// block spot
	for y := 0; y < 2; y++ {
		for x := 0; x < 3; x++ {
			str := ""
			if y == 0 {
				if x == 0 {
					str = "grate_bl"
				} else if x == 1 {
					str = "grate_b"
				} else {
					str = "grate_br"
				}
			} else {
				if x == 0 {
					str = "grate_tl"
				} else if x == 1 {
					str = "grate_t"
				} else {
					str = "grate_tr"
				}
			}
			spr = img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x-1)*data.MSize, (float64(y)-0.5)*data.MSize))
			data.BlockSpot = append(data.BlockSpot, spr)
		}
	}
	// conveyor base
	for x := 0; x < 53; x++ {
		str := "conv_base_m"
		if x == 0 {
			str = "conv_base_l"
		} else if x == 52 {
			str = "conv_base_r"
		}
		spr = img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x-53)*data.MSize, 0))
		data.ConveyorBase = append(data.ConveyorBase, spr)
	}
}
