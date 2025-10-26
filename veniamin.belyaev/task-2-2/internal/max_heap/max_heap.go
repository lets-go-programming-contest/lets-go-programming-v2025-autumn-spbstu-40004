package maxheap

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	if i >= len(*h) || j >= len(*h) {
		return
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	num, TACheck := x.(int)
	if TACheck {
		*h = append(*h, num)
	} else {
		panic("type assertion failed")
	}
}

func (h *MaxHeap) Pop() interface{} {
	old := *h

	length := len(old)
	if length == 0 {
		return nil
	}

	element := old[length-1]
	*h = old[:length-1]

	return element
}

func InitHeap(array []int) *MaxHeap {
	maxHeap := &MaxHeap{}
	*maxHeap = array
	heap.Init(maxHeap)

	return maxHeap
}
