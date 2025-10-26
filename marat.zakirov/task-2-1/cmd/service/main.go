package main

import (
	"fmt"
)

func main() {
	var (
		departNum, employNum int
		newTemp              int
		operator             string
	)

	const (
		leftBorder  = 15
		rightBorder = 30
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

		currLeftBorder, currRightBorder := leftBorder, rightBorder

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

			if currLeftBorder == -1 {
				fmt.Println(currLeftBorder)

				continue
			}

			temperatureHandling(operator, &currLeftBorder, &currRightBorder, newTemp)

			if currRightBorder < currLeftBorder {
				currLeftBorder, currRightBorder = -1, -1
				fmt.Println(currLeftBorder)

				continue
			}

			fmt.Println(currLeftBorder)
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
