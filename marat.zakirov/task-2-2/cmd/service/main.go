package main

import (
	"container/heap"
	"fmt"

	"github.com/ZakirovMS/task-2-2/internal/maxheap"
)

func main() {
	var (
		dishesNum   int
		currentPref int
		orderPred   int
		customHeap  maxheap.MaxHeap
	)

	_, err := fmt.Scan(&dishesNum)
	if err != nil {
		fmt.Println("ERROR Incorrect number of dishes")

		return
	}

	heap.Init(&customHeap)

	for range dishesNum {
		_, err = fmt.Scan(&currentPref)
		if err != nil {
			fmt.Println("ERROR Incorrect dish priority")

			return
		}

		heap.Push(&customHeap, currentPref)
	}

	_, err = fmt.Scan(&orderPred)
	if err != nil {
		fmt.Println("ERROR Incorrect employee priority")

		return
	}

	for range orderPred - 1 {
		heap.Pop(&customHeap)
	}

	result, ok := heap.Pop(&customHeap).(int)
	if ok {
		fmt.Println(result)
	}
}
