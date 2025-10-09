package main

import "fmt"

func main() {
	var firstNumber, secondNumber, res int
	var operand string
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
	_, err = fmt.Scanln(&operand)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}
	switch operand {
	case "+":
		res = firstNumber + secondNumber
	case "-":
		res = firstNumber - secondNumber
	case "/":
		if secondNumber == 0 {
			fmt.Println("Division by zero")
			return
		}
		res = firstNumber / secondNumber
	case "*":
		res = firstNumber * secondNumber
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(res)
}
