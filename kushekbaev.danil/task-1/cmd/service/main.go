package main

import "fmt"

func main() {
	var (
		firstNumber  int
		secondNumber int
		result       int
		operator     string
	)

	_, err := fmt.Scanln(&firstNumber)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&secondNumber)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operator {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "*":
		result = firstNumber * secondNumber
	case "/":
		if secondNumber == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = firstNumber / secondNumber
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(result)
}
