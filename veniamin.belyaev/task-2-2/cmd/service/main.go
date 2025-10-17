package main

import (
	"container/heap"
	"fmt"

	maxHeap "github.com/belyaevEDU/task-2-2/internal/max_heap"
)

func main() {
	var amount, kNumber, result int

	_, err := fmt.Scan(&amount)
	if err != nil || amount < 1 {
		fmt.Println("Invalid amount")

		return
	}

	mealArray := make([]int, amount)
	for index := range amount {
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

	_, err = fmt.Scan(&kNumber)
	if err != nil || kNumber < 1 || kNumber > amount {
		fmt.Println("Invalid k")

		return
	}

	mealHeap := maxHeap.InitHeap(mealArray)

	for range kNumber {
		val, TACheck := heap.Pop(mealHeap).(int)
		if TACheck {
			result = val
		}
	}

	fmt.Println(result)
}
