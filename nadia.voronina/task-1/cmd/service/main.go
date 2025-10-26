package main

import (
	"fmt"
)

func main() {
	var firstOperand int
	_, errFirst := fmt.Scanln(&firstOperand)
	if errFirst != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var secondOperand int
	_, errSecond := fmt.Scanln(&secondOperand)
	if errSecond != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operator string
	_, _ = fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Println(firstOperand + secondOperand)
	case "-":
		fmt.Println(firstOperand - secondOperand)
	case "*":
		fmt.Println(firstOperand * secondOperand)
	case "/":
		if secondOperand == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(firstOperand / secondOperand)
		}
	default:
		fmt.Println("Invalid operation")
		return
	}
}
