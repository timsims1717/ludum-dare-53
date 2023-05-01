package state

import (
	"github.com/faiface/pixel"
	"math/rand"
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/internal/systems"
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
	botTruck := &data.Truck{}
	botTruckObj := object.New().WithID("south_truck")
	botTruckObj.Pos.X = PadX1
	botTruckObj.Pos.Y = botTruckYS
	botTruckObj.Layer = 11
	start := botTruckObj.Pos
	end := botTruckObj.Pos
	end.Y = botTruckYE
	botTruck.Object = botTruckObj
	botTruck.Start = start
	botTruck.End = end
	botTruck.Pad = data.SouthPad
	botTruck.Entity = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, botTruckObj).
		AddComponent(myecs.Drawable, data.BotTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(botTruck, false)))
	// southeast truck
	seTruck := &data.Truck{}
	seTruckObj := object.New().WithID("southeast_truck")
	seTruckObj.Pos.X = midTruckXS
	seTruckObj.Pos.Y = PadY3 + 2.*data.MSize
	seTruckObj.Layer = 11
	start = seTruckObj.Pos
	seTruckObj.Pos.X -= 10. * data.MSize
	end = seTruckObj.Pos
	end.X = midTruckXE
	seTruck.Object = seTruckObj
	seTruck.Start = start
	seTruck.End = end
	seTruck.Pad = data.SouthEastPad
	seTruck.Entity = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, seTruckObj).
		AddComponent(myecs.Drawable, data.MidTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(seTruck, true)))
	// east truck
	eTruck := &data.Truck{}
	eTruckObj := object.New().WithID("east_truck")
	eTruckObj.Pos.X = midTruckXS
	eTruckObj.Pos.Y = PadY2 + 2.*data.MSize
	eTruckObj.Layer = 11
	start = eTruckObj.Pos
	end = eTruckObj.Pos
	end.X = midTruckXE
	eTruck.Object = eTruckObj
	eTruck.Start = start
	eTruck.End = end
	eTruck.Pad = data.EastPad
	eTruck.Entity = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, eTruckObj).
		AddComponent(myecs.Drawable, data.MidTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(eTruck, false)))
	// northeast truck
	neTruck := &data.Truck{}
	neTruckObj := object.New().WithID("northeast_truck")
	neTruckObj.Pos.X = PadX2
	neTruckObj.Pos.Y = topTruckYS
	neTruckObj.Layer = 11
	start = neTruckObj.Pos
	neTruckObj.Pos.Y -= 20. * data.MSize
	end = neTruckObj.Pos
	end.Y = topTruckYE
	neTruck.Object = neTruckObj
	neTruck.Start = start
	neTruck.End = end
	neTruck.Pad = data.NorthEastPad
	neTruck.Entity = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, neTruckObj).
		AddComponent(myecs.Drawable, data.TopTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(neTruck, true)))
	// north truck
	nTruck := &data.Truck{}
	nTruckObj := object.New().WithID("north_truck")
	nTruckObj.Pos.X = PadX1
	nTruckObj.Pos.Y = topTruckYS
	nTruckObj.Layer = 11
	start = nTruckObj.Pos
	end = nTruckObj.Pos
	end.Y = topTruckYE
	nTruck.Object = nTruckObj
	nTruck.Start = start
	nTruck.End = end
	nTruck.Pad = data.NorthPad
	nTruck.Entity = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, nTruckObj).
		AddComponent(myecs.Drawable, data.TopTruck).
		AddComponent(myecs.Update, data.NewFn(TruckUpdate(nTruck, false)))
}

func TruckUpdate(truck *data.Truck, startMove bool) func() {
	delivery := startMove
	moving := startMove
	if startMove {
		systems.GenerateLoad(truck)
		systems.RandomizeLoad(truck)
	}
	timer := timing.New(rand.Float64() * 4.)
	return func() {
		timer.Update()
		if !moving && !delivery && truck.Pad.Tet == nil && timer.Done() {
			// generate load
			systems.GenerateLoad(truck)
			systems.RandomizeLoad(truck)
			// start moving
			moving = true
			// switch to delivery
			delivery = true
		} else if moving && delivery {
			// move and check for arrived
			var arrived bool
			truck.Object.Pos, arrived = MoveToward(truck.Start, truck.End, truck.Object.Pos)
			if arrived {
				moving = false
			}
		} else if !moving && delivery && timer.Done() {
			if truck.Pad.Tet == nil {
				// unload a block
				if len(truck.DeliveryLoad) > 0 {
					block := truck.DeliveryLoad[0]
					block.Entity.AddComponent(myecs.ViewPort, data.FactoryViewport)
					block.Entity.AddComponent(myecs.Input, gameInput)
					block.Object.Pos = truck.Pad.Object.Pos
					block.Object.Layer = 10
					block.Object.Hide = false
					truck.Pad.Tet = block
					if len(truck.DeliveryLoad) == 1 {
						truck.DeliveryLoad = []*data.Factromino{}
					} else {
						truck.DeliveryLoad = truck.DeliveryLoad[1:]
					}
				}
				// if last unloaded, move and switch to !delivery
				if len(truck.DeliveryLoad) == 0 {
					moving = true
					delivery = false
					timer = timing.New(rand.Float64()*4. + 2.)
				}
			} else {
				timer = timing.New(rand.Float64()*2. + 1.)
			}
		} else if moving && !delivery {
			// move and check for arrived
			var arrived bool
			truck.Object.Pos, arrived = MoveToward(truck.End, truck.Start, truck.Object.Pos)
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
