package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("task-1")
}

func calculate(num_1 int, num_2 int, operator byte) (int, error) {
	switch operator {
	case '+':
		return num_1 + num_2, nil
	case '-':
		return num_1 - num_2, nil
	case '*':
		return num_1 * num_2, nil
	case '/':
		if num_2 == 0 {
			return 0, errors.New("division by zero")
		}

		return num_1 / num_2, nil
	default:
		return 0, errors.New("invalid operator")
	}

}
