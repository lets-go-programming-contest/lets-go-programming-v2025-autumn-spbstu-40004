package main

import "fmt"

func main() {
	var firstNum, secondNum, result int
	var operator string

	fmt.Println("Type the first number: ")
	_, errFirstNum := fmt.Scanln(&firstNum)

	fmt.Println("Type the second number: ")
	_, errSecondNum := fmt.Scanln(&secondNum)

	fmt.Println("Type an operator (+, -, *, /): ")
	_, errOperator := fmt.Scanln(&operator)

	if errFirstNum != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if errSecondNum != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if errOperator != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operator {
	case "+":
		result = firstNum + secondNum
	case "-":
		result = firstNum - secondNum
	case "/":
		if secondNum == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = firstNum / secondNum
	case "*":
		result = firstNum * secondNum
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}
