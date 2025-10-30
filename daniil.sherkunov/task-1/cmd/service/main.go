package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(r *bufio.Reader) (string, bool) {
	s, err := r.ReadString('\n')
	if err != nil && len(s) == 0 {
		return "", false
	}
	return strings.TrimSpace(s), true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	firstLine, ok := readLine(in)
	if !ok {
		fmt.Println("Invalid first operand")
		return
	}
	a, err := strconv.Atoi(firstLine)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	secondLine, ok := readLine(in)
	if !ok {
		fmt.Println("Invalid second operand")
		return
	}
	b, err := strconv.Atoi(secondLine)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	opLine, ok := readLine(in)
	if !ok || len(opLine) == 0 {
		fmt.Println("Invalid operation")
		return
	}
	op := strings.TrimSpace(opLine)

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
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operation")
	}
}
