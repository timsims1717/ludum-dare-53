package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"math/rand"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/world"
)

func CreateFactoryTet(pos pixel.Vec, col data.TColor, size int) *data.FacTetronimo {
	t := &data.FacTetronimo{}
	t.LastPos = pos
	t.Object = object.New().WithID("factory-tet")
	t.Object.Pos = pos
	t.Object.Layer = 12
	w := constants.FactoryTile
	h := world.TileSize + 6.
	if size == 0 {
		size = rand.Intn(18)
	}
	switch size {
	case 0, 1, 2, 3, 4, 5: //1
		a := CreateFactoryBlock(pixel.ZV, col)
		a.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
	case 6, 7, 8: //2 Horizontal
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.X -= constants.FactoryTile * 0.5
		b.Object.Offset.X += constants.FactoryTile * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		w += constants.FactoryTile
	case 9, 10, 11: //2 Vertical
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.Y += world.TileSize * 0.5
		b.Object.Offset.Y -= world.TileSize * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		h += world.TileSize
	case 12: //3 Horizontal
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		c := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.X += constants.FactoryTile
		b.Object.Offset.X -= constants.FactoryTile
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		c.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		t.Blocks = append(t.Blocks, c)
		w += constants.FactoryTile * 2.
	case 13: //3 Vertical
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		c := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.Y += world.TileSize
		c.Object.Offset.Y -= world.TileSize
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		c.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		t.Blocks = append(t.Blocks, c)
		h += world.TileSize * 2.
	case 14: //3 r
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		c := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.Y += world.TileSize * 0.5
		a.Object.Offset.X -= constants.FactoryTile * 0.5
		b.Object.Offset.Y += world.TileSize * 0.5
		b.Object.Offset.X += constants.FactoryTile * 0.5
		c.Object.Offset.Y -= world.TileSize * 0.5
		c.Object.Offset.X -= constants.FactoryTile * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		c.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		t.Blocks = append(t.Blocks, c)
		w += constants.FactoryTile
		h += world.TileSize
	case 15: //3 7
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		c := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.Y += world.TileSize * 0.5
		a.Object.Offset.X -= constants.FactoryTile * 0.5
		b.Object.Offset.Y += world.TileSize * 0.5
		b.Object.Offset.X += constants.FactoryTile * 0.5
		c.Object.Offset.Y -= world.TileSize * 0.5
		c.Object.Offset.X += constants.FactoryTile * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		c.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		t.Blocks = append(t.Blocks, c)
		w += constants.FactoryTile
		h += world.TileSize
	case 16: //3 L
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		c := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.Y += world.TileSize * 0.5
		a.Object.Offset.X -= constants.FactoryTile * 0.5
		b.Object.Offset.Y -= world.TileSize * 0.5
		b.Object.Offset.X += constants.FactoryTile * 0.5
		c.Object.Offset.Y -= world.TileSize * 0.5
		c.Object.Offset.X -= constants.FactoryTile * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		c.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		t.Blocks = append(t.Blocks, c)
		w += constants.FactoryTile
		h += world.TileSize
	case 17: //3 J
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		c := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.Y += world.TileSize * 0.5
		a.Object.Offset.X += constants.FactoryTile * 0.5
		b.Object.Offset.Y -= world.TileSize * 0.5
		b.Object.Offset.X += constants.FactoryTile * 0.5
		c.Object.Offset.Y -= world.TileSize * 0.5
		c.Object.Offset.X -= constants.FactoryTile * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		c.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		t.Blocks = append(t.Blocks, c)
		w += constants.FactoryTile
		h += world.TileSize
	}
	t.Object.Rect = pixel.R(0., 0., w, h)
	t.Entity = myecs.Manager.NewEntity()
	t.Entity.AddComponent(myecs.Object, t.Object).
		AddComponent(myecs.Block, t)
	return t
}

func ConstructTetFromBlocks(pos pixel.Vec, blocks []*data.FactoryBlock) *data.FacTetronimo {
	ft := &data.FacTetronimo{}
	ft.LastPos = pos
	ft.Object = object.New().WithID("factory-tet")
	ft.Object.Pos = pos
	ft.Object.Layer = 20
	l := -1.
	r := -1.
	b := -1.
	t := -1.
	for _, block := range blocks {
		p := world.MapToWorldC(block.Coords, pixel.V(constants.FactoryTile, world.TileSize))
		if l == -1 || p.X < l {
			l = p.X
		}
		if r == -1 || p.X > r {
			r = p.X
		}
		if t == -1 || p.Y > t {
			t = p.Y
		}
		if b == -1 || p.Y < b {
			b = p.Y
		}
	}
	size := pixel.V(r-l, t-b)
	center := pixel.V(l+size.X*0.5, b+size.Y*0.5)
	for _, block := range blocks {
		block.Object.Offset = block.Object.Pos.Sub(center)
		block.Object.Pos = pixel.ZV
		block.Object.Layer = 20
		block.Entity.AddComponent(myecs.Parent, ft.Object)
		ft.Blocks = append(ft.Blocks, block)
	}
	ft.Object.Rect = pixel.R(0., 0., r-l+constants.FactoryTile, t-b+world.TileSize+6.)
	ft.Entity = myecs.Manager.NewEntity()
	ft.Entity.AddComponent(myecs.Object, ft.Object).
		AddComponent(myecs.Block, ft)
	return ft
}

func CreateFactoryBlock(pos pixel.Vec, col data.TColor) *data.FactoryBlock {
	block := &data.FactoryBlock{
		Coords: world.Origin,
		Color:  col,
	}
	block.Object = object.New().WithID("factory-block")
	block.Object.Pos = pos
	block.Object.Layer = 12
	block.Object.Rect = pixel.R(0., 0., constants.FactoryTile, world.TileSize)
	spr := img.NewSprite(fmt.Sprintf("%s_f", col.String()), constants.BlockKey)
	block.Entity = myecs.Manager.NewEntity()
	block.Entity.
		AddComponent(myecs.Object, block.Object).
		AddComponent(myecs.Drawable, spr)
	return block
}