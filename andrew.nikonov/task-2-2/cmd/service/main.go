package main

import (
	"container/heap"
	"errors"
	"fmt"

	intmaxheap "github.com/ysffmn/task-2-2/internal/intMaxHeap"
)

var errInput = errors.New("incorrect input")

const (
	maxInputVal = 10000
	minInputVal = -10000
)

func main() {
	var dishesNum int

	_, err := fmt.Scanln(&dishesNum)
	if err != nil || dishesNum < minInputVal || dishesNum > maxInputVal {
		fmt.Println(errInput)

		return
	}

	dishesHeap := &intmaxheap.IntMaxHeap{}
	heap.Init(dishesHeap)

	for range dishesNum {
		var rating int

		_, err = fmt.Scan(&rating)
		if err != nil || rating < minInputVal || rating > maxInputVal {
			fmt.Println(errInput)

			return
		}

		heap.Push(dishesHeap, rating)
	}

	var preferredNum int

	_, err = fmt.Scanln(&preferredNum)
	if err != nil || preferredNum < 1 || preferredNum > dishesNum {
		fmt.Println(errInput)

		return
	}

	for range preferredNum - 1 {
		heap.Pop(dishesHeap)
	}

	fmt.Println(heap.Pop(dishesHeap))
}
