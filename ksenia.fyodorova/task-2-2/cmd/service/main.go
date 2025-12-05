package main

import (
	"container/heap"
	"fmt"

	customheap "github.com/lolnyok/task-2-2/heap"
)

var (
	ErrInvalidInput       = fmt.Errorf("invalid input format")
	ErrUnexpectedHeapType = fmt.Errorf("unexpected type from heap")
)

func main() {
	var totalDishes, preferenceRank int

	_, err := fmt.Scan(&totalDishes)
	if err != nil {
		panic(fmt.Errorf("%w: %v", ErrInvalidInput, err))
	}

	dishRatings := make([]int, totalDishes)

	for i := 0; i < totalDishes; i++ {
		_, err := fmt.Scan(&dishRatings[i])
		if err != nil {
			panic(fmt.Errorf("%w: %v", ErrInvalidInput, err))
		}
	}

	_, err = fmt.Scan(&preferenceRank)
	if err != nil {
		panic(fmt.Errorf("%w: %v", ErrInvalidInput, err))
	}

	dishHeap := &customheap.DishHeap{}
	heap.Init(dishHeap)

	for _, rating := range dishRatings {
		heap.Push(dishHeap, rating)

		if dishHeap.Len() > preferenceRank {
			heap.Pop(dishHeap)
		}
	}

	result := heap.Pop(dishHeap)

	kthPreference, ok := result.(int)
	if !ok {
		panic(ErrUnexpectedHeapType)
	}

	fmt.Println(kthPreference)
}
