package main

import (
	"container/heap"
	"fmt"

	"github.com/jambii1/task-2-2/internal/maxheap"
)

func main() {
	var (
		dishesAmount uint
		dishPriority int
		preference   uint
	)

	_, err := fmt.Scan(&dishesAmount)
	if err != nil {
		fmt.Println("invalid number of dishes")

		return
	}

	preferences := &maxheap.MaxHeap{}
	heap.Init(preferences)

	for range dishesAmount {
		_, err = fmt.Scan(&dishPriority)
		if err != nil {
			fmt.Println("invalid dish priority")

			return
		}

		heap.Push(preferences, dishPriority)
	}

	_, err = fmt.Scan(&preference)
	if err != nil {
		fmt.Println("invalid dish preference")

		return
	}

	for range preference - 1 {
		heap.Pop(preferences)
	}

	result, ok := heap.Pop(preferences).(int)
	if ok {
		fmt.Println(result)
	}
}
