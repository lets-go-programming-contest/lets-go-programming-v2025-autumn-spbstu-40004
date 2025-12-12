package maxheap

type MaxHeap []int

func (maxHeap *MaxHeap) Len() int {
	return len(*maxHeap)
}

func (maxHeap *MaxHeap) Less(indexFirst, indexSecond int) bool {
	return (*maxHeap)[indexFirst] >= (*maxHeap)[indexSecond]
}

func (maxHeap *MaxHeap) Swap(indexFirst, indexSecond int) {
	(*maxHeap)[indexFirst], (*maxHeap)[indexSecond] = (*maxHeap)[indexSecond], (*maxHeap)[indexFirst]
}

func (maxHeap *MaxHeap) Push(value any) {
	val, ok := value.(int)
	if !ok {
		panic("Value must be int in MaxHeap")
	}

	*maxHeap = append(*maxHeap, val)
}

func (maxHeap *MaxHeap) Pop() any {
	oldLen := len(*maxHeap)
	if oldLen == 0 {
		return nil
	}

	oldHeap := *maxHeap
	poppedValue := oldHeap[oldLen-1]
	*maxHeap = oldHeap[0 : oldLen-1]

	return poppedValue
}
