package main

import (
	"fmt"
)

func main() {
	var (
		dishesAmount int
	)

	_, err := fmt.Scan(&dishesAmount)
	if err != nil {
		fmt.Println("invalid number of dishes")

		return
	}
}
