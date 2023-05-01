package data

import (
	"timsims1717/ludum-dare-53/internal/constants"
)

type Truck struct {
	DeliveryLoad []*Factromino
	MyBatchType  constants.DeliveryLoadBatchTypes
}
