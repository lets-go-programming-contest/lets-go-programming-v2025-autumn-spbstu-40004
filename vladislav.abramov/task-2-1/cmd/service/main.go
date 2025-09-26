package main

import (
	"fmt"
)

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
)

func processDepartment(kCount int) {
	minTemp := minAllowedTemp
	maxTemp := maxAllowedTemp
	valid := true

	for range kCount {
		var (
			operation string
			temp      int
		)

		_, err := fmt.Scan(&operation, &temp)
		if err != nil {
			fmt.Print("Invalid data\n")

			continue
		}

		if temp > maxAllowedTemp || temp < minAllowedTemp {
			fmt.Print("Invalid data\n")

			continue
		}

		if !valid {
			fmt.Print("-1\n")

			continue
		}

		switch operation {
		case ">=":
			if temp > minTemp {
				minTemp = temp
			}
		case "<=":
			if temp < maxTemp {
				maxTemp = temp
			}
		default:
			fmt.Print("Invalid data\n")
		}

		if minTemp > maxTemp {
			valid = false

			fmt.Print("-1\n")
		} else {
			fmt.Println(minTemp)
		}
	}
}

func main() {
	var (
		nCount int
		err    error
	)

	_, err = fmt.Scanln(&nCount)
	if err != nil || nCount < 1 || nCount > 1000 {
		fmt.Print("Invalid var N\n")

		return
	}

	for range nCount {
		var kCount int

		_, err = fmt.Scanln(&kCount)
		if err != nil || kCount < 1 || kCount > 1000 {
			fmt.Print("Invalid var K\n")

			continue
		}

		processDepartment(kCount)
	}
}
