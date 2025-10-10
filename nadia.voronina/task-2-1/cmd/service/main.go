package main

import (
	"fmt"
)

func main() {
	var n, k int

	_, errN := fmt.Scanln(&n)
	if errN != nil {
		fmt.Println("Invalid number of departments")
		return
	}

	for i := 0; i < n; i++ {

		_, errK := fmt.Scanln(&k)
		if errK != nil {
			fmt.Println("Invalid number of employees")
			return
		}

		var maxDegree int = 31
		var minDegree int = 14

		for j := 0; j < k; j++ {
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
			if sign == "<=" {
				if (maxDegree >= degree) && (minDegree <= degree) {
					maxDegree = degree
					fmt.Println(minDegree)
				} else if (maxDegree <= degree) && (minDegree <= degree) {
					fmt.Println(minDegree)
				} else {
					fmt.Println(-1)
				}
			} else if sign == ">=" {
				if (minDegree <= degree) && (maxDegree >= degree) {
					minDegree = degree
					fmt.Println(minDegree)
				} else {
					fmt.Println(-1)
				}
			} else {
				fmt.Println("Wrong sign has been added")
				return
			}
		}
	}
}
