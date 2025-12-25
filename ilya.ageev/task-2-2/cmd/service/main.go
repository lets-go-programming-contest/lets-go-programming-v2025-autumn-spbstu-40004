package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (dishHeap *IntHeap) Len() int {
	return len(*dishHeap)
}

func (dishHeap *IntHeap) Less(i, j int) bool {
	return (*dishHeap)[i] < (*dishHeap)[j]
}

func (dishHeap *IntHeap) Swap(i, j int) {
	(*dishHeap)[i], (*dishHeap)[j] = (*dishHeap)[j], (*dishHeap)[i]
}

func (dishHeap *IntHeap) Push(x interface{}) {
	if num, ok := x.(int); ok {
		*dishHeap = append(*dishHeap, num)
	} else {
		fmt.Println("Error: expected int type")
	}
}

func (dishHeap *IntHeap) Pop() interface{} {
	old := *dishHeap
	n := len(old)
	x := old[n-1]
	*dishHeap = old[0 : n-1]

	return x
}

func main() {
	var NumOfDishes, Priority int

	_, err := fmt.Scan(&NumOfDishes)
	if err != nil {
		fmt.Println("Invalid number of dishes")

		return
	}

	arr := make([]int, NumOfDishes)
	for i := range arr {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			fmt.Println("Invalid priority of dishes")

			return
		}
	}

	_, err = fmt.Scan(&Priority)
	if err != nil {
		fmt.Println("Invalid priority dish")

		return
	}

	if Priority > len(arr) {
		fmt.Println("Priority exceeds number of dishes")

		return
	}

	DishHeap := &IntHeap{}
	heap.Init(DishHeap)

	for _, num := range arr[:Priority] {
		heap.Push(DishHeap, num)
	}

	for _, num := range arr[Priority:] {
		if num > (*DishHeap)[0] {
			heap.Pop(DishHeap)
			heap.Push(DishHeap, num)
		}
	}

	result := (*DishHeap)[0]
	fmt.Println(result)
}
