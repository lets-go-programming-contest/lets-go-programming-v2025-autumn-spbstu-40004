package main

import (
	"errors"
	"fmt"
)

func safeDivision(operand1 int, operand2 int) (int, error) {
	if operand2 == 0 {
		return 0, errors.New("division by zero")
	}
	return operand1 / operand2, nil
}

func main() {
	var operand1 int
	_, scanErr := fmt.Scanln(&operand1)
	if scanErr != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var operand2 int
	_, scanErr = fmt.Scanln(&operand2)
	if scanErr != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operation string
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Println(operand1 + operand2)
	case "-":
		fmt.Println(operand1 - operand2)
	case "*":
		fmt.Println(operand1 * operand2)
	case "/":
		divisionResult, divisionErr := safeDivision(operand1, operand2)
		if divisionErr != nil {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(divisionResult)
	default:
		fmt.Println("Invalid operation")
	}
}
