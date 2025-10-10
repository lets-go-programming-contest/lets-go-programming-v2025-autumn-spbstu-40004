package main

import (
	"fmt"
)

func main() {
	var (
		a, b      int
		operation string
	)
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err := fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err3 := fmt.Scanln(&operation)
	if err3 != nil {
		fmt.Println("Invalid operation")
		return
	}
	var result int
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = a / b
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}
