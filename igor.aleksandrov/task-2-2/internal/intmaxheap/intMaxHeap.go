package intmaxheap

import "container/heap"

type IntMaxHeap []int

func (h *IntMaxHeap) Len() int {
	return len(*h)
}

func (h *IntMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMaxHeap) Push(element interface{}) {
	num, success := element.(int)
	if success {
		panic("push expects an int element")
	} else {
		*h = append(*h, num)
	}
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	length := len(old)

	if length == 0 {
		return nil
	}

	element := old[length-1]
	*h = old[0 : length-1]

	return element
}

func InitIntMaxHeap(array []int) *IntMaxHeap {
	h := &IntMaxHeap{}
	*h = array
	heap.Init(h)

	return h
}
