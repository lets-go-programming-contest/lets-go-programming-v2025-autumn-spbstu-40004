package main

import "fmt"

func main() {
	var dishesCount int

	_, err := fmt.Scanln(&dishesCount)
	if err != nil || dishesCount < 1 || dishesCount > 10000 {
		fmt.Println("Invalid dishes count!")

		return
	}
}
