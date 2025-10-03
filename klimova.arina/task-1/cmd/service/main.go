package main

import (
	"fmt"
)

func main() {
	var a, b int
	var operation string
	_, err1 := fmt.Scanln(&a)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err2 := fmt.Scanln(&b)
	if err2 != nil {
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
		fmt.Println("Invalid opertion")
		return
	}

	fmt.Println(result)
}
