package main

import (
	"fmt"
)

func main() {
	var x, y int
	var operator string
	_, err1 := fmt.Scanln(&x)
	_, err2 := fmt.Scanln(&y)
	_, err3 := fmt.Scanln(&operator)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if err3 != nil {
		fmt.Println("Invalid input for operator")
		return
	}
	if operator == "/" && y == 0 {
		fmt.Println("Division by zero")
		return
	}
	switch operator {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		fmt.Println(x / y)
	default:
		fmt.Println("Invalid operation")
	}
}
