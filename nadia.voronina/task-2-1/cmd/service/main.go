package main

import (
	"fmt"
)

func processLessEq(maxDegree *int, minDegree *int, degree int) {
	switch {
	case (*maxDegree >= degree) && (*minDegree <= degree):
		*maxDegree = degree

		fmt.Println(*minDegree)
	case (*maxDegree <= degree) && (*minDegree <= degree):
		fmt.Println(*minDegree)
	default:
		*maxDegree = 0
		*minDegree = 0

		fmt.Println(-1)
	}
}

func processGreaterEq(maxDegree *int, minDegree *int, degree int) {
	if (*minDegree <= degree) && (*maxDegree >= degree) {
		*minDegree = degree

		fmt.Println(*minDegree)
	} else {
		*maxDegree = 0
		*minDegree = 0

		fmt.Println(-1)
	}
}

func processCondition(maxDegree *int, minDegree *int, degree int, sign string) {
	if *minDegree == 0 && *maxDegree == 0 {
		fmt.Println(-1)

		return
	}

	switch sign {
	case "<=":
		processLessEq(maxDegree, minDegree, degree)
	case ">=":
		processGreaterEq(maxDegree, minDegree, degree)
	default:
		fmt.Println("Wrong sign has been added")

		return
	}
}

func main() {
	var numberOfDepartments, numberOfEmployees int

	_, errN := fmt.Scanln(&numberOfDepartments)
	if errN != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range numberOfDepartments {
		_, errK := fmt.Scanln(&numberOfEmployees)
		if errK != nil {
			fmt.Println("Invalid number of employees")

			return
		}

		maxDegree := 31
		minDegree := 14

		for range numberOfEmployees {
			var sign string

			var degree int

			_, err1 := fmt.Scan(&sign)
			if err1 != nil {
				fmt.Println("Invalid sign")

				return
			}

			_, err2 := fmt.Scanln(&degree)
			if err2 != nil {
				fmt.Println("Invalid degree")

				return
			}

			processCondition(&maxDegree, &minDegree, degree, sign)
		}
	}
}
