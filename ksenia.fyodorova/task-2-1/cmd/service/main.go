package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minTemp          = 15
	maxTemp          = 30
	minPartsLength   = 2
	greaterEqualSign = ">="
	lessEqualSign    = "<="
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numOfDepartments, _ := strconv.Atoi(scanner.Text())

	for range numOfDepartments {
		scanner.Scan()
		numOfStaff, _ := strconv.Atoi(scanner.Text())

		start := minTemp
		end := maxTemp

		for range numOfStaff {
			scanner.Scan()
			preference := scanner.Text()

			parts := strings.Fields(preference)
			if len(parts) < minPartsLength {
				fmt.Println(-1)

				continue
			}

			sign := parts[0]
			temperature, err := strconv.Atoi(parts[1])

			if err != nil {
				fmt.Println(-1)

				continue
			}

			switch sign {
			case greaterEqualSign:
				if temperature > start {
					start = temperature
				}
			case lessEqualSign:
				if temperature < end {
					end = temperature
				}
			default:
				fmt.Println(-1)

				continue
			}

			if start > end {
				fmt.Println(-1)
			} else {
				fmt.Println(start)
			}
		}
	}
}
