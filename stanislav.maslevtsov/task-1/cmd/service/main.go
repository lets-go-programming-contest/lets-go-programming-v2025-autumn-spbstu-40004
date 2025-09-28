package main

import (
	"errors"
	"fmt"
)

var ErrDivisionByZero = errors.New("Division by zero")

func safeDivision(operand1 int, operand2 int) (int, error) {
	if operand2 == 0 {
		return 0, ErrDivisionByZero
	}
	return operand1 / operand2, nil
}

func main() {
	var leftOperand int
	_, err := fmt.Scanln(&leftOperand)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var rightOperand int
	_, err = fmt.Scanln(&rightOperand)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operation string
	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(leftOperand + rightOperand)
	case "-":
		fmt.Println(leftOperand - rightOperand)
	case "*":
		fmt.Println(leftOperand * rightOperand)
	case "/":
		var divisionResult int
		divisionResult, err = safeDivision(leftOperand, rightOperand)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(divisionResult)
	default:
		fmt.Println("Invalid operation")
	}
}
