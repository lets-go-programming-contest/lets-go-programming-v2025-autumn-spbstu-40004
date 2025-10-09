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
	*h = old[:n-1]

	return x
}

func main() {
	var numberOfDishes int
	_, err := fmt.Scanln(&numberOfDishes)
	if err != nil {
		fmt.Println("invalid numberOfDishes")
	}

	myHeap := &MaxHeap{}
	heap.Init(myHeap)

	for range numberOfDishes {
		var dish int
		_, err := fmt.Scan(&dish)
		if err != nil {
			fmt.Println("invalid dish")
		}
		heap.Push(myHeap, dish)
	}

	var k int
	_, err1 := fmt.Scanln(&k)
	if err1 != nil {
		fmt.Println("invalid k")
	}

	var result int
	for range k {
		result = heap.Pop(myHeap).(int)
	}

	fmt.Println(result)
}
