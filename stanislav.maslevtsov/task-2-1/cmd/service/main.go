package main

import "fmt"

func main() {
	var (
		departmentNum uint
		employeeNum   uint
		cmpOperator   string
		leftBorder    uint
		rightBorder   uint
		newBorder     uint
	)

	_, err := fmt.Scan(&departmentNum)
	if err != nil {
		fmt.Println("Error: invalid department number")

		return
	}

	for range departmentNum {
		leftBorder = 15
		rightBorder = 30

		_, err = fmt.Scan(&employeeNum)
		if err != nil {
			fmt.Println("Error: invalid employee number")

			return
		}

		for range employeeNum {
			_, err = fmt.Scan(&cmpOperator, &newBorder)
			if err != nil {
				fmt.Println("Error: invalid temperature border")

				return
			}

			switch cmpOperator {
			case ">=":
				if newBorder > rightBorder {
					fmt.Println(-1)

					continue
				} else if newBorder > leftBorder {
					leftBorder = newBorder
				}
			case "<=":
				if newBorder < leftBorder {
					fmt.Println(-1)

					continue
				} else if newBorder < rightBorder {
					rightBorder = newBorder
				}
			default:
				fmt.Println("Error: invalid compare operator")

				return
			}

			fmt.Println(leftBorder)
		}
	}
}
