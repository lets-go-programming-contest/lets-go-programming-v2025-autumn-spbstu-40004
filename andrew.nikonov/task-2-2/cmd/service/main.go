package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/ysffmn/task-2-2/internal/intmaxheap"
)

var errInput = errors.New("incorrect input")

func main() {
	var dishesNum int

	_, err := fmt.Scanln(&dishesNum)
	if err != nil || dishesNum < -10000 || dishesNum > 10000 {
		fmt.Println(errInput)

		return
	}

	dishesHeap := &intmaxheap.IntMaxHeap{}
	heap.Init(dishesHeap)

	for range dishesNum {
		var rating int

		_, err = fmt.Scan(&rating)
		if err != nil || rating < -10000 || rating > 10000 {
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
		dishesHeap.Pop()
	}

	fmt.Println(dishesHeap.Pop())
}
