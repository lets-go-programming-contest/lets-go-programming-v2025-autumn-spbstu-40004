package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/svesh3000/task-2-2/internal/maxheap"
)

var (
	errInvalidDishNum       = errors.New("invalid number of dishes")
	errInvalidDishPrefLevel = errors.New("invalid preference level of the dish")
	errInvalidKthDishNumber = errors.New("invalid number of the kth dish by preference")
)

func main() {
	const (
		minDishNum = 1
	)

	var numDishes int

	_, err := fmt.Scan(&numDishes)
	if err != nil || numDishes < minDishNum {
		fmt.Println(errInvalidDishNum)

		return
	}

	buffet := &maxheap.MaxHeap{}
	heap.Init(buffet)

	for range numDishes {
		var preference int

		_, err = fmt.Scan(&preference)
		if err != nil {
			fmt.Println(errInvalidDishPrefLevel)

			return
		}

		heap.Push(buffet, preference)
	}

	var num int

	_, err = fmt.Scan(&num)
	if err != nil || num < minDishNum || num > numDishes {
		fmt.Println(errInvalidKthDishNumber)

		return
	}

	for range num - 1 {
		heap.Pop(buffet)
	}

	var dish int

	dish, ok := heap.Pop(buffet).(int)
	if ok {
		fmt.Println(dish)
	}
}
