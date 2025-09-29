package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Ksenia-rgb/task-2-2/internal/maxheap"
)

var errInput = errors.New("incorrect input")

func main() {
	var cntDish int
	heapDish := &maxheap.MaxHeap{}
	heap.Init(heapDish)

	_, err := fmt.Scanln(&cntDish)
	if err != nil || cntDish < 1 || cntDish > 10000 {
		fmt.Println(errInput)

		return
	}

	for range cntDish {
		var valueDish int

		_, err = fmt.Scan(&valueDish)
		if err != nil || valueDish < -10000 || valueDish > 10000 {
			fmt.Println(errInput)

			return
		}

		heap.Push(heapDish, valueDish)
	}

	var ratingDish int

	_, err = fmt.Scan(&ratingDish)
	if err != nil || ratingDish < 1 || ratingDish > heapDish.Len() {
		fmt.Println(errInput)

		return
	}

	for range ratingDish - 1 {
		heap.Pop(heapDish)
	}

	fmt.Println(heap.Pop(heapDish))
}
