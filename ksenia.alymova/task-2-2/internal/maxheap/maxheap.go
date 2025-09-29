package maxheap

type MaxHeap []int

func (maxHeap *MaxHeap) Len() int {
	return len(*maxHeap)
}

func (maxHeap *MaxHeap) Less(indexLhs, indexRhs int) bool {
	return (*maxHeap)[indexLhs] >= (*maxHeap)[indexRhs]
}

func (maxHeap *MaxHeap) Swap(indexLhs, indexRhs int) {
	(*maxHeap)[indexLhs], (*maxHeap)[indexRhs] = (*maxHeap)[indexRhs], (*maxHeap)[indexLhs]
}

func (maxHeap *MaxHeap) Push(value any) {
	intValue, ok := value.(int)
	if ok {
		*maxHeap = append(*maxHeap, intValue)
	}
}

func (maxHeap *MaxHeap) Pop() any {
	oldLen := len(*maxHeap)
	returnValue := (*maxHeap)[oldLen-1]
	*maxHeap = (*maxHeap)[:oldLen-1]

	return returnValue
}
