package main

import "fmt"

func main() {
	var departmentsCount int
	_, err := fmt.Scanln(&departmentsCount)
	if err != nil || departmentsCount < 1 || departmentsCount > 1000 {
		fmt.Println("Invalid departments count!")
		return
	}

	for range departmentsCount {
		var employeesCount int
		_, err = fmt.Scanln(&employeesCount)
		if err != nil || employeesCount < 1 || employeesCount > 1000 {
			fmt.Println("Invalid employees count!")
			return
		}
	}
}
