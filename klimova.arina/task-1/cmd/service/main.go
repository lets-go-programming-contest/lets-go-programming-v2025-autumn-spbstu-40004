package main

import (
	"fmt"
)

func main() {
	var a, b int
	var operation string
	_, err1 := fmt.Scanln(&a)
	_, err2 := fmt.Scanln(&b)
	_, err3 := fmt.Scanln(&operation)
	if err1 != nil {
		fmt.Println("Invalid first operand")
	}
	if err2 != nil {
		fmt.Println("Invalid second operand")
	}
	if err3 != nil {
		fmt.Println("Invalid operation")
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
		fmt.Println("Invalid opertion")
		return
	}

	fmt.Println(result)
}
