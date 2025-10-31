package main

import (
	"container/heap"
	"fmt"

	"github.com/Z-1337/task-2-2/internal/maxheap"
)

func main() {
	var (
		amount     uint
		priority   int
		preference uint
	)

	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Error while reading amount of meals")

		return
	}

	preferences := &maxheap.MaxHeap{}
	heap.Init(preferences)

	for range amount {
		_, err = fmt.Scan(&priority)
		if err != nil {
			fmt.Println("Error while reading meal priority")

			return
		}

		heap.Push(preferences, priority)
	}

	_, err = fmt.Scan(&preference)
	if err != nil {
		fmt.Println("Error while reading meal preference")

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
