package main

import (
	"fmt"
	"container/heap"

	IntHeap "github.com/faxryzen/task-2-2/internal/IntHeap"
)

func main() {
	var (
		amount, kPrefer uint16
		result          int
	)

	_, err := fmt.Scanln(&amount)
	if err != nil || amount == 0 || amount > 10000 {
		fmt.Println("invalid food amount")

		return
	}

	foodRating := make([]int, amount)

	for i := range amount {
		_, err = fmt.Scanln(&foodRating[i])
		if err != nil || foodRating[i] < -10000 || foodRating[i] > 10000 {
			fmt.Println("invalid food init")

			return
		}
	}

	_, err = fmt.Scan(&kPrefer)
	if err != nil || kPrefer == 0 || kPrefer > amount {
		fmt.Println("invalid preference")

		return
	}

	foodHeap := IntHeap.InitIntHeap(foodRating)

	for range kPrefer {
		value, isGood := heap.Pop(foodHeap).(int)
		if isGood {
			result = value
		}
	}

	fmt.Println(result)
}
