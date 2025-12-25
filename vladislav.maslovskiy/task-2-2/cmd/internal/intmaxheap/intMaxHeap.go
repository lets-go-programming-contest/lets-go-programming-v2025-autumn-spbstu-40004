package intmaxheap

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Size() int {
	return len(*h)
}

func (h *MaxHeap) Compare(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(element interface{}) {
	num, success := element.(int)
	if success {
		panic("push expects an int element")
	} else {
		*h = append(*h, num)
	}
}

func (h *MaxHeap) Extract() interface{} {
	old := *h
	length := len(old)

	if length == 0 {
		return nil
	}

	element := old[length-1]
	*h = old[0 : length-1]

	return element
}

func CreateMaxHeap(array []int) *MaxHeap {
	h := &MaxHeap{}
	*h = array
	heap.Init(h)

	return h
}
