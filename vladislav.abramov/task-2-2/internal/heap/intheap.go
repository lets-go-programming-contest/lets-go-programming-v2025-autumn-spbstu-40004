package heap

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	if num, good := x.(int); good {
		*h = append(*h, num)
	} else {
		panic("Expected int")
	}
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	length := len(old)

	if length == 0 {
		panic("Heap is empty")
	}

	x := old[length-1]
	*h = old[0 : length-1]

	return x
}

func NewIntHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)

	return h
}
