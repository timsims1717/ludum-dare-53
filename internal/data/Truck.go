package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/pkg/object"
)

type Truck struct {
	DeliveryLoad []*Factromino
	MyBatchType  constants.DeliveryLoadBatchTypes

	Object *object.Object
	Entity *ecs.Entity
	Start  pixel.Vec
	End    pixel.Vec
	Pad    *FactoryPad
}
