package main

import "fmt"

func main() {
	var (
		departmentNum uint
		employeeNum   uint
		cmpOperator   string
		border        uint
	)

	_, err := fmt.Scan(&departmentNum)
	if err != nil {
		fmt.Println("Error: invalid department number")

		return
	}

	for range departmentNum {
		temperature_borders := make([]uint, 2)
		temperature_borders[0] = 15
		temperature_borders[1] = 30

		_, err = fmt.Scan(&employeeNum)
		if err != nil {
			fmt.Println("Error: invalid employee number")

			return
		}

		for range employeeNum {
			_, err = fmt.Scan(&cmpOperator, &border)
			if err != nil {
				fmt.Println("Error: invalid temperature border")

				return
			}

			switch cmpOperator {
			case ">=":
				if border > temperature_borders[1] {
					fmt.Println(-1)

					continue
				}
				if border > temperature_borders[0] {
					temperature_borders[0] = border
				}
			case "<=":
				if border < temperature_borders[0] {
					fmt.Println(-1)

					continue
				}
				if border < temperature_borders[1] {
					temperature_borders[1] = border
				}
				temperature_borders[1] = border
			default:
				fmt.Println("Error: invalid compare operator")

				return
			}

			fmt.Println(temperature_borders[0])
		}
	}
}
