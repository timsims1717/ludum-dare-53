package data

import (
	"timsims1717/ludum-dare-53/internal/constants"
)

type Truck struct {
	DeliveryLoad []*FacTetronimo
	MyBatchType  constants.DeliveryLoadBatchTypes
}

func DeliveryLoadBatchAssignment(randI int) *constants.DeliveryLoadBatchTypes {
	BatchWeights := map[constants.DeliveryLoadBatchTypes][2]int{}
	previous := 0
	for i := 0; i < constants.NumberofBatchTypes; i++ {
		BatchWeights[constants.DeliveryLoadBatchTypes(i)] = [2]int{previous + 1, previous + constants.DeliveryLoadBatchWeights[constants.DeliveryLoadBatchTypes(i)]}
		previous = BatchWeights[constants.DeliveryLoadBatchTypes(i)][1]
	}
	for i, batchrange := range BatchWeights {
		if randI >= batchrange[0] && randI <= batchrange[1] {
			return &i
		}
	}
	return nil
}
