package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		numSt    int
		numNd    int
		operator string
	)

	_, err := fmt.Fscanln(os.Stdin, &numSt)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Fscanln(os.Stdin, &numNd)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Fscanln(os.Stdin, &operator)
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
		if numSt == 0 {
			fmt.Println("Division by zero")
		}

		fmt.Println(numSt / numNd)
	default:
		fmt.Println("Invalid operation")
	}
}
