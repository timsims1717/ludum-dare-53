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
)

type DeliveryLoadBatchTypes int

const (
	BatchA = iota //3,2,1,1
	BatchB        //3,1,1,1
	BatchC        //2,2,1,1,1
	BatchD        //2,2,2,1,

)

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
