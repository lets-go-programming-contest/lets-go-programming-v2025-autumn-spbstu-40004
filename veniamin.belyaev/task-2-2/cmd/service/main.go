package main

import (
	"container/heap"
	"fmt"

	"github.com/belyaevEDU/task-2-2/internal/MaxHeap"
)

func main() {
	var amount, kNumber, result int

	_, err := fmt.Scan(&amount)
	if err != nil {
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
	}

	_, err = fmt.Scan(&kNumber)
	if err != nil {
		fmt.Println("Invalid k")

		return
	}

	mealHeap := MaxHeap.InitHeap(mealArray)

	for range kNumber {
		val, TACheck := heap.Pop(mealHeap).(int)
		if TACheck {
			result = val
		}
	}

	fmt.Println(result)
}
