package main

import "fmt"

func main() {
	var (
		departNum int
		employNum int
		currTemp  int
		newTemp   int
		operator  string
	)

	fmt.Scan(&departNum)

	for dep := 0; dep < departNum; dep++ {
		fmt.Scan(&employNum)
		currTemp = 0
		for empl := 0; empl < employNum; empl++ {
			fmt.Scan(&operator)
			fmt.Scan(&newTemp)
		}
	}
}
