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
	_, scanErr := fmt.Scan(&dishCount)
	if scanErr != nil {
		return
	}

	dishRatings := make([]int, dishCount)
	for index := range dishCount {
		_, scanErr = fmt.Scan(&dishRatings[index])
		if scanErr != nil {
			return
		}
	}

	_, scanErr = fmt.Scan(&kthPreference)
	if scanErr != nil {
		return
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, rating := range dishRatings {
		heap.Push(maxHeap, rating)
	}

	var result int
	for range kthPreference {
		result = heap.Pop(maxHeap).(int)
	}

	fmt.Println(result)
}
