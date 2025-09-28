package main

import (
	"errors"
	"fmt"
)

var ErrInvalidCmpOperator = errors.New("invalid compare operator")

func setBorders(leftBorder uint, rightBorder uint, cmpOperator string, newBorder uint) (uint, uint, error) {
	switch cmpOperator {
	case ">=":
		if newBorder > rightBorder {
			return 0, 0, nil
		} else if newBorder > leftBorder {
			leftBorder = newBorder
		}
	case "<=":
		if newBorder < leftBorder {
			return 0, 0, nil
		} else if newBorder < rightBorder {
			rightBorder = newBorder
		}
	default:
		return 0, 0, ErrInvalidCmpOperator
	}

	return leftBorder, rightBorder, nil
}

func main() {
	var (
		departmentNum uint
		employeeNum   uint
		leftBorder    uint
		rightBorder   uint
		cmpOperator   string
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

			leftBorder, rightBorder, err = setBorders(leftBorder, rightBorder, cmpOperator, newBorder)
			if err != nil {
				fmt.Println(err)

				return
			}

			if leftBorder == 0 {
				fmt.Println(-1)
			} else {
				fmt.Println(leftBorder)
			}
		}
	}
}
