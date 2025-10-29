package main

import (
	"container/heap"
	"fmt"

	intMaxHeap "github.com/MrMels625/task-2-2/internal/intmaxheap"
)

func readEmployeeMind(dishes []int, preferredDishNumber int) int {
	var resultDish int

	dishesHeap := intMaxHeap.InitIntMaxHeap(dishes)

	for range preferredDishNumber {
		dish, popped := heap.Pop(dishesHeap).(int)
		if popped {
			resultDish = dish
		}
	}

	return resultDish
}

func main() {
	var dishesCount int

	_, err := fmt.Scanln(&dishesCount)
	if err != nil || dishesCount < 1 || dishesCount > 10000 {
		fmt.Println("Invalid dishes count!")

		return
	}

	dishes := make([]int, dishesCount)

	for index := range dishesCount {
		_, err = fmt.Scan(&dishes[index])
		if err != nil {
			fmt.Println("Invalid dish description!")

			return
		}

		if dishes[index] < -10000 || dishes[index] > 10000 {
			fmt.Println("Unsupported dish description!")

			return
		}
	}

	var preferredDishNumber int

	_, err = fmt.Scanln(&preferredDishNumber)
	if err != nil || preferredDishNumber < 1 || preferredDishNumber > dishesCount {
		fmt.Println("Invalid number of a dish in buffet!")

		return
	}

	fmt.Println(readEmployeeMind(dishes, preferredDishNumber))
}
