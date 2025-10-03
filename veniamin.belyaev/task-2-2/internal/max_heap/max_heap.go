package max_heap

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	num, TACheck := x.(int)
	if TACheck {
		*h = append(*h, num)
	}
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	if n == 0 {
		return nil
	}

	x := old[n-1]
	*h = old[:n-1]

	return x
}

func InitHeap(array []int) *MaxHeap {
	maxHeap := &MaxHeap{}
	*maxHeap = array
	heap.Init(maxHeap)

	return maxHeap
}
