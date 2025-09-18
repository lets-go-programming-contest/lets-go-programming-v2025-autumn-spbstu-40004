package main

import "fmt"

func main() {
	var operand1 int
	_, scanErr := fmt.Scanln(&operand1)
	if scanErr != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var operand2 int
	_, scanErr = fmt.Scanln(&operand2)
	if scanErr != nil {
		fmt.Println("Invalid second operand")
		return
	}
}
