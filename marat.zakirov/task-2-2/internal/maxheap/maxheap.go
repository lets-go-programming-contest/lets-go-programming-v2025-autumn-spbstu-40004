package maxheap

type MaxHeap []int

func (maxHeap *MaxHeap) Len() int {
	return len(*maxHeap)
}

func (maxHeap *MaxHeap) Less(indexSt, indexNd int) bool {
	return (*maxHeap)[indexSt] >= (*maxHeap)[indexNd]
}

func (maxHeap *MaxHeap) Swap(indexSt, indexNd int) {
	(*maxHeap)[indexSt], (*maxHeap)[indexNd] = (*maxHeap)[indexNd], (*maxHeap)[indexSt]
}

func (maxHeap *MaxHeap) Push(value any) {
	assertedValue, ok := value.(int)
	if ok {
		*maxHeap = append(*maxHeap, assertedValue)
	} else {
		panic("Some errors in append")
	}
}

func (maxHeap *MaxHeap) Pop() any {
	oldLen := len(*maxHeap)
	if oldLen == 0 {
		return nil
	}

	oldHeap := *maxHeap
	lastValue := oldHeap[oldLen-1]
	*maxHeap = oldHeap[0 : oldLen-1]

	return lastValue
}
