package main

import (
	"container/heap"
	"fmt"
	"log"
)

type DishHeap []int

func (heap *DishHeap) Len() int           { return len(*heap) }
func (heap *DishHeap) Less(i, j int) bool { return (*heap)[i] < (*heap)[j] }
func (heap *DishHeap) Swap(i, j int)      { (*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i] }

func (heap *DishHeap) Push(dish interface{}) {
	dishValue, ok := dish.(int)
	if !ok {
		log.Fatal("invalid type in Push")
	}

	*heap = append(*heap, dishValue)
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

	_, err := fmt.Scan(&totalDishes)
	if err != nil {
		log.Fatal(err)
	}

	dishRatings := make([]int, totalDishes)

	for i := range totalDishes {
		_, err := fmt.Scan(&dishRatings[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = fmt.Scan(&preferenceRank)
	if err != nil {
		log.Fatal(err)
	}

	dishHeap := &DishHeap{}
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
		log.Fatal("invalid type from heap")
	}

	fmt.Println(kthPreference)
}
