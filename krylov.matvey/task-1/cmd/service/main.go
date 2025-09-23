package main

import "fmt"

func main() {
	var operand1 int
	_, err := fmt.Scanln(&operand1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var operand2 int
	_, err = fmt.Scanln(&operand2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operator string
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	var result int
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = operand1 / operand2
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(result)
}
