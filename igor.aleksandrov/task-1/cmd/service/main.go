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

	fmt.Println(lhs + rhs)
}
