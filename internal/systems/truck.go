package systems

import (
	"github.com/faiface/pixel"
	"math/rand"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
)

func GenerateLoad(t *data.Truck) {
	t.MyBatchType = *constants.DeliveryBatchRoll()
	for _, size := range constants.DeliveryLoadBatch[t.MyBatchType] {
		switch size {
		case 3:
			t.DeliveryLoad = append(t.DeliveryLoad, CreateFactoryTet(pixel.ZV, data.RandColor(), constants.FactrominoType(constants.FacThree)))
		case 2:
			t.DeliveryLoad = append(t.DeliveryLoad, CreateFactoryTet(pixel.ZV, data.RandColor(), constants.FactrominoType(constants.FacTwo)))
		case 1:
			t.DeliveryLoad = append(t.DeliveryLoad, CreateFactoryTet(pixel.ZV, data.RandColor(), constants.FactrominoType(constants.FacOne)))
		}
	}
	//pixel.zv
}

func RandomizeLoad(t *data.Truck) {
	if t != nil && len(t.DeliveryLoad) > 1 {
		for i := len(t.DeliveryLoad) - 1; i > 0; i-- {
			j := rand.Intn(i)
			t.DeliveryLoad[i], t.DeliveryLoad[j] = t.DeliveryLoad[j], t.DeliveryLoad[i]
		}
	}
}
