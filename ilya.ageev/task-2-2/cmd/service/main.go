package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (DishHeap IntHeap) Len() int           { return len(DishHeap) }
func (DishHeap IntHeap) Less(i, j int) bool { return DishHeap[i] < DishHeap[j] }
func (DishHeap IntHeap) Swap(i, j int)      { DishHeap[i], DishHeap[j] = DishHeap[j], DishHeap[i] }

func (DishHeap *IntHeap) Push(x interface{}) {
	*DishHeap = append(*DishHeap, x.(int))
}

func (DishHeap *IntHeap) Pop() interface{} {
	old := *DishHeap
	n := len(old)
	x := old[n-1]
	*DishHeap = old[0 : n-1]
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
			fmt.Println("Invalid priority")
			return
		}
	}

	fmt.Scan(&Priority)

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
