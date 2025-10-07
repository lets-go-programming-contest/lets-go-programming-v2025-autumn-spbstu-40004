package main

import "fmt"

func main() {
	var (
		departNum, employNum    int
		leftBorder, rightBorder int
		newTemp                 int
		operator                string
	)

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

		leftBorder, rightBorder = 15, 30

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

			if leftBorder == -1 {
				fmt.Println(leftBorder)

				continue
			}

			temperatureHandling(operator, &leftBorder, &rightBorder, newTemp)

			if rightBorder < leftBorder {
				leftBorder, rightBorder = -1, -1
				fmt.Println(leftBorder)

				continue
			}

			fmt.Println(leftBorder)
		}
	}
}

func temperatureHandling(op string, lBorder *int, rBorder *int, newValue int) {
	switch op {
	case "<=":
		if *rBorder > newValue {
			*rBorder = newValue
		}

	case ">=":
		if *lBorder < newValue {
			*lBorder = newValue
		}
	}
}
