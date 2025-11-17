package intheap

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i >= len(*h) || j >= len(*h) {
		return
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	num, TACheck := x.(int)
	if TACheck {
		*h = append(*h, num)
	} else {
		panic("type assertion failed")
	}
}

func (h *IntHeap) Pop() interface{} {
	old := *h

	length := len(old)
	if length == 0 {
		return nil
	}

	element := old[length-1]
	*h = old[:length-1]

	return element
}

func InitHeap(array []int) *IntHeap {
	maxHeap := &IntHeap{}
	*maxHeap = array
	heap.Init(maxHeap)

	return maxHeap
}
