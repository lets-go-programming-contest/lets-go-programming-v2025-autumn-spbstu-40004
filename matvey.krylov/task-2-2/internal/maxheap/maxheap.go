package maxheap

type MaxHeap []int

func (heap *MaxHeap) Len() int {
	return len(*heap)
}

func (heap *MaxHeap) Less(left, right int) bool {
	return (*heap)[left] > (*heap)[right]
}

func (heap *MaxHeap) Swap(left, right int) {
	(*heap)[left], (*heap)[right] = (*heap)[right], (*heap)[left]
}

func (heap *MaxHeap) Push(value any) {
	assertedValue, ok := value.(int)
	if ok {
		*heap = append(*heap, assertedValue)
	} else {
		panic("Incorrect value")
	}
}

func (heap *MaxHeap) Pop() any {
	heapLen := len(*heap)
	if heapLen == 0 {
		return nil
	}

	oldHeap := *heap
	lastValue := oldHeap[heapLen-1]
	*heap = oldHeap[0 : heapLen-1]

	return lastValue
}
