package main

import (
	"container/heap"
	"fmt"
)

type DishHeap []int

func (heap DishHeap) Len() int           { return len(heap) }
func (heap DishHeap) Less(i, j int) bool { return heap[i] < heap[j] }
func (heap DishHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }

func (heap *DishHeap) Push(dish interface{}) {
	*heap = append(*heap, dish.(int))
}

func (heap *DishHeap) Pop() interface{} {
	oldHeap := *heap
	lastIndex := len(oldHeap) - 1
	dish := oldHeap[lastIndex]
	*heap = oldHeap[:lastIndex]
	return dish
}

func main() {
	var totalDishes, preferenceRank int
	fmt.Scan(&totalDishes)

	dishRatings := make([]int, totalDishes)
	for i := 0; i < totalDishes; i++ {
		fmt.Scan(&dishRatings[i])
	}

	fmt.Scan(&preferenceRank)

	dishHeap := &DishHeap{}
	heap.Init(dishHeap)

	for _, rating := range dishRatings {
		heap.Push(dishHeap, rating)
		if dishHeap.Len() > preferenceRank {
			heap.Pop(dishHeap)
		}
	}

	kthPreference := heap.Pop(dishHeap).(int)
	fmt.Println(kthPreference)
}
