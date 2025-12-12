package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lolnyok/task-2-1/temperature"
)

const (
	minPartsLength = 2
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	numOfDepartments, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(-1)

		return
	}

	for range numOfDepartments {
		scanner.Scan()

		numOfStaff, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(-1)

			continue
		}

		tempRange := temperature.NewTemperatureRange()

		for range numOfStaff {
			scanner.Scan()
			preference := scanner.Text()

			parts := strings.Fields(preference)
			if len(parts) < minPartsLength {
				fmt.Println(-1)

				continue
			}

			sign := parts[0]

			temp, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println(-1)

				continue
			}

			err = tempRange.Update(sign, temp)
			if err != nil {
				fmt.Println(-1)

				continue
			}

			comfortTemp, err := tempRange.GetComfortableTemp()
			if err != nil {
				fmt.Println(-1)
			} else {
				fmt.Println(comfortTemp)
			}
		}
	}
}
