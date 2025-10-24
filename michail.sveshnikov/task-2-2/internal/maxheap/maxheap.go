package maxheap

type MaxHeap []int

func (heap *MaxHeap) Push(val any) {
	value, ok := val.(int)
	if !ok {
		panic("MaxHeap.Push: value must be int")
	}

	*heap = append(*heap, value)
}

func (heap *MaxHeap) Pop() any {
	old := *heap

	lenHeap := len(old)
	if lenHeap == 0 {
		return nil
	}

	element := (old)[lenHeap-1]
	*heap = (old)[0 : lenHeap-1]

	return element
}

func (heap *MaxHeap) Len() int {
	return len(*heap)
}

func (heap *MaxHeap) Less(i, j int) bool {
	return (*heap)[i] > (*heap)[j]
}

func (heap *MaxHeap) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}
