package main

import (
	"container/heap"
	"fmt"

	intHeap "github.com/cherkasoov/lets-go-programming-v2025-autumn-spbstu-40004/task-2-2/internal/intheap"
)

func main() {
	var sum, count, res int

	_, err := fmt.Scan(&sum)
	if err != nil || sum < 1 {
		fmt.Println("Invalid sum")

		return
	}

	mealArray := make([]int, sum)
	for index := range sum {
		_, err = fmt.Scan(&mealArray[index])
		if err != nil {
			fmt.Println("Invalid data")

			return
		}

		if mealArray[index] < -10000 || mealArray[index] > 10000 {
			fmt.Println("Invalid data")

			return
		}
	}

	_, err = fmt.Scan(&count)
	if err != nil || count < 1 || count > sum {
		fmt.Println("Invalid k")

		return
	}

	mealHeap := intHeap.InitHeap(mealArray)

	for range count {
		val, TACheck := heap.Pop(mealHeap).(int)
		if TACheck {
			res = val
		}
	}

	fmt.Println(res)
}
