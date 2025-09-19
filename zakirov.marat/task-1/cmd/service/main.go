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
		operator byte
		result   int
	)

	n, err := fmt.Fscan(os.Stdin, &num_1, &num_2, &operator)
	if err != nil {
		switch n {
		case 0:
			fmt.Println("Invalid first operand")
		case 1:
			fmt.Println("Invalid second operand")
		}

		return
	}
	result, err = calculate(num_1, num_2, operator)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func calculate(operand_1 int, operand_2 int, operator byte) (int, error) {
	switch operator {
	case '+':
		return operand_1 + operand_2, nil
	case '-':
		return operand_1 - operand_2, nil
	case '*':
		return operand_1 * operand_2, nil
	case '/':
		if operand_2 == 0 {
			return 0, errors.New("Division by zero")
		}

		return operand_1 / operand_2, nil
	default:
		return 0, errors.New("Invalid operation")
	}

}
