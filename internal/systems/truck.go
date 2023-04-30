package systems

import (
	"github.com/faiface/pixel"
	"math/rand"
	"time"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/data"
)

func GenerateLoad(t *data.Truck) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomLoadInt := r.Intn(100) + 1
	t.MyBatchType = *data.DeliveryLoadBatchAssignment(randomLoadInt)
	for i, size := range constants.DeliveryLoadBatch[t.MyBatchType] {
		switch size {
		case 3:
		case 2:
		case 1:
			t.DeliveryLoad = append(t.DeliveryLoad, CreateFactoryTet(pixel.ZV, data.RandColor(), i))
		}
	}
	//pixel.zv
}
