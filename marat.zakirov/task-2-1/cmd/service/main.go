package main

import "fmt"

func main() {
	var (
		departNum int
		employNum int
		newTemp   int
		operator  string
	)

	leftBorder, rightBorder := 15, 30

	_, err := fmt.Scan(&departNum)
	if err != nil {
		fmt.Println("ERROR in getting the number of departments")

		return
	}

	for range departNum {
		_, err = fmt.Scan(&employNum)
		if err != nil {
			fmt.Println("ERROR in getting the number of employees")

			return
		}

		for range employNum {
			readNum, err := fmt.Scan(&operator, &newTemp)
			if err != nil {
				if readNum == 0 {
					fmt.Println("ERROR in getting operator")

					return
				} else {
					fmt.Println("ERROR in getting new temperature")

					return
				}
			}

			if rightBorder == -1 {
				continue
			}

			switch operator {
			case "<=":
				rightBorder = newTemp
			case ">=":
				leftBorder = newTemp
			}

			if rightBorder < leftBorder {
				rightBorder = -1
				leftBorder = -1
				fmt.Println(leftBorder)

				continue
			}

			fmt.Println(leftBorder)
		}
	}
}
