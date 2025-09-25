package main

import "fmt"

func main() {
	var departmentNum int
	_, err := fmt.Scan(&departmentNum)
	if err != nil {
		fmt.Println("Error: invalid department number")
		return
	}

	for iDepartment := 0; iDepartment != departmentNum; iDepartment++ {
		temperatures := [2]int{15, 30}
		var employeeNum int
		_, err = fmt.Scan(&employeeNum)
		if err != nil {
			fmt.Println("Error: invalid employee number")
			return
		}

		for iEmployee := 0; iEmployee != employeeNum; iEmployee++ {
			var cmpOperator string
			var border int
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
				temperatures[0] = border
			case "<=":
				if border < temperatures[0] {
					fmt.Println(-1)
					continue
				}
				temperatures[1] = border
			default:
				fmt.Println("Error: invalid compare operator")
				return
			}
		}
	}
}
