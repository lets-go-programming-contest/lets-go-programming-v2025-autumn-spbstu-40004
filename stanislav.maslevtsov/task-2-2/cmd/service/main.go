package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (heap *MaxHeap) Len() int {
	return len(*heap)
}

func (heap *MaxHeap) Less(left, right int) bool {
	return (*heap)[left] > (*heap)[right]
}

func (heap *MaxHeap) Swap(left, right int) {
	(*heap)[left], (*heap)[right] = (*heap)[right], (*heap)[left]
}

func (heap *MaxHeap) Push(value any) {
	assertedValue, ok := value.(int)
	if ok {
		*heap = append(*heap, assertedValue)
	}
}

func (heap *MaxHeap) Pop() any {
	oldHeap := *heap
	heapLen := len(oldHeap)
	lastValue := oldHeap[heapLen-1]
	*heap = oldHeap[0 : heapLen-1]

	return lastValue
}

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

	preferences := &MaxHeap{}
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
