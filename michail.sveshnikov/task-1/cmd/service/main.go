package main

import "fmt"

func main() {
	var operand1, operand2 int
	_, err := fmt.Scan(&operand1)
	if err != nil {
		fmt.Println("Invalid first operand")
	}
	_, err = fmt.Scan(&operand2)
	if err != nil {
		fmt.Println("Invalid second operand")
	}
	var operator string
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
	}
}
