package main

import (
	"container/heap"
	"fmt"

	intheap "github.com/15446-rus75/task-2-2/internal/heap"
)

const (
	minElem = -10000
	maxElem = 10000
)

func main() {
	var nCount, kCount int

	_, err := fmt.Scan(&nCount)
	if err != nil {
		fmt.Print("Failed to read N\n")

		return
	}

	arr := make([]int, nCount)

	for count := range nCount {
		_, err = fmt.Scan(&arr[count])
		if err != nil {
			fmt.Print("Failed to read data\n")
		}

		if arr[count] < minElem || arr[count] > maxElem {
			fmt.Print("Invalid arr[i]\n")

			return
		}
	}

	_, err = fmt.Scan(&kCount)
	if err != nil || kCount < 1 || kCount > nCount {
		fmt.Print("Failed to read K\n")
	}

	heapOfMeals := intheap.NewIntHeap()

	for _, num := range arr {
		heap.Push(heapOfMeals, num)
	}

	var result int

	for range kCount {
		if value, good := heap.Pop(heapOfMeals).(int); good {
			result = value
		}
	}

	fmt.Println(result)
}
