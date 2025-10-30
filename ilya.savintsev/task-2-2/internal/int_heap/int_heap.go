package intheap

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] >= (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	n, isGood := x.(int)
	if !isGood {
		panic("expected int")
	}

	*h = append(*h, n)
}

func (h *IntHeap) Pop() any {
	if len(*h) == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func InitIntHeap(array []int) *IntHeap {
	h := &IntHeap{}
	*h = array
	heap.Init(h)

	return h
}
