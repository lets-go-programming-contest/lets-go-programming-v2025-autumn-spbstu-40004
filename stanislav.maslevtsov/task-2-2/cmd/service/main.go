package main

import (
	"fmt"
)

type MinHeap []uint

func (heap MinHeap) Len() int {
	return len(heap)
}

func (heap MinHeap) Less(left, right int) bool {
	return heap[left] < heap[right]
}

func (heap MinHeap) Swap(left, right int) {
	heap[left], heap[right] = heap[right], heap[left]
}

func (heap *MinHeap) Push(value any) {
	*heap = append(*heap, value.(uint))
}

func (heap *MinHeap) Pop() any {
	oldHeap := *heap
	heapLen := len(oldHeap)
	lastValue := oldHeap[heapLen-1]
	*heap = oldHeap[0 : heapLen-1]
	return lastValue
}

func main() {
	var (
		dishesAmount int
	)

	_, err := fmt.Scan(&dishesAmount)
	if err != nil {
		fmt.Println("invalid number of dishes")

		return
	}
}
