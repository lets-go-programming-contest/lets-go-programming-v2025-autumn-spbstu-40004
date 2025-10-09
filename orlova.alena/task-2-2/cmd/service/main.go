package main

import (
	"container/heap"
	"fmt"

	intheap "task-2-2/cmd/service/internal/intheap"
)

func main() {
	var (
		numOfDishes   int
		rate          int
		numOfPrefDish int
		myHeap        intheap.IntHeap
	)

	_, err := fmt.Scanln(&numOfDishes)
	if err != nil {
		fmt.Println("Wrong input")

		return
	}

	heap.Init(&myHeap)

	for range numOfDishes {
		_, err = fmt.Scanln(&rate)
		if err != nil {
			fmt.Println("Wrong input")

			return
		}

		heap.Push(&myHeap, rate)
	}

	_, err = fmt.Scanln(&numOfPrefDish)
	if err != nil || numOfPrefDish > myHeap.Len() {
		fmt.Println("Wrong input")

		return
	}

	for range numOfPrefDish - 1 {
		heap.Pop(&myHeap)
	}

	fmt.Println(heap.Pop(&myHeap))
}
