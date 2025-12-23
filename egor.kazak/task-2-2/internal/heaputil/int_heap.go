package heaputil

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int { return len(*h) }

func (h *IntHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] } // max-heap

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	if num, ok := x.(int); ok {
		*h = append(*h, num)
	} else {
		panic("heaputil: Push received non-int value")
	}
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	oldLen := len(old)

	if oldLen == 0 {
		panic("heaputil: Pop called on empty heap")
	}

	x := old[oldLen-1]
	*h = old[0 : oldLen-1]

	return x
}

func Init(h *IntHeap) {
	heap.Init(h)
}

func Push(h *IntHeap, x int) {
	heap.Push(h, x)
}

func Pop(h *IntHeap) int {
	item := heap.Pop(h)
	if num, ok := item.(int); ok {
		return num
	}
	panic("heaputil: Pop returned non-int value")
}
