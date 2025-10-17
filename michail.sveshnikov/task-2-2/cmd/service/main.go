package main

import (
	"container/heap"
	"fmt"

	"github.com/svesh3000/task-2-2/internal/maxheap"
)

func main() {
	var numDishes int

	_, err := fmt.Scan(&numDishes)
	if err != nil || numDishes < 1 {
		fmt.Println("ERROR: invalid number of dishes!")

		return
	}

	buffet := &maxheap.MaxHeap{}
	heap.Init(buffet)

	for range numDishes {
		var preference int

		_, err = fmt.Scan(&preference)
		if err != nil {
			fmt.Println("ERROR: invalid preference level of the dish!")

			return
		}

		heap.Push(buffet, preference)
	}

	var num int

	_, err = fmt.Scan(&num)
	if err != nil || num < 1 || num > numDishes {
		fmt.Println("ERROR: invalid number of the kth dish by preference!")

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
