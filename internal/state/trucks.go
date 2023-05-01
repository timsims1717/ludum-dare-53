package state

import (
	"github.com/faiface/pixel"
	"math/rand"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/timing"
	"timsims1717/ludum-dare-53/pkg/util"
)

var (
	botTruckYS = PadY3 - 25.*data.MSize
	botTruckYE = PadY3 - 6.*data.MSize
	midTruckXS = PadX2 + 25.*data.MSize
	midTruckXE = PadX2 + 6.*data.MSize
	topTruckYS = PadY1 + 25.*data.MSize
	topTruckYE = PadY1 + 4.*data.MSize
)

func CreateTrucks() {
	// bottom truck
	botTruckObj := object.New()
	botTruckObj.Pos.X = PadX1
	botTruckObj.Pos.Y = botTruckYS
	botTruckObj.Layer = 11
	start := botTruckObj.Pos
	end := botTruckObj.Pos
	end.Y = botTruckYE
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, botTruckObj).
		AddComponent(myecs.Drawable, data.BotTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(start, end, botTruckObj, data.SouthPad, false)))
	// southeast truck
	seTruckObj := object.New()
	seTruckObj.Pos.X = midTruckXS
	seTruckObj.Pos.Y = PadY3 + 2.*data.MSize
	seTruckObj.Layer = 11
	start = seTruckObj.Pos
	seTruckObj.Pos.X -= 10. * data.MSize
	end = seTruckObj.Pos
	end.X = midTruckXE
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, seTruckObj).
		AddComponent(myecs.Drawable, data.MidTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(start, end, seTruckObj, data.SouthEastPad, true)))
	// east truck
	eTruckObj := object.New()
	eTruckObj.Pos.X = midTruckXS
	eTruckObj.Pos.Y = PadY2 + 2.*data.MSize
	eTruckObj.Layer = 11
	start = eTruckObj.Pos
	end = eTruckObj.Pos
	end.X = midTruckXE
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, eTruckObj).
		AddComponent(myecs.Drawable, data.MidTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(start, end, eTruckObj, data.EastPad, false)))
	// northeast truck
	neTruckObj := object.New()
	neTruckObj.Pos.X = PadX2
	neTruckObj.Pos.Y = topTruckYS
	neTruckObj.Layer = 11
	start = neTruckObj.Pos
	neTruckObj.Pos.Y -= 20. * data.MSize
	end = neTruckObj.Pos
	end.Y = topTruckYE
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, neTruckObj).
		AddComponent(myecs.Drawable, data.TopTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(start, end, neTruckObj, data.NorthEastPad, true)))
	// north truck
	nTruckObj := object.New()
	nTruckObj.Pos.X = PadX1
	nTruckObj.Pos.Y = topTruckYS
	nTruckObj.Layer = 11
	start = nTruckObj.Pos
	end = nTruckObj.Pos
	end.Y = topTruckYE
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, nTruckObj).
		AddComponent(myecs.Drawable, data.TopTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(start, end, nTruckObj, data.NorthPad, false)))
}

func TruckUpdate(start, end pixel.Vec, obj *object.Object, pad *data.FactoryPad, startMove bool) func() {
	delivery := startMove
	moving := startMove
	timer := timing.New(rand.Float64() * 4.)
	return func() {
		timer.Update()
		if !moving && !delivery && pad.Tet == nil && timer.Done() {
			// generate load

			// start moving
			moving = true
			// switch to delivery
			delivery = true
		} else if moving && delivery {
			// move and check for arrived
			var arrived bool
			obj.Pos, arrived = MoveToward(start, end, obj.Pos)
			if arrived {
				moving = false
			}
		} else if delivery && pad.Tet == nil {
			// unload a block
			// if last unloaded, move and switch to !delivery
			moving = true
			delivery = false
			timer = timing.New(rand.Float64()*4. + 2.)
		} else if moving && !delivery {
			// move and check for arrived
			var arrived bool
			obj.Pos, arrived = MoveToward(end, start, obj.Pos)
			if arrived {
				moving = false
			}
		}
	}
}

func MoveToward(start, end, curr pixel.Vec) (pixel.Vec, bool) {
	move := util.Normalize(end.Sub(curr)).Scaled(85.)
	curr.X += move.X * timing.DT
	curr.Y += move.Y * timing.DT
	arrived := false
	if end.X-start.X > 0 && curr.X > end.X {
		curr.X = end.X
		arrived = true
	} else if end.X-start.X < 0 && curr.X < end.X {
		curr.X = end.X
		arrived = true
	}
	if end.Y-start.Y > 0 && curr.Y > end.Y {
		curr.Y = end.Y
		arrived = true
	} else if end.Y-start.Y < 0 && curr.Y < end.Y {
		curr.Y = end.Y
		arrived = true
	}
	return curr, arrived
}
