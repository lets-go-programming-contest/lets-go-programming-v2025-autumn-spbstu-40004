package main

import "fmt"

func main() {
	var operand1, operand2 int
	var operator string
	cnt, err := fmt.Scan(&operand1, &operand2, &operator)
	if err != nil {
		switch cnt {
		case 0:
			fmt.Println("Invalid first operand")
		case 1:
			fmt.Println("Invalid second operand")
		}
		return
	}
	switch operator {
	case "+":
		fmt.Print(operand1 + operand2)
	case "-":
		fmt.Print(operand1 - operand2)
	case "*":
		fmt.Print(operand1 * operand2)
	case "/":
		if operand2 == 0 {
			fmt.Print("Division by zero")
		} else {
			fmt.Print(operand1 / operand2)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
