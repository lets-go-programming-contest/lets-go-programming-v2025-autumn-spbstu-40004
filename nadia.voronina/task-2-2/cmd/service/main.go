package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
func main() {
	var numberOfDishes int

	_, errN := fmt.Scanln(&numberOfDishes)
	if errN != nil {
		fmt.Println("Invalid number of dishes")

		return
	}

	var dishes IntHeap
	heap.Init(&dishes)
	var dish int

	for range numberOfDishes {

		_, errK := fmt.Scan(&dish)
		if errK != nil {
			fmt.Println("Invalid dish")

			return
		}

		heap.Push(&dishes, dish)
	}

	var preferredDish int
	_, errK := fmt.Scan(&preferredDish)
	if errK != nil {
		fmt.Println("Invalid preferred dish")

		return
	}

	for range numberOfDishes - preferredDish {
		heap.Pop(&dishes)
	}
	fmt.Printf("%d.", heap.Pop(&dishes))
}
