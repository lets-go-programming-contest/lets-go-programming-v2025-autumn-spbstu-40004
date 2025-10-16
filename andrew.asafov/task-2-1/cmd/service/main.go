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

	scanner.Scan()
	groupCount, _ := strconv.Atoi(scanner.Text())

	results := make([]int, 0, groupCount)

	for range groupCount {
		scanner.Scan()
		expressionCountText := scanner.Text()
		expressionCount, err := strconv.Atoi(expressionCountText)

		if err != nil {
			results = append(results, -1)
			continue
		}

		minValue := 15
		maxValue := 30
		valid := true

		for range expressionCount {
			if !scanner.Scan() {
				break
			}

			expression := scanner.Text()
			parts := strings.Fields(expression)

			const minPartsCount = 2
			if len(parts) < minPartsCount {
				continue
			}

			operator := parts[0]
			value, parseErr := strconv.Atoi(parts[1])

			if parseErr != nil {
				continue
			}

			if operator == ">=" {
				if value > minValue {
					minValue = value
				}
			} else if operator == "<=" {
				if value < maxValue {
					maxValue = value
				}
			}

			if minValue > maxValue {
				valid = false
			}
		}

		if valid {
			results = append(results, minValue)
		} else {
			results = append(results, -1)
		}
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
