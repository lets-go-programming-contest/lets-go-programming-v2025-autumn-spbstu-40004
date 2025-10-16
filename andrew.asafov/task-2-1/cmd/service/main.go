package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	A, _ := strconv.Atoi(scanner.Text())

	results := make([]int, 0, A)

	for i := 0; i < A; i++ {
		if !scanner.Scan() {
			break
		}
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			results = append(results, -1)
			continue
		}

		minVal := 15
		maxVal := 30
		valid := true

		for j := 0; j < n; j++ {
			if !scanner.Scan() {
				break
			}
			expression := scanner.Text()
			parts := strings.Fields(expression)

			if len(parts) < 2 {
				continue
			}

			operator := parts[0]
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}

			if operator == ">=" {
				if value > minVal {
					minVal = value
				}
			} else if operator == "<=" {
				if value < maxVal {
					maxVal = value
				}
			}

			if minVal > maxVal {
				valid = false
			}
		}

		if valid {
			results = append(results, minVal)
		} else {
			results = append(results, -1)
		}
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
