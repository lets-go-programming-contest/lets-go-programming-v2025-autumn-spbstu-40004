package main

import ("fmt")

func main() {
	var a, b int

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var ch string

	_, err = fmt.Scan(&ch)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch ch {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			if b == 0 {
				fmt.Println("Division by zero")
				return
			}
			fmt.Println(a / b)
		default:
			fmt.Println("Invalid operation")
	}
}
