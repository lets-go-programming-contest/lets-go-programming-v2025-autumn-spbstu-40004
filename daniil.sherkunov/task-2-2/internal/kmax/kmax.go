package kmax

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int { return len(*h) }

func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }

func (h *MinHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(x any) {
	v, ok := x.(int)
	if !ok {
		panic("kmax: Push non-int")
	}
	*h = append(*h, v)
}

func (h *MinHeap) Pop() any {
	old := *h
	if len(old) == 0 {
		panic("kmax: Pop from empty heap")
	}
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func (h *MinHeap) Top() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}
	return (*h)[0], true
}

func KthLargest(values []int, kth int) (int, bool) {
	mh := &MinHeap{}
	heap.Init(mh)
	for _, value := range values {
		if mh.Len() < kth {
			heap.Push(mh, value)
			continue
		}
		top, _ := mh.Top()
		if value > top {
			heap.Pop(mh)
			heap.Push(mh, value)
		}
	}
	res, ok := mh.Top()
	if !ok {
		return 0, false
	}
	return res, true
}
