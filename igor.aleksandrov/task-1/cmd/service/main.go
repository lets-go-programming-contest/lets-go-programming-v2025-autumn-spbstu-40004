package main

import "fmt"

func main() {
	var lhs int
	_, err := fmt.Scanln(&lhs)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var rhs int
	_, err = fmt.Scanln(&rhs)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operation string
	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(lhs + rhs)
	case "-":
		fmt.Println(lhs - rhs)
	case "*":
		fmt.Println(lhs * rhs)
	default:
		fmt.Println("Unsupported operation")
	}
}
