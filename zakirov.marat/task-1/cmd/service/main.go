package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var (
		num_1    int
		num_2    int
		operator string
		result   int
	)

	_, err := fmt.Fscanln(os.Stdin, &num_1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Fscanln(os.Stdin, &num_2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Fscanln(os.Stdin, &operator)
	if err != nil {
		fmt.Println("Invalid third operand")
		return
	}

	result, err = calculate(num_1, num_2, operator)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func calculate(operand_1 int, operand_2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return operand_1 + operand_2, nil
	case "-":
		return operand_1 - operand_2, nil
	case "*":
		return operand_1 * operand_2, nil
	case "/":
		if operand_2 == 0 {
			return 0, errors.New("Division by zero")
		}

		return operand_1 / operand_2, nil
	default:
		return 0, errors.New("Invalid operation")
	}
}
