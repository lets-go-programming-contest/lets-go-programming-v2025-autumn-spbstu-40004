package main

import "fmt"

func main() {
	var (
		operandFirst, operandSecond int
		sign                        string
	)

	_, err := fmt.Scanln(&operandFirst)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&operandSecond)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&sign)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch sign {
	case "+":
		fmt.Println(operandFirst + operandSecond)
	case "-":
		fmt.Println(operandFirst - operandSecond)
	case "*":
		fmt.Println(operandFirst * operandSecond)
	case "/":
		if operandSecond == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(operandFirst / operandSecond)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
