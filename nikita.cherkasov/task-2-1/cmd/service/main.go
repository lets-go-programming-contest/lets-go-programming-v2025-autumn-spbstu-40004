package main

import (
	"fmt"

	"github.com/cherkasoov/task-2-1/internal/temperature"
)


func handleDepartmentRequests(employeeCount int) error {
	tr := temperature.NewTemperatureRange(temperature.MinTemperature, temperature.MaxTemperature)

	for i := 0; i < employeeCount; i++ {
		var condition string
		var targetTemp int

		_, err := fmt.Scanln(&condition, &targetTemp)
		if err != nil || targetTemp < temperature.MinTemperature || targetTemp > temperature.MaxTemperature {
			return temperature.ErrInvalidInput
		}

		err = tr.Update(targetTemp, condition)
		if err != nil {
			return err
		}

		fmt.Println(tr.GetMin())
	}

	return nil
}

func main() {
	const (
		minDepartmentCount = 1
		maxDepartmentCount = 1000
		minEmployeeCount   = 1
		maxEmployeeCount   = 1000
	)

	var departmentCount, employeeCount int

	_, err := fmt.Scanln(&departmentCount)
	if err != nil || departmentCount < minDepartmentCount || departmentCount > maxDepartmentCount {
		fmt.Println("Invalid department count")
		return
	}

	for i := 0; i < departmentCount; i++ {
		_, err = fmt.Scanln(&employeeCount)
		if err != nil || employeeCount < minEmployeeCount || employeeCount > maxEmployeeCount {
			fmt.Println("Invalid employee count")
			return
		}

		err := handleDepartmentRequests(employeeCount)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
