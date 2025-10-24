package main

import (
	"container/heap"
	"fmt"

	maxheap "github.com/arinaklimova/task-2-2/internal/maxheap"
)

func main() {
	var numberOfDishes int

	_, err := fmt.Scanln(&numberOfDishes)
	if err != nil {
		fmt.Println("invalid numberOfDishes")

		return
	}

	myHeap := &maxheap.MaxHeap{}
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
