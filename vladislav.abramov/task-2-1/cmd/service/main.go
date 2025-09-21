package main

import (
	"fmt"
)

func main() {
	var (
		nCount int
		err    error
	)

	_, err = fmt.Scanln(&nCount)
	if err != nil {
		fmt.Print("Invalid var N\n")

		return
	}

	for range nCount {
		var kCount int
		_, err = fmt.Scanln(&kCount)
    if err != nil {
      fmt.Print("Invalid var K\n")
    }

		minTemp := 15
		maxTemp := 30
		valid := true

		for range kCount {
			var (
				operation string
				temp      int
			)

      _, err = fmt.Scan(&operation, &temp)
      if err != nil {
        fmt.Print("Invalid data\n")
      }

			if temp > 30 || temp < 15 {
				fmt.Print("Invalid temperature\n")
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
}
