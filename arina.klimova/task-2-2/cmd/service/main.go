package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeap) Push(x interface{}) {
	val, ok := x.(int)
	if !ok {
		fmt.Println("Type assertion failed")

		return
	}

	*h = append(*h, val)
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

		return
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

	var kthLargest int

	_, err1 := fmt.Scanln(&kthLargest)
	if err1 != nil {
		fmt.Println("invalid k")

		return
	}

	var result int

	for range kthLargest {
		val, ok := heap.Pop(myHeap).(int)
		if !ok {
			fmt.Println("Type assertion failed")

			return
		}

		result = val
	}

	fmt.Println(result)
}
