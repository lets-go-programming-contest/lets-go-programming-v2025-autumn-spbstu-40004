package main

import (
	"container/heap"
	"errors"
	"fmt"

	intheap "github.com/faxryzen/task-2-2/internal/int_heap"
)

const (
	minRating = -10000
	maxRating = 10000
	maxAmount = 10000
)

var (
	ErrInvalidFoodAmount = errors.New("invalid food amount")
	ErrInvalidFoodInit   = errors.New("invalid food init")
	ErrInvalidPreference = errors.New("invalid preference")
)

func main() {
	var amount, kPrefer uint16

	_, err := fmt.Scanln(&amount)
	if err != nil || amount == 0 || amount > maxAmount {
		fmt.Println(ErrInvalidFoodAmount)

		return
	}

	foodRating := make([]int, amount)

	for i := range amount {
		_, err = fmt.Scan(&foodRating[i])
		if err != nil || foodRating[i] < minRating || foodRating[i] > maxRating {
			fmt.Println(ErrInvalidFoodInit)

			return
		}
	}

	_, err = fmt.Scan(&kPrefer)
	if err != nil || kPrefer == 0 || kPrefer > amount {
		fmt.Println(ErrInvalidPreference)

		return
	}

	resultPrefer(kPrefer, foodRating)
}

func resultPrefer(pref uint16, ratings []int) {
	var result int

	foodHeap := intheap.InitIntHeap(ratings)

	for range pref {
		value, isGood := heap.Pop(foodHeap).(int)
		if isGood {
			result = value
		}
	}

	fmt.Println(result)
}
