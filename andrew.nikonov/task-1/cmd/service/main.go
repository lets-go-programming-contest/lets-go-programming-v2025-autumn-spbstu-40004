package main

import "fmt"

func main() {
	var (
		operand1 int
		operand2 int
		operator string
	)
	_, err := fmt.Scanln(&operand1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scanln(&operand2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid input for operator")
		return
	}
	switch operator {
	case "+":
		fmt.Println(operand1 + operand2)
	case "-":
		fmt.Println(operand1 - operand2)
	case "*":
		fmt.Println(operand1 * operand2)
	case "/":
		if operand2 == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(operand1 / operand2)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
