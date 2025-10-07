package main

import "fmt"

func main() {
	var (
		departNum int
		employNum int
		newTemp   int
		operator  string
	)

	leftBorder, rightBorder := 0, 30

	fmt.Scan(&departNum)
	for dep := 0; dep < departNum; dep++ {
		fmt.Scan(&employNum)
		for empl := 0; empl < employNum; empl++ {
			fmt.Scan(&operator)
			fmt.Scan(&newTemp)
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
				fmt.Println("-1")
				continue
			}

			fmt.Println((rightBorder + leftBorder) / 2)
		}
	}
}
