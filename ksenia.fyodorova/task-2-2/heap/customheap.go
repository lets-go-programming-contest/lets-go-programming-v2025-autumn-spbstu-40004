package customheap

type DishHeap []int

func (h *DishHeap) Len() int           { return len(*h) }
func (h *DishHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *DishHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *DishHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		panic("invalid type in Push: expected int")
	}

	*h = append(*h, value)
}

func (h *DishHeap) Pop() interface{} {
	if len(*h) == 0 {
		panic("Pop called on empty heap")
	}

	old := *h
	n := len(old)
	value := old[n-1]
	*h = old[:n-1]

	return value
}
