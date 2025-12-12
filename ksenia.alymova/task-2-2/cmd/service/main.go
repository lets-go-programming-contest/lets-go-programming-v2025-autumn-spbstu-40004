package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Ksenia-rgb/task-2-2/internal/maxheap"
)

var errInput = errors.New("incorrect input")

const (
	minDish      = 1
	maxDish      = 10000
	minValueDish = -10000
	maxValueDish = 10000
)

func main() {
	var cntDish int

	_, err := fmt.Scanln(&cntDish)
	if err != nil || cntDish < minDish || cntDish > maxDish {
		fmt.Println(errInput)

		return
	}

	heapDish := &maxheap.MaxHeap{}
	heap.Init(heapDish)

	for range cntDish {
		var valueDish int

		_, err = fmt.Scan(&valueDish)
		if err != nil || valueDish < minValueDish || valueDish > maxValueDish {
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
