package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var nCount, kCount int

	fmt.Scan(&nCount)

	arr := make([]int, nCount)
	for count := 0; count < nCount; count++ {
		fmt.Scan(&arr[count])
	}

	fmt.Scan(&kCount)

	heap_of_meals := &IntHeap{}
	heap.Init(heap_of_meals)

	for _, num := range arr {
		heap.Push(heap_of_meals, num)
	}

	var result int

	for range kCount {
		result = heap.Pop(heap_of_meals).(int)
	}

	fmt.Println(result)
}
