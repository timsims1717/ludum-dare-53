package loading

import (
	"fmt"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/reanimator"
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

	stickyNote, err := img.LoadImage("assets/stickynote.png")
	if err != nil {
		panic(err)
	}
	data.StickyNote = pixel.NewSprite(stickyNote, stickyNote.Bounds())
	data.StickyObj = object.New()
	data.StickyObj.Rect = stickyNote.Bounds()
	data.TinyNote = img.NewOffsetSprite("sticky_note", constants.FactoryKey, pixel.V(0, 0))

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
	for x := 1; x < 26; x++ {
		spr = img.NewOffsetSprite("wall_in", constants.FactoryKey, pixel.V(float64(x)*data.MSize, 0.))
		data.SideSection = append(data.SideSection, spr)
	}
	// corner wall
	spr = img.NewOffsetSprite("wall_corner", constants.FactoryKey, pixel.V(0., 0.))
	data.CornerSection = append(data.CornerSection, spr)
	for x := 1; x < 26; x++ {
		spr = img.NewOffsetSprite("wall_in", constants.FactoryKey, pixel.V(float64(x)*data.MSize, 0.))
		data.CornerSection = append(data.CornerSection, spr)
	}
	for y := 0; y < 35; y++ {
		for x := 0; x < 27; x++ {
			spr = img.NewOffsetSprite("wall_in", constants.FactoryKey, pixel.V(float64(x)*data.MSize, float64(y+1)*data.MSize))
			data.CornerSection = append(data.CornerSection, spr)
		}
	}
	// side doors
	for y := 0; y < 12; y++ {
		str := "caution_side_1"
		if y%2 == 0 {
			str = "caution_side_2"
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
	for x := 0; x < data.BeltSize; x++ {
		str := "conv_base_m"
		if x == 0 {
			str = "conv_base_l"
		} else if x == data.BeltSize-1 {
			str = "conv_base_r"
		}
		spr = img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x-53)*data.MSize, 0))
		data.ConveyorBase = append(data.ConveyorBase, spr)
	}
	// conveyor sections
	for x := 0; x < 3; x++ {
		var str string
		if x == 0 {
			str = "conv_l%s"
		} else if x == 1 {
			str = "conv_m%s"
		} else {
			str = "conv_r%s"
		}
		for y := 0; y < 6; y++ {
			key := fmt.Sprintf(str, "")
			if y == 0 {
				key = fmt.Sprintf(str, "b")
			} else if y == 5 {
				key = fmt.Sprintf(str, "t")
			}
			sprs := img.Reverse(img.Batchers[constants.FactoryKey].GetAnimation(key).S)
			anim := reanimator.NewAnimFromSprites(key, sprs, reanimator.Loop).
				WithBatch(constants.FactoryKey).
				WithOffset(pixel.V(0, float64(y)*data.MSize))
			tree := reanimator.NewSimple(anim)
			if x == 0 {
				data.ConvLeftEdge = append(data.ConvLeftEdge, tree)
			} else if x == 1 {
				data.ConvMiddle = append(data.ConvMiddle, tree)
			} else {
				data.ConvRightEdge = append(data.ConvRightEdge, tree)
			}
		}
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, object.New()).
		AddComponent(myecs.Animation, data.ConvLeftEdge)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, object.New()).
		AddComponent(myecs.Animation, data.ConvMiddle)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, object.New()).
		AddComponent(myecs.Animation, data.ConvRightEdge)
	// trucks
	for y := 0; y < 10; y++ {
		for x := 0; x < data.TruckWidth; x++ {
			str := "truck_roof"
			if y == 0 {
				str = "truck_roof_top_edge"
			}
			spr = img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x-data.TruckWidth/2)*data.MSize, float64(-y)*data.MSize))
			data.BotTruck = append(data.BotTruck, spr)
		}
	}
	for y := 0; y < data.TruckHeight; y++ {
		for x := 0; x < 12; x++ {
			var str string
			if y == 0 {
				str = "truck_side_b_mid"
				if x == 0 {
					str = "truck_side_b_end"
				}
			} else if y < 7 {
				str = "truck_side_mid"
				if x == 0 {
					str = "truck_side_end"
				}
			} else if y == 7 {
				str = "truck_side_roof_b_mid"
				if x == 0 {
					str = "truck_side_roof_b_end"
				}
			} else if y == 11 {
				str = "truck_side_roof_t_mid"
				if x == 0 {
					str = "truck_side_roof_t_end"
				}
			} else {
				str = "truck_side_roof_m_mid"
				if x == 0 {
					str = "truck_side_roof_m_end"
				}
			}
			spr = img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x)*data.MSize, float64(y-data.TruckHeight/2)*data.MSize))
			data.MidTruck = append(data.MidTruck, spr)
		}
	}
	for y := 0; y < 20; y++ {
		for x := 0; x < data.TruckWidth; x++ {
			var str string
			if y == 0 {
				str = "truck_back_b"
				if x == 0 {
					str = "truck_back_bl"
				} else if x == data.TruckWidth-1 {
					str = "truck_back_br"
				} else if x == 1 || x == data.TruckWidth-2 {
					str = "truck_back_bf"
				}
			} else if y < 8 {
				str = "truck_back"
				if x == 0 {
					str = "truck_back_l"
				} else if x == data.TruckWidth-1 {
					str = "truck_back_r"
				}
			} else if y == 8 {
				str = "truck_back_t"
				if x == 0 {
					str = "truck_back_tl"
				} else if x == data.TruckWidth-1 {
					str = "truck_back_tr"
				} else if x == data.TruckWidth/2-1 {
					str = "truck_back_thl"
				} else if x == data.TruckWidth/2 {
					str = "truck_back_thr"
				}
			} else if y == 11 {
				str = "truck_roof_bot_edge"
			} else {
				str = "truck_roof"
			}
			spr = img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x-data.TruckWidth/2)*data.MSize, float64(y)*data.MSize))
			data.TopTruck = append(data.TopTruck, spr)
		}
	}
	// TV
	h := 18
	w := 20
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			str := "tv_m"
			if y == 0 {
				str = "tv_b"
				if x == 0 {
					str = "tv_bl"
				} else if x == w-1 {
					str = "tv_br"
				}
			} else if y == h-1 {
				str = "tv_t"
				if x == 0 {
					str = "tv_tl"
				} else if x == w-1 {
					str = "tv_tr"
				}
			} else {
				if x == 0 {
					str = "tv_ml"
				} else if x == w-1 {
					str = "tv_mr"
				}
			}
			spr = img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x)*data.MSize, float64(y)*data.MSize))
			data.TV = append(data.TV, spr)
		}
	}
	// Platform
	pw := 23
	ph := 44
	tnb := 32
	tnt := 40
	for y := 0; y < ph; y++ {
		for x := 0; x < pw; x++ {
			var str string
			if y < 3 {
				str = fmt.Sprintf("bulb_mb%d", y+1)
				if x == 0 {
					str = fmt.Sprintf("bulb_r%dc1", y+1)
				} else if x == 1 {
					str = fmt.Sprintf("bulb_r%dc2", y+1)
				} else if x == pw-2 {
					str = fmt.Sprintf("bulb_r%dc3", y+1)
				} else if x == pw-1 {
					str = fmt.Sprintf("bulb_r%dc4", y+1)
				}
			} else if y < tnb {
				if x == 0 {
					str = "bulb_r4c1"
				} else if x == 1 {
					str = "bulb_r4c2"
				} else if x == pw-2 {
					str = "bulb_r4c3"
				} else if x == pw-1 {
					str = "bulb_r4c4"
				}
			} else if y == tnb {
				if x == 0 {
					str = "bulb_r5c1"
				} else if x == 1 {
					str = "bulb_r5c2"
				} else if x == pw-2 {
					str = "bulb_r5c3"
				} else if x == pw-1 {
					str = "bulb_r5c4"
				}
			} else if y == tnb+1 {
				if x == 1 {
					str = "bulb_r6c2"
				} else if x == pw-2 {
					str = "bulb_r6c3"
				}
			} else if y > tnb+1 && y < tnt {
				if x == 1 {
					str = "bulb_r7c2"
				} else if x == pw-2 {
					str = "bulb_r7c3"
				}
			} else if y == tnt {
				if x == 0 {
					str = "bulb_r8c1"
				} else if x == 1 {
					str = "bulb_r8c2"
				} else if x == pw-2 {
					str = "bulb_r8c3"
				} else if x == pw-1 {
					str = "bulb_r8c4"
				}
			} else if y > tnt && y < ph-1 {
				if x == 0 {
					str = "bulb_r9c1"
				} else if x == 1 {
					str = "bulb_r9c2"
				} else if x == pw-2 {
					str = "bulb_r9c3"
				} else if x == pw-1 {
					str = "bulb_r9c4"
				}
			} else if y == ph-1 {
				str = "bulb_mt"
				if x == 0 {
					str = "bulb_r0c1"
				} else if x == 1 {
					str = "bulb_r0c2"
				} else if x == pw-2 {
					str = "bulb_r0c3"
				} else if x == pw-1 {
					str = "bulb_r0c4"
				}
			}
			if str != "" {
				spr = img.NewOffsetSprite(str, constants.FactoryKey, pixel.V(float64(x)*data.MSize, float64(y)*data.MSize))
				data.Bulb = append(data.Bulb, spr)
			}
		}
	}
	// buttons
	baseSpr1 := img.NewOffsetSprite("button_base_1", constants.BlockKey, pixel.V(0., 0.))
	pauseBtn := img.NewOffsetSprite("pause_button", constants.BlockKey, pixel.V(0., 0.))
	restartBtn := img.NewOffsetSprite("restart_button", constants.BlockKey, pixel.V(0., 0.))
	baseSpr2 := img.NewOffsetSprite("button_base_2", constants.BlockKey, pixel.V(0., 0.))
	data.RestartButSprs = append(data.RestartButSprs, baseSpr1)
	data.RestartButSprs = append(data.RestartButSprs, restartBtn)
	data.RestartButSprs = append(data.RestartButSprs, baseSpr2)
	data.PauseButSprs = append(data.PauseButSprs, baseSpr1)
	data.PauseButSprs = append(data.PauseButSprs, pauseBtn)
	data.PauseButSprs = append(data.PauseButSprs, baseSpr2)
	// tetris display
	data.TVShapes = append(data.TVShapes, img.NewOffsetSprite("O", constants.FactoryKey, pixel.V(-6.*data.MSize, 0.)))
	data.TVShapes = append(data.TVShapes, img.NewOffsetSprite("S", constants.FactoryKey, pixel.V(-4.*data.MSize, 0.)))
	data.TVShapes = append(data.TVShapes, img.NewOffsetSprite("Z", constants.FactoryKey, pixel.V(-2.*data.MSize, 0.)))
	data.TVShapes = append(data.TVShapes, img.NewOffsetSprite("J", constants.FactoryKey, pixel.V(0., 0.)))
	data.TVShapes = append(data.TVShapes, img.NewOffsetSprite("L", constants.FactoryKey, pixel.V(2.*data.MSize, 0.)))
	data.TVShapes = append(data.TVShapes, img.NewOffsetSprite("T", constants.FactoryKey, pixel.V(4.*data.MSize, 0.)))
	data.TVShapes = append(data.TVShapes, img.NewOffsetSprite("I", constants.FactoryKey, pixel.V(6.*data.MSize, 0.)))

	// cursor
	data.HandPoint = img.NewOffsetSprite("hand_point", constants.FactoryKey, pixel.ZV)
	data.HandOpen = img.NewOffsetSprite("hand_open", constants.FactoryKey, pixel.ZV)
	data.HandGrab = img.NewOffsetSprite("hand_grab", constants.FactoryKey, pixel.ZV)
}
