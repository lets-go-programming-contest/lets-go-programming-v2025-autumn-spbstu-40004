package main

import (
	"fmt"
)

func processCondition(maxDegree *int, minDegree *int, degree int, sign string) {
	switch sign {
	case "<=":
		switch {
		case (*maxDegree >= degree) && (*minDegree <= degree):
			*maxDegree = degree

			fmt.Println(*minDegree)
		case (*maxDegree <= degree) && (*minDegree <= degree):
			fmt.Println(*minDegree)
		default:
			fmt.Println(-1)
		}
	case ">=":
		if (*minDegree <= degree) && (*maxDegree >= degree) {
			*minDegree = degree

			fmt.Println(*minDegree)
		} else {
			fmt.Println(-1)
		}
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
