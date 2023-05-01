package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/world"
)

func CreateFactoryTet(pos pixel.Vec, col data.TColor, factrominoType constants.FactrominoType) *data.FacTetronimo {
	t := &data.FacTetronimo{MyFactronimoType: factrominoType}
	t.LastPos = pos
	t.Object = object.New().WithID("factory-tet")
	t.Object.Pos = pos
	t.Object.Layer = 12
	t.Color = col
	w := constants.FactoryTile
	h := world.TileSize + 6.
	if t.MyFactronimoType == constants.FacUndefined {
		t.MyFactronimoType = constants.RawFactrominoRoll()
	}
	switch t.MyFactronimoType {
	case constants.FacOne: //1
		CreateFactrominoSizeOne(t)
	case constants.FacTwo:
		CreateFactrominoSizeTwo(t, w, h)
	case constants.FacThree:
		CreateFactrominoSizeThree(t, w, h)
	}
	t.Object.Rect = pixel.R(0., 0., w, h)
	t.Entity = myecs.Manager.NewEntity()
	t.Entity.AddComponent(myecs.Object, t.Object).
		AddComponent(myecs.Block, t)
	return t
}
func CreateFactrominoSizeOne(t *data.FacTetronimo) {
	a := CreateFactoryBlock(pixel.ZV, t.Color)
	a.Entity.AddComponent(myecs.Parent, t.Object)
	t.Blocks = append(t.Blocks, a)
}
func CreateFactrominoSizeTwo(t *data.FacTetronimo, w float64, h float64) {
	t.MyFactronimoVariant = constants.FactVariantUndefined
	t.MyFactronimoVariant = constants.FactrominoVariant(constants.GlobalSeededRandom.Intn(2) + 1)

	switch t.MyFactronimoVariant {
	case constants.Horizontal: //2 Horizontal
		a := CreateFactoryBlock(pixel.ZV, t.Color)
		b := CreateFactoryBlock(pixel.ZV, t.Color)
		a.Object.Offset.X -= constants.FactoryTile * 0.5
		b.Object.Offset.X += constants.FactoryTile * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		w += constants.FactoryTile
	case constants.Vertical: //2 Vertical
		a := CreateFactoryBlock(pixel.ZV, t.Color)
		b := CreateFactoryBlock(pixel.ZV, t.Color)
		a.Object.Offset.Y += world.TileSize * 0.5
		b.Object.Offset.Y -= world.TileSize * 0.5
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		h += world.TileSize
	}
}
func CreateFactrominoSizeThree(t *data.FacTetronimo, w float64, h float64) {
	t.MyFactronimoVariant = constants.FactVariantUndefined
	t.MyFactronimoVariant = constants.FactrominoThreeVariationRoll()
	switch t.MyFactronimoVariant {
	case constants.Horizontal:
		a := CreateFactoryBlock(pixel.ZV, t.Color)
		b := CreateFactoryBlock(pixel.ZV, t.Color)
		c := CreateFactoryBlock(pixel.ZV, t.Color)
		a.Object.Offset.X += constants.FactoryTile
		b.Object.Offset.X -= constants.FactoryTile
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		c.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		t.Blocks = append(t.Blocks, c)
		w += constants.FactoryTile * 2.
	case constants.Vertical:
		a := CreateFactoryBlock(pixel.ZV, t.Color)
		b := CreateFactoryBlock(pixel.ZV, t.Color)
		c := CreateFactoryBlock(pixel.ZV, t.Color)
		a.Object.Offset.Y += world.TileSize
		c.Object.Offset.Y -= world.TileSize
		a.Entity.AddComponent(myecs.Parent, t.Object)
		b.Entity.AddComponent(myecs.Parent, t.Object)
		c.Entity.AddComponent(myecs.Parent, t.Object)
		t.Blocks = append(t.Blocks, a)
		t.Blocks = append(t.Blocks, b)
		t.Blocks = append(t.Blocks, c)
		h += world.TileSize * 2.
	case constants.BabyR:
		a := CreateFactoryBlock(pixel.ZV, t.Color)
		b := CreateFactoryBlock(pixel.ZV, t.Color)
		c := CreateFactoryBlock(pixel.ZV, t.Color)
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
	case constants.BabySeven:
		a := CreateFactoryBlock(pixel.ZV, t.Color)
		b := CreateFactoryBlock(pixel.ZV, t.Color)
		c := CreateFactoryBlock(pixel.ZV, t.Color)
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
	case constants.BabyL:
		a := CreateFactoryBlock(pixel.ZV, t.Color)
		b := CreateFactoryBlock(pixel.ZV, t.Color)
		c := CreateFactoryBlock(pixel.ZV, t.Color)
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
	case constants.BabyJ:
		a := CreateFactoryBlock(pixel.ZV, t.Color)
		b := CreateFactoryBlock(pixel.ZV, t.Color)
		c := CreateFactoryBlock(pixel.ZV, t.Color)
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
