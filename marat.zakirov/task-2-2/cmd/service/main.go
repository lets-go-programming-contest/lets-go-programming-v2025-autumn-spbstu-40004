package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/ZakirovMS/task-2-2/internal/maxheap"
)

var errorInput = errors.New("ERROR Incorrect input")

func main() {
	var (
		dishesNum   int
		currentPref int
		orderPred   int
		customHeap  maxheap.MaxHeap
	)

	const (
		minDishQuantity = 1
		maxDishQuantity = 10000
		minPrefQunatity = -10000
		maxPrefQunatity = 10000
	)

	_, err := fmt.Scan(&dishesNum)
	if err != nil || dishesNum < minDishQuantity || dishesNum > maxDishQuantity {
		fmt.Println(errorInput)

		return
	}

	heap.Init(&customHeap)

	for range dishesNum {
		_, err = fmt.Scan(&currentPref)
		if err != nil || currentPref < minPrefQunatity || currentPref > maxPrefQunatity {
			fmt.Println(errorInput)

			return
		}

		heap.Push(&customHeap, currentPref)
	}

	_, err = fmt.Scan(&orderPred)
	if err != nil || orderPred < 1 || orderPred > customHeap.Len() {
		fmt.Println(errorInput)

		return
	}

	for range orderPred - 1 {
		heap.Pop(&customHeap)
	}

	result, ok := heap.Pop(&customHeap).(int)
	if ok {
		fmt.Println(result)
	}
}
