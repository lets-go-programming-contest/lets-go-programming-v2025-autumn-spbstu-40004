package maxheap

type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeap) Push(x interface{}) {
	val, ok := x.(int)
	if !ok {
		panic("Type assertion failed: expected int")
	}

	*h = append(*h, val)
}

func (h *MaxHeap) Pop() interface{} {
	if len(*h) == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}
