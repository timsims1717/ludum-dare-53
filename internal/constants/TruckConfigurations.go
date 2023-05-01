package constants

const (
	NumberofBatchTypes = 4
)

var (
	DeliveryLoadBatchWeights = map[DeliveryLoadBatchTypes]int{
		BatchA: 25,
		BatchB: 25,
		BatchC: 25,
		BatchD: 25,
	}
	DeliveryLoadBatch = map[DeliveryLoadBatchTypes][]int{
		BatchA: {3, 2, 1, 1},
		BatchB: {3, 1, 1, 1, 1},
		BatchC: {2, 2, 1, 1, 1},
		BatchD: {2, 2, 2, 1},
	}
	RawFactrominoWeights = map[FactrominoType]int{
		FacOne:   34,
		FacTwo:   33,
		FacThree: 33,
	}
	FactrominoThreeVariationWeights = map[FactrominoVariant]int{
		Vertical:   18,
		Horizontal: 18,
		BabyR:      16,
		BabySeven:  16,
		BabyL:      16,
		BabyJ:      16,
	}
)

func FactrominoThreeVariationRoll() FactrominoVariant {
	randomLoadInt := GlobalSeededRandom.Intn(100) + 1
	FactrominoVariantWeights := map[FactrominoVariant][2]int{}
	previous := 0
	for i := 1; i < 7; i++ {
		FactrominoVariantWeights[FactrominoVariant(i)] = [2]int{previous + 1, previous + FactrominoThreeVariationWeights[FactrominoVariant(i)]}
		previous = FactrominoVariantWeights[FactrominoVariant(i)][1]
	}
	for i, batchrange := range FactrominoVariantWeights {
		if randomLoadInt >= batchrange[0] && randomLoadInt <= batchrange[1] {
			return i
		}
	}
	return FactVariantUndefined
}

func RawFactrominoRoll() FactrominoType {
	randomLoadInt := GlobalSeededRandom.Intn(100) + 1
	FactrominoWeights := map[FactrominoType][2]int{}
	previous := 0
	for i := 1; i < 4; i++ {
		FactrominoWeights[FactrominoType(i)] = [2]int{previous + 1, previous + RawFactrominoWeights[FactrominoType(i)]}
		previous = FactrominoWeights[FactrominoType(i)][1]
	}
	for i, batchrange := range FactrominoWeights {
		if randomLoadInt >= batchrange[0] && randomLoadInt <= batchrange[1] {
			return i
		}
	}
	return FacUndefined
}

type DeliveryLoadBatchTypes int

const (
	BatchA = iota //3,2,1,1
	BatchB        //3,1,1,1
	BatchC        //2,2,1,1,1
	BatchD        //2,2,2,1,

)

func DeliveryBatchRoll() *DeliveryLoadBatchTypes {
	randomLoadInt := GlobalSeededRandom.Intn(100) + 1
	BatchWeights := map[DeliveryLoadBatchTypes][2]int{}
	previous := 0
	for i := 0; i < NumberofBatchTypes; i++ {
		BatchWeights[DeliveryLoadBatchTypes(i)] = [2]int{previous + 1, previous + DeliveryLoadBatchWeights[DeliveryLoadBatchTypes(i)]}
		previous = BatchWeights[DeliveryLoadBatchTypes(i)][1]
	}
	for i, batchrange := range BatchWeights {
		if randomLoadInt >= batchrange[0] && randomLoadInt <= batchrange[1] {
			return &i
		}
	}
	return nil
}

func (b DeliveryLoadBatchTypes) String() string {
	switch b {
	case BatchA:
		return "Batch A: {3,2,1,1}"
	case BatchB:
		return "Batch B: {3,1,1,1,1}"
	case BatchC:
		return "Batch C: {2,2,1,1,1}"
	case BatchD:
		return "Batch D: {2,2,2,1}"
	}
	return ""
}
