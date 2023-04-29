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

func CreateFactoryTet(pos pixel.Vec, col data.TColor) *data.FacTetronimo {
	t := &data.FacTetronimo{}
	t.LastPos = pos
	t.Object = object.New().WithID("factory-tet")
	t.Object.Pos = pos
	t.Object.Layer = 12
	w := constants.FactoryTile
	h := world.TileSize + 6.
	switch rand.Intn(9) {
	case 0:
		a := CreateFactoryBlock(pixel.ZV, col)
		a.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
	case 1:
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.X -= constants.FactoryTile * 0.5
		b.Object.Offset.X += constants.FactoryTile * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		w += constants.FactoryTile
	case 2:
		a := CreateFactoryBlock(pixel.ZV, col)
		b := CreateFactoryBlock(pixel.ZV, col)
		a.Object.Offset.Y += world.TileSize * 0.5
		b.Object.Offset.Y -= world.TileSize * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		h += world.TileSize
	case 3:
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
	case 4:
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
	case 5:
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
	case 6:
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
	case 7:
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
	case 8:
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
