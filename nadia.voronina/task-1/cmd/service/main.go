package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var firstOperand int
	_, errFirst := fmt.Scan(&firstOperand)
	if errFirst != nil {
		fmt.Println("Invalid first operand")
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}

	var secondOperand int
	_, errSecond := fmt.Scan(&secondOperand)
	if errSecond != nil {
		fmt.Println("Invalid second operand")
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}

	var operator string
	_, _ = fmt.Scan(&operator)

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
		fmt.Println("Invalid operator")
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}
}
