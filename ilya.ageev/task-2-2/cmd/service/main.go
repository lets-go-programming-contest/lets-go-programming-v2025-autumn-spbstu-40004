package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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
	var N, k int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&k)

	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}

	for i := k; i < N; i++ {
		if arr[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}

	result := (*h)[0]
	fmt.Println(result)
}
