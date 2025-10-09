package main

import "fmt"

func main() {
	var departmentsCount int
	_, err := fmt.Scanln(&departmentsCount)
	if err != nil || departmentsCount < 0 {
		fmt.Println("Invalid departments count!")
		return
	}

	var employeesCount int
	_, err = fmt.Scanln(&employeesCount)
	if err != nil || employeesCount < 0 {
		fmt.Println("Invalid employees count!")
		return
	}
}
