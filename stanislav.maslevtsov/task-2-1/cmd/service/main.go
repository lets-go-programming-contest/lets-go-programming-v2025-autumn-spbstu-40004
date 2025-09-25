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
		temperatures := [2]uint{15, 30}
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
				if border > temperatures[1] {
					fmt.Println(-1)

					continue
				}
				if border > temperatures[0] {
					temperatures[0] = border
				}
			case "<=":
				if border < temperatures[0] {
					fmt.Println(-1)

					continue
				}
				if border < temperatures[1] {
					temperatures[1] = border
				}
				temperatures[1] = border
			default:
				fmt.Println("Error: invalid compare operator")

				return
			}

			fmt.Println(temperatures[0])
		}
	}
}
