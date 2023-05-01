package state

import (
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/object"
)

var wallHeight = data.MSize * 26.

func LoadTileMaps() {
	// floor tile
	for y := 0; y < 10; y++ {
		for x := 0; x < 18; x++ {
			obj := object.New()
			obj.Pos.Y = float64(-4+y)*data.MSize*7. - data.MSize
			obj.Pos.X = float64(-9+x) * data.MSize * 11.
			obj.Layer = 9
			myecs.Manager.NewEntity().
				AddComponent(myecs.Object, obj).
				AddComponent(myecs.Drawable, data.FloorSection)
		}
	}
	// wall section
	for x := 0; x < 100; x++ {
		obj := object.New()
		obj.Pos.Y = wallHeight
		obj.Pos.X = float64(-50+x) * data.MSize * 2.
		if obj.Pos.X > 0. {
			continue
		}
		obj.Layer = 12
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, obj).
			AddComponent(myecs.Drawable, data.WallSection)
	}
	// big back doors
	door1Obj := object.New()
	door1Obj.Pos.Y = wallHeight
	door1Obj.Pos.X = 2. * data.MSize
	door1Obj.Layer = 12
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, door1Obj).
		AddComponent(myecs.Drawable, data.DoorSection)
	door2Obj := object.New()
	door2Obj.Pos.Y = wallHeight
	door2Obj.Pos.X = 18. * data.MSize
	door2Obj.Layer = 12
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, door2Obj).
		AddComponent(myecs.Drawable, data.DoorSection)
	tinyWallObj := object.New()
	tinyWallObj.Pos.Y = wallHeight
	tinyWallObj.Pos.X = 34. * data.MSize
	tinyWallObj.Layer = 12
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, tinyWallObj).
		AddComponent(myecs.Drawable, data.WallSection)
	// side wall
	for y := 0; y < 66; y++ {
		obj := object.New()
		obj.Pos.X = 36. * data.MSize
		obj.Pos.Y = float64(-26+y) * data.MSize
		obj.Layer = 12
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, obj).
			AddComponent(myecs.Drawable, data.SideSection)
	}
	cornerObj := object.New()
	cornerObj.Pos.X = 36. * data.MSize
	cornerObj.Pos.Y = float64(40) * data.MSize
	cornerObj.Layer = 12
	// side doors
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, cornerObj).
		AddComponent(myecs.Drawable, data.CornerSection)
	for y := 0; y < 2; y++ {
		obj := object.New()
		obj.Pos.X = 35. * data.MSize
		if y == 0 {
			obj.Pos.Y = data.MSize * -11.
		} else {
			obj.Pos.Y = data.MSize * 3.
		}
		obj.Layer = 12
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, obj).
			AddComponent(myecs.Drawable, data.SideDSection)
	}
	// conveyor belt
	objConvBase := object.New()
	objConvBase.Pos.X = -2. * data.MSize
	objConvBase.Pos.Y = 17. * data.MSize
	objConvBase.Layer = 10
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, objConvBase).
		AddComponent(myecs.Drawable, data.ConveyorBase)
	// conveyor sections
	for x := 0; x < data.BeltSize; x++ {
		obj := object.New()
		obj.Pos.X = (float64(x-data.BeltSize) - 2.) * data.MSize
		obj.Pos.Y = 18. * data.MSize
		obj.Layer = 11
		e := myecs.Manager.NewEntity().AddComponent(myecs.Object, obj)
		if x == 0 {
			e.AddComponent(myecs.Drawable, data.ConvLeftEdge)
		} else if x == data.BeltSize-1 {
			e.AddComponent(myecs.Drawable, data.ConvRightEdge)
		} else {
			e.AddComponent(myecs.Drawable, data.ConvMiddle)
		}
	}
	// TV
	objTV := object.New()
	objTV.Pos.X = -22. * data.MSize
	objTV.Pos.Y = -2. * data.MSize
	objTV.Layer = 10
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, objTV).
		AddComponent(myecs.Drawable, data.TV)
	// bulb
	objBulb := object.New()
	objBulb.Pos.X = -46. * data.MSize
	objBulb.Pos.Y = -16. * data.MSize
	objBulb.Layer = 13
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, objBulb).
		AddComponent(myecs.Drawable, data.Bulb)
}
