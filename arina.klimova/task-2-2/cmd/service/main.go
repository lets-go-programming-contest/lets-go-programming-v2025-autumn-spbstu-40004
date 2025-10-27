package main

import (
	"container/heap"
	"fmt"

	maxheap "github.com/arinaklimova/task-2-2/internal/maxheap"
)

func main() {
	var (
		errScanDishes = fmt.Errorf("invalid numberOfDishes")
		errScanDish   = fmt.Errorf("invalid dish")
		errScanK      = fmt.Errorf("invalid k")
		errTypeAssert = fmt.Errorf("type assertion failed")
	)

	var numberOfDishes int

	_, err := fmt.Scanln(&numberOfDishes)
	if err != nil {
		fmt.Println(errScanDishes)

		return
	}

	myHeap := &maxheap.MaxHeap{}
	heap.Init(myHeap)

	for range numberOfDishes {
		var dish int

		_, err := fmt.Scan(&dish)
		if err != nil {
			fmt.Println(errScanDish)
		}

		heap.Push(myHeap, dish)
	}

	var kthLargest int

	_, err1 := fmt.Scanln(&kthLargest)
	if err1 != nil {
		fmt.Println(errScanK)

		return
	}

	var result int

	for range kthLargest {
		val, ok := heap.Pop(myHeap).(int)
		if !ok {
			fmt.Println(errTypeAssert)

			return
		}

		result = val
	}

	fmt.Println(result)
}
