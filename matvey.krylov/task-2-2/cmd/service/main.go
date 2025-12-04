package main

import (
	"container/heap"
	"fmt"

	"github.com/mkryloff/task-2-2/internal/maxheap"
)

func main() {
	var (
		mealsAmount  uint
		mealPriority int
		preference   uint
	)

	_, err := fmt.Scan(&mealsAmount)
	if err != nil {
		fmt.Println("Invalid number of meals")

		return
	}

	preferences := &maxheap.MaxHeap{}
	heap.Init(preferences)

	for range mealsAmount {
		_, err = fmt.Scan(&mealPriority)
		if err != nil {
			fmt.Println("Invalid meal priority")

			return
		}

		heap.Push(preferences, mealPriority)
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
