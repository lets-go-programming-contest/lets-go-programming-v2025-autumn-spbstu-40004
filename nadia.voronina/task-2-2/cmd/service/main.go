package main

import (
	"container/heap"
	"errors"
	"fmt"

	"spbstu.ru/nadia.voronina/task-2-2/pkg/intheap"
)

var (
	ErrInvalidNumberOfDishes = errors.New("invalid number of dishes")
	ErrInvalidDish           = errors.New("invalid dish")
	ErrInvalidPreferredDish  = errors.New("invalid preferred dish")
)

func main() {
	var numberOfDishes int

	_, errN := fmt.Scanln(&numberOfDishes)
	if errN != nil {
		fmt.Println(ErrInvalidNumberOfDishes)

		return
	}

	var dishes intheap.IntHeap

	heap.Init(&dishes)

	var dish int

	for range numberOfDishes {
		_, errK := fmt.Scan(&dish)

		if errK != nil {
			fmt.Println(ErrInvalidDish)

			return
		}

		heap.Push(&dishes, dish)
	}

	var preferredDish int

	_, errK := fmt.Scan(&preferredDish)

	if errK != nil {
		fmt.Println(ErrInvalidPreferredDish)

		return
	}

	for range numberOfDishes - preferredDish {
		heap.Pop(&dishes)
	}

	fmt.Printf("%d\n", heap.Pop(&dishes))
}
