package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processGroup(scanner *bufio.Scanner, expressionCount int) []int {
	minValue := 15
	maxValue := 30
	results := make([]int, 0, expressionCount)
	valid := true

	for range expressionCount {
		if !scanner.Scan() {
			break
		}

		expression := scanner.Text()
		parts := strings.Fields(expression)

		if !valid {
			results = append(results, -1)

			continue
		}

		const minPartsCount = 2
		if len(parts) < minPartsCount {
			results = append(results, minValue)

			continue
		}

		operator := parts[0]
		value, parseErr := strconv.Atoi(parts[1])

		if parseErr != nil {
			results = append(results, minValue)

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

			results = append(results, -1)
		} else {
			results = append(results, minValue)
		}
	}

	return results
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	groupCount, _ := strconv.Atoi(scanner.Text())

	for range groupCount {
		scanner.Scan()
		expressionCountText := scanner.Text()

		expressionCount, err := strconv.Atoi(expressionCountText)
		if err != nil {
			fmt.Println(-1)

			continue
		}

		results := processGroup(scanner, expressionCount)
		for _, result := range results {
			fmt.Println(result)
		}
	}
}
