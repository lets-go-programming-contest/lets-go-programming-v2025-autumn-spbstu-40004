package intheap

type IntHeap []int

func (intHeap IntHeap) Len() int {
	return len(intHeap)
}

func (intHeap IntHeap) Less(index1, index2 int) bool {
	return intHeap[index1] < intHeap[index2]
}

func (intHeap IntHeap) Swap(index1, index2 int) {
	intHeap[index1], intHeap[index2] = intHeap[index2], intHeap[index1]
}

func (intHeap *IntHeap) Push(value any) {
	*intHeap = append(*intHeap, value.(int))
}

func (intHeap *IntHeap) Pop() any {
	oldLen := len(*intHeap)
	if oldLen == 0 {
		return nil
	}

	oldHeap := *intHeap
	lastValue := oldHeap[oldLen-1]
	*intHeap = oldHeap[0 : oldLen-1]

	return lastValue
}
