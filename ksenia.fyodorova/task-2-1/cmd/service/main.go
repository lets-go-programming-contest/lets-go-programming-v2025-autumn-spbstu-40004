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
	numOfDepartments, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < numOfDepartments; i++ {
		scanner.Scan()
		numOfStaff, _ := strconv.Atoi(scanner.Text())

		start := 15
		end := 30

		for j := 0; j < numOfStaff; j++ {
			scanner.Scan()
			preference := scanner.Text()

			parts := strings.Fields(preference)
			if len(parts) < 2 {
				fmt.Println(-1)
				continue
			}

			sign := parts[0]
			temperature, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println(-1)
				continue
			}

			if sign == ">=" {
				if temperature > start {
					start = temperature
				}
			} else if sign == "<=" {
				if temperature < end {
					end = temperature
				}
			} else {
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
