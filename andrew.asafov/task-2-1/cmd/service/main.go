package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processGroup(scanner *bufio.Scanner, expressionCount int) int {
	minValue := 15
	maxValue := 30

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
			return -1
		}
	}

	return minValue
}

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

		result := processGroup(scanner, expressionCount)
		results = append(results, result)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
