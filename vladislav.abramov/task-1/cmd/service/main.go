package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var op string
	var err error

	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Print("Invalid first operand\n")
		return
	}

	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Print("Invalid second operand\n")
		return
	}

	_, err = fmt.Scanln(&op)
	if err != nil {
		fmt.Print("Invalid operation\n")
		return
	}

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Print("Division by zero\n")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Print("Invalid operation\n")
	}
}
