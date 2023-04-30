package systems

import (
	"math/rand"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/world"
)

var (
	BlockUpdate bool
	MoveDown    bool
	MoveRight   bool
	MoveLeft    bool
	Rotate      bool
)

func BlockSystem() {
	if FailCondition {
		return
	}
	data.TetrisBoard.Timer.Update()
	if data.TetrisBoard.Shape != nil {
		if data.TetrisBoard.Timer.Done() || MoveDown {
			// check to see if any are at the bottom
			done := false
			for _, block := range data.TetrisBoard.Shape.Blocks {
				if block != nil {
					if !BlockLegal(block.Coords.D()) {
						done = true
						break
					}
					below := data.TetrisBoard.Get(block.Coords.D())
					if below != nil && !below.Moving {
						done = true
						break
					}
				}
			}
			if done {
				for _, block := range data.TetrisBoard.Shape.Blocks {
					if block != nil {
						block.Moving = false
					}
				}
				data.TetrisBoard.Shape = nil
				PieceDone = true
			} else {
				for _, block := range data.TetrisBoard.Shape.Blocks {
					if block != nil {
						data.TetrisBoard.Set(block.Coords, nil)
						block.Coords.Y--
					}
				}
				for _, block := range data.TetrisBoard.Shape.Blocks {
					if block != nil {
						data.TetrisBoard.Set(block.Coords, block)
					}
				}
			}
		}
		if MoveLeft && !PieceDone {
			canMove := true
			for _, block := range data.TetrisBoard.Shape.Blocks {
				if block != nil {
					if !BlockLegal(block.Coords.L()) {
						canMove = false
						break
					}
					left := data.TetrisBoard.Get(block.Coords.L())
					if left != nil && !left.Moving {
						canMove = false
						break
					}
				}
			}
			if canMove {
				for _, block := range data.TetrisBoard.Shape.Blocks {
					if block != nil {
						data.TetrisBoard.Set(block.Coords, nil)
						block.Coords.X--
					}
				}
				for _, block := range data.TetrisBoard.Shape.Blocks {
					if block != nil {
						data.TetrisBoard.Set(block.Coords, block)
					}
				}
			}
		} else if MoveRight && !PieceDone {
			canMove := true
			for _, block := range data.TetrisBoard.Shape.Blocks {
				if block != nil {
					if !BlockLegal(block.Coords.R()) {
						canMove = false
						break
					}
					right := data.TetrisBoard.Get(block.Coords.R())
					if right != nil && !right.Moving {
						canMove = false
						break
					}
				}
			}
			if canMove {
				for _, block := range data.TetrisBoard.Shape.Blocks {
					if block != nil {
						data.TetrisBoard.Set(block.Coords, nil)
						block.Coords.X++
					}
				}
				for _, block := range data.TetrisBoard.Shape.Blocks {
					if block != nil {
						data.TetrisBoard.Set(block.Coords, block)
					}
				}
			}
		}
		if !PieceDone && Rotate && !data.TetrisBoard.Shape.NoRot {
			canRot := true
			pivot := data.TetrisBoard.Shape.Blocks[0]
			if pivot != nil {
				for _, block := range data.TetrisBoard.Shape.Blocks {
					if block != nil {
						if block.Coords != pivot.Coords {
							xDiff := block.Coords.X - pivot.Coords.X
							yDiff := block.Coords.Y - pivot.Coords.Y
							try := world.Coords{X: pivot.Coords.X + yDiff, Y: pivot.Coords.Y - xDiff}
							if !BlockLegal(try) {
								canRot = false
								break
							}
							tryHere := data.TetrisBoard.Get(try)
							if tryHere != nil && !tryHere.Moving {
								canRot = false
								break
							}
						}
					}
				}
			}
			if canRot {
				for i, block := range data.TetrisBoard.Shape.Blocks {
					if i != 0 && block != nil {
						data.TetrisBoard.Set(block.Coords, nil)
						xDiff := block.Coords.X - pivot.Coords.X
						yDiff := block.Coords.Y - pivot.Coords.Y
						block.Coords = world.Coords{X: pivot.Coords.X + yDiff, Y: pivot.Coords.Y - xDiff}
					}
				}
				for i, block := range data.TetrisBoard.Shape.Blocks {
					if i != 0 && block != nil {
						data.TetrisBoard.Set(block.Coords, block)
					}
				}
			}
		}
	}
	for _, result := range myecs.Manager.Query(myecs.IsBlock) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		block, ok := result.Components[myecs.Block].(*data.TetrisBlock)
		if okO && ok {
			obj.Pos = world.MapToWorld(block.Coords)
		}
	}
	BlockUpdate = false
	MoveDown = false
	MoveLeft = false
	MoveRight = false
	Rotate = false
	if data.TetrisBoard.Timer.Done() {
		data.TetrisBoard.ResetTimer()
	}
}

// Places Tetronimo on the Board
func PlaceTetronimo() bool {
	//Validate Blocks
	for _, block := range data.TetrisBoard.NextShape.Blocks {
		if BlockLegal(block.Coords) {
			if data.TetrisBoard.Board[block.Coords.Y][block.Coords.X] == nil {
				data.TetrisBoard.Board[block.Coords.Y][block.Coords.X] = block
				obj := object.New()
				obj.Pos = world.MapToWorld(block.Coords)
				obj.Layer = 2
				spr := img.NewSprite(block.Color.String(), constants.BlockKey)
				block.Entity = myecs.Manager.NewEntity()
				block.Entity.
					AddComponent(myecs.Object, obj).
					AddComponent(myecs.Block, block).
					AddComponent(myecs.Drawable, spr)
			} else {
				return false
			}
		} else {
			return false
		}
	}
	//Put Next Tetronimo in Shape
	data.TetrisBoard.Shape = data.TetrisBoard.NextShape
	data.TetrisBoard.NextShape = NewTetronimo()
	data.TetrisBoard.ResetTimer()
	data.TetrisBoard.Stats.Tetronimos++
	return true
}

// Creates Standalone Tetronimo
func NewTetronimo() *data.Tetronimo {
	col := data.RandColor()
	t := &data.Tetronimo{}
	t.TetType = constants.TetronimoType(rand.Intn(7))
	switch t.TetType {
	case constants.O:
		t = CreateOTetronimo(col)
	case constants.I:
		t = CreateITetronimo(col)
	case constants.L:
		t = CreateLTetronimo(col)
	case constants.J:
		t = CreateJTetronimo(col)
	case constants.S:
		t = CreateSTetronimo(col)
	case constants.Z:
		t = CreateZTetronimo(col)
	case constants.T:
		t = CreateTTetronimo(col)
	}
	return t
}
func FactoTet(f *data.FacTetronimo) {
	//detect Fac Type
	var originalCoords [4]world.Coords
	for i, block := range f.Blocks {
		originalCoords[i] = world.Coords{block.Coords.X, block.Coords.Y}
	}
	newCoords := data.Normalize(originalCoords)
	if len(newCoords) == 4 {
		var tetType *constants.TetronimoType
		for i, kv := range constants.NormalizedTetronimos {
			if data.TetronimoCoordsEqual(i, newCoords) {
				tetType = &kv
				break
			}
		}
		if tetType != nil {
			switch *tetType {
			case constants.O:
				data.TetrisBoard.NextShape = CreateOTetronimo(f.Blocks[0].Color)
				return
			case constants.I:
				data.TetrisBoard.NextShape = CreateITetronimo(f.Blocks[0].Color)
				return
			case constants.L:
				data.TetrisBoard.NextShape = CreateLTetronimo(f.Blocks[0].Color)
				return
			case constants.J:
				data.TetrisBoard.NextShape = CreateJTetronimo(f.Blocks[0].Color)
				return
			case constants.S:
				data.TetrisBoard.NextShape = CreateSTetronimo(f.Blocks[0].Color)
				return
			case constants.Z:
				data.TetrisBoard.NextShape = CreateZTetronimo(f.Blocks[0].Color)
				return
			case constants.T:
				data.TetrisBoard.NextShape = CreateTTetronimo(f.Blocks[0].Color)
				return
			}
		}
	}

}

func CreateITetronimo(col data.TColor) *data.Tetronimo {
	t := &data.Tetronimo{}
	s := constants.TetrisStart
	t.TetType = constants.I
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X++
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X++
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X -= 3
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	return t
}

func CreateOTetronimo(col data.TColor) *data.Tetronimo {
	t := &data.Tetronimo{}
	s := constants.TetrisStart
	t.TetType = constants.O
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X++
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.Y--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	t.NoRot = true
	return t
}
func CreateTTetronimo(col data.TColor) *data.Tetronimo {
	t := &data.Tetronimo{}
	s := constants.TetrisStart
	t.TetType = constants.T
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X++
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X -= 2
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X++
	s.Y--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	return t
}
func CreateSTetronimo(col data.TColor) *data.Tetronimo {
	t := &data.Tetronimo{}
	s := constants.TetrisStart
	t.TetType = constants.S
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X++
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X--
	s.Y--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	return t
}
func CreateZTetronimo(col data.TColor) *data.Tetronimo {
	t := &data.Tetronimo{}
	s := constants.TetrisStart
	t.TetType = constants.Z
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X++
	s.Y--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X++
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	return t
}
func CreateJTetronimo(col data.TColor) *data.Tetronimo {
	t := &data.Tetronimo{}
	s := constants.TetrisStart
	t.TetType = constants.J
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X += 2
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.Y--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	return t
}
func CreateLTetronimo(col data.TColor) *data.Tetronimo {
	t := &data.Tetronimo{}
	s := constants.TetrisStart
	t.TetType = constants.L
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X++
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.X -= 2
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	s.Y--
	t.Blocks = append(t.Blocks, StandaloneBlock(s, col))
	return t
}
func StandaloneBlock(c world.Coords, col data.TColor) *data.TetrisBlock {
	block := &data.TetrisBlock{
		Coords: c,
		Color:  col,
		Moving: true,
	}
	return block
}

func CreateBlock(c world.Coords, col data.TColor) *data.TetrisBlock {
	if BlockLegal(c) {
		if data.TetrisBoard.Board[c.Y][c.X] == nil {
			block := &data.TetrisBlock{
				Coords: c,
				Color:  col,
				Moving: true,
			}
			data.TetrisBoard.Board[c.Y][c.X] = block
			obj := object.New()
			obj.Pos = world.MapToWorld(c)
			obj.Layer = 2
			spr := img.NewSprite(col.String(), constants.BlockKey)
			block.Entity = myecs.Manager.NewEntity()
			block.Entity.
				AddComponent(myecs.Object, obj).
				AddComponent(myecs.Block, block).
				AddComponent(myecs.Drawable, spr)
			return block
		}
	}
	return nil
}

func BlockLegal(c world.Coords) bool {
	return c.X >= 0 && c.X < constants.TetrisWidth &&
		c.Y >= 0 && c.Y < constants.TetrisHeight
}
