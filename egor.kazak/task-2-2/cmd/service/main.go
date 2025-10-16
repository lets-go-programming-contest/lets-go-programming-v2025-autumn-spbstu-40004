package main

import (
	"container/heap"
	"fmt"
	"log"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	num, ok := x.(int)
	if !ok {
		log.Fatal("type assertion to int failed")
	}
	*h = append(*h, num)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	var dishCount, preferenceOrder int
	_, err := fmt.Scan(&dishCount)
	if err != nil {
		log.Fatal(err)
	}

	ratings := make([]int, dishCount)
	for index := range ratings {
		_, err := fmt.Scan(&ratings[index])
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = fmt.Scan(&preferenceOrder)
	if err != nil {
		log.Fatal(err)
	}

	heapInstance := &IntHeap{}
	heap.Init(heapInstance)

	for _, rating := range ratings {
		heap.Push(heapInstance, rating)
	}

	var result int
	for range preferenceOrder {
		item := heap.Pop(heapInstance)
		num, ok := item.(int)
		if !ok {
			log.Fatal("type assertion to int failed")
		}
		result = num
	}

	fmt.Println(result)
}
