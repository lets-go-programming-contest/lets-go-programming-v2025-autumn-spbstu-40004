package main

import "fmt"

func main() {
	var a, b int
	var op string

	scan := func(dest interface{}, msg string) bool {
		if _, err := fmt.Scan(dest); err != nil {
			fmt.Println(msg)
			return false
		}
		return true
	}

	if scan(&a, "Invalid first operand") &&
		scan(&b, "Invalid second operand") &&
		scan(&op, "Invalid operation") {
		switch op {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			if b == 0 {
				fmt.Println("Division by zero")
			} else {
				fmt.Println(a / b)
			}
		default:
			fmt.Println("Invalid operation")
		}
	}
}
