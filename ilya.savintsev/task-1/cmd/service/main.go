package main

import (
	"fmt"
	"os"
)

func main() {
	var x, y int
	var operator string
	_, err1 := fmt.Scanln(&x)
	_, err2 := fmt.Scanln(&y)
	_, err3 := fmt.Scanln(&operator)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		os.Exit(1)
	}
	if err2 != nil {
		fmt.Println("Invalid second operand")
		os.Exit(1)
	}
	if err3 != nil {
		fmt.Println("Invalid input for operator")
		os.Exit(1)
	}
	if operator == "/" && y == 0 {
		fmt.Println("Division by zero")
		os.Exit(1)
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
