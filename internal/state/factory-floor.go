package state

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/internal/systems"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/world"
)

var (
	FactoryBGEntities []*ecs.Entity
)

func BuildFactoryFloor() {
	for y := 0; y < constants.FactoryHeight; y++ {
		for x := 0; x < constants.FactoryWidth; x++ {
			obj := object.New()
			obj.Pos = world.MapToWorldC(world.Coords{X: x, Y: y}, pixel.V(constants.FactoryTile, world.TileSize))
			obj.Layer = 10
			e := myecs.Manager.NewEntity()
			e.AddComponent(myecs.Object, obj).
				AddComponent(myecs.Drawable, data.BlockSpot)
			FactoryBGEntities = append(FactoryBGEntities, e)
		}
	}
	data.FactoryFloor.Object = object.New()
	data.FactoryFloor.Entity = myecs.Manager.NewEntity()
	data.FactoryFloor.Object.Pos = pixel.V(constants.FactoryTile*0.5*(constants.FactoryWidth-1), world.TileSize*0.5*(constants.FactoryHeight-1))
	data.FactoryFloor.Object.Rect = pixel.R(0., 0., constants.FactoryWidth*constants.FactoryTile, constants.FactoryHeight*world.TileSize)
	data.FactoryFloor.Entity.AddComponent(myecs.Object, data.FactoryFloor.Object).
		AddComponent(myecs.ViewPort, data.FactoryViewport).
		AddComponent(myecs.Input, gameInput).
		AddComponent(myecs.Update, data.NewFn(FactoryFloorUpdate())).
		AddComponent(myecs.Click, data.NewFn(FactoryFloorClicked()))
}

func FactoryFloorUpdate() func() {
	return func() {
		if data.FactoryFloor.Object.PointInside(data.FactoryViewport.Projected(gameInput.World)) {
			if data.DraggingPiece != nil {
				if ActuallyOnFloor() {
					legal := true
					spr := img.NewSprite("ff_hlerr", constants.BlockKey)
					for _, block := range data.DraggingPiece.Blocks {
						pos := data.DraggingPiece.Object.Pos.Add(block.Object.Offset)
						x, y := world.WorldToMapC(pos.X+constants.FactoryTile*0.5, pos.Y, pixel.V(constants.FactoryTile, world.TileSize))
						c := world.Coords{X: x, Y: y}
						mPos := world.MapToWorldC(c, pixel.V(constants.FactoryTile, world.TileSize))
						if data.FactoryFloor.Get(c) != nil {
							legal = false
						}
						obj := object.New()
						obj.Pos = mPos
						obj.Layer = 19
						myecs.Manager.NewEntity().
							AddComponent(myecs.Object, obj).
							AddComponent(myecs.Drawable, spr).
							AddComponent(myecs.Temp, myecs.ClearFlag(true))
					}
					if legal {
						spr.Key = "ff_hl"
					}
				}
			} else {
				// todo: hover for picking up
			}
		}
	}
}

func FactoryFloorClicked() func() {
	return func() {
		if data.DraggingPiece != nil && !systems.FailCondition {
			if ActuallyOnFloor() {
				legal := true
				for _, block := range data.DraggingPiece.Blocks {
					pos := data.DraggingPiece.Object.Pos.Add(block.Object.Offset)
					x, y := world.WorldToMapC(pos.X+constants.FactoryTile*0.5, pos.Y, pixel.V(constants.FactoryTile, world.TileSize))
					c := world.Coords{X: x, Y: y}
					if data.FactoryFloor.Get(c) != nil {
						legal = false
						break
					}
				}
				if legal {
					for _, block := range data.DraggingPiece.Blocks {
						pos := data.DraggingPiece.Object.Pos.Add(block.Object.Offset)
						x, y := world.WorldToMapC(pos.X+constants.FactoryTile*0.5, pos.Y, pixel.V(constants.FactoryTile, world.TileSize))
						c := world.Coords{X: x, Y: y}
						mPos := world.MapToWorldC(c, pixel.V(constants.FactoryTile, world.TileSize))
						mPos.Y += 6.
						block.Object.Pos = mPos
						block.Coords = c
						block.Object.Offset = pixel.ZV
						block.Object.Layer = 19 - y
						block.Entity.RemoveComponent(myecs.Parent)
						data.FactoryFloor.Set(c, block)
					}
					myecs.Manager.DisposeEntity(data.DraggingPiece.Entity)
					data.DraggingPiece.Blocks = []*data.FactoryBlock{}
					data.DraggingPiece = nil

					if !constants.Achievements["GridFullOBlocks"].Achieved {
						total := 0
					empty:
						for _, row := range data.FactoryFloor.Blocks {
							for _, block := range row {
								if block != nil {
									total++
								} else {
									break empty
								}
							}
						}
						if total == 35 {
							temp := constants.Achievements["GridFullOBlocks"]
							temp.Achieved = true
							constants.Achievements["GridFullOBlocks"] = temp
						}
					}

					PlayPlaceSound()

				}
			}
		} else {
			pos := data.FactoryViewport.Projected(gameInput.World)
			x, y := world.WorldToMapC(pos.X+constants.FactoryTile*0.5, pos.Y, pixel.V(constants.FactoryTile, world.TileSize))
			c := world.Coords{X: x, Y: y}
			if data.FactoryLegal(c) {
				blockA := data.FactoryFloor.Get(c)
				if blockA != nil {
					blocks := []*data.FactoryBlock{blockA}
					blocks2 := GetAllColorNeighbors(blockA, blocks)
					blocks = GetAllColorNeighbors(blockA, blocks)
					tet := systems.ConstructTetFromBlocks(pos, blocks)
					tet.Entity.AddComponent(myecs.ViewPort, data.FactoryViewport)
					tet.Entity.AddComponent(myecs.Input, gameInput)
					data.DraggingPiece = tet
					data.DraggingPiece.Entity.AddComponent(myecs.Drag, &gameInput.World)
					data.DraggingPiece.Object.Layer = 20
					for _, block := range blocks2 {
						myecs.Manager.DisposeEntity(block.Entity)
						data.FactoryFloor.Set(block.Coords, nil)
					}
					PlayPickupSound()
				}
			}
		}
	}
}

func ActuallyOnFloor() bool {
	onFloor := true
	for _, block := range data.DraggingPiece.Blocks {
		pos := data.DraggingPiece.Object.Pos.Add(block.Object.Offset)
		if pos.X+constants.FactoryTile*0.5 < 0 || pos.Y < 0 {
			onFloor = false
			break
		}
		x, y := world.WorldToMapC(pos.X+constants.FactoryTile*0.5, pos.Y, pixel.V(constants.FactoryTile, world.TileSize))
		c := world.Coords{X: x, Y: y}
		if !data.FactoryLegal(c) {
			onFloor = false
			break
		}
	}
	return onFloor
}

func GetAllColorNeighbors(block *data.FactoryBlock, blocks []*data.FactoryBlock) []*data.FactoryBlock {
	col := block.Color
	for _, n := range block.Coords.Neighbors() {
		if data.FactoryLegal(n) {
			nBlock := data.FactoryFloor.Get(n)
			if nBlock != nil && nBlock.Color == col && !BlockInArray(nBlock, blocks) {
				blocks = append(blocks, nBlock)
				blocks = GetAllColorNeighbors(nBlock, blocks)
			}
		}
	}
	return blocks
}

func BlockInArray(block *data.FactoryBlock, blocks []*data.FactoryBlock) bool {
	for _, b := range blocks {
		if b.Coords == block.Coords {
			return true
		}
	}
	return false
}
