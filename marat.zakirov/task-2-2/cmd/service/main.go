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
		h           maxheap.MaxHeap
	)

	_, err := fmt.Scan(&dishesNum)
	if err != nil {
		fmt.Println("ERROR Incorrect number of dishes")

		return
	}

	heap.Init(&h)

	for range dishesNum {
		_, err = fmt.Scan(&currentPref)
		if err != nil {
			fmt.Println("ERROR Incorrect dish priority")

			return
		}

		heap.Push(&h, currentPref)
	}

	_, err = fmt.Scan(&orderPred)
	if err != nil {
		fmt.Println("ERROR Incorrect employee priority")

		return
	}

	for range orderPred - 1 {
		heap.Pop(&h)
	}

	result, ok := heap.Pop(&h).(int)
	if ok {
		fmt.Println(result)
	}
}
