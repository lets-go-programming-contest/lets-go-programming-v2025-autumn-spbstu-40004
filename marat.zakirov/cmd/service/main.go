package main

import (
	"fmt"
)

func main() {
	var (
		numSt    int
		numNd    int
		operator string
	)

	_, err := fmt.Scan(&numSt)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&numNd)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Invalid third operand")
		return
	}

	switch operator {
	case "+":
		fmt.Println(numSt + numNd)
	case "-":
		fmt.Println(numSt - numNd)
	case "*":
		fmt.Println(numSt * numNd)
	case "/":
		if numNd == 0 {
			fmt.Println("Division by zero")
			return
		}

		fmt.Println(numSt / numNd)
	default:
		fmt.Println("Invalid operation")
	}
}
