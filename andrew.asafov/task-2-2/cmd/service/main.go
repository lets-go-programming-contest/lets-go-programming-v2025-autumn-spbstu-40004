package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var dishCount, kthPreference int
	_, err := fmt.Scan(&dishCount)
	if err != nil {
		return
	}

	arr := make([]int, dishCount)
	for i := 0; i < dishCount; i++ {
		_, err = fmt.Scan(&arr[i])
		if err != nil {
			return
		}
	}

	_, err = fmt.Scan(&kthPreference)
	if err != nil {
		return
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for i := 0; i < dishCount; i++ {
		heap.Push(maxHeap, arr[i])
	}

	var result int
	for i := 0; i < kthPreference; i++ {
		result = heap.Pop(maxHeap).(int)
	}

	fmt.Println(result)
}
