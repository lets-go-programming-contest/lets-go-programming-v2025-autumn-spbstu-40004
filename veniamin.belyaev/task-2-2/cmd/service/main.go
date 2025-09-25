package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func initHeap(array []int) *MaxHeap {
	maxHeap := &MaxHeap{}
	*maxHeap = array
	heap.Init(maxHeap)

	return maxHeap
}

func main() {
	var amount, kNumber, result int

	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Invalid amount")

		return
	}

	mealArray := make([]int, amount)
	for index := range amount {
		_, err = fmt.Scan(&mealArray[index])
		if err != nil {
			fmt.Println("Invalid data")

			return
		}
	}

	_, err = fmt.Scan(&kNumber)
	if err != nil {
		fmt.Println("Invalid k")

		return
	}

	mealHeap := initHeap(mealArray)

	for range kNumber {
		result = heap.Pop(mealHeap).(int)
	}

	fmt.Println(result)
}
