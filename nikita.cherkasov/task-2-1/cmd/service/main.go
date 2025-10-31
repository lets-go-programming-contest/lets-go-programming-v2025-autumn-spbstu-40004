package main

import (
	"fmt"

	"github.com/cherkasoov/lets-go-programming-v2025-autumn-spbstu-40004/task-2-1/internal/temperature"
)

const (
	minDepartmentCount = 1
	maxDepartmentCount = 1000
	minEmployeeCount   = 1
	maxEmployeeCount   = 1000
)

func handleDepartmentRequests(employeeCount int, conditions []string, targetTemps []int) error {
	temp := temperature.NewTemperatureRange(temperature.MinTemperature, temperature.MaxTemperature)

	for i := range employeeCount {
		err := temp.Update(targetTemps[i], conditions[i])
		if err != nil {
			return fmt.Errorf("update temperature: %w", err)
		}

		fmt.Println(temp.GetMin())
	}

	return nil
}

func readEmployeeInput(employeeCount int) ([]string, []int, error) {
	conditions := make([]string, employeeCount)
	targetTemps := make([]int, employeeCount)

	for j := range employeeCount {
		_, err := fmt.Scanln(&conditions[j], &targetTemps[j])
		if err != nil || targetTemps[j] < temperature.MinTemperature ||
			targetTemps[j] > temperature.MaxTemperature {
			return nil, nil, fmt.Errorf("invalid employee input: %w", err)
		}
	}

	return conditions, targetTemps, nil
}

func processDepartment() error {
	var employeeCount int

	_, err := fmt.Scanln(&employeeCount)
	if err != nil || employeeCount < minEmployeeCount || employeeCount > maxEmployeeCount {
		return fmt.Errorf("invalid employee count: %w", err)
	}

	conditions, targetTemps, err := readEmployeeInput(employeeCount)
	if err != nil {
		return fmt.Errorf("read employee input: %w", err)
	}

	err = handleDepartmentRequests(employeeCount, conditions, targetTemps)
	if err != nil {
		return fmt.Errorf("handle department requests: %w", err)
	}

	return nil
}

func run() error {
	var departmentCount int

	_, err := fmt.Scanln(&departmentCount)
	if err != nil || departmentCount < minDepartmentCount || departmentCount > maxDepartmentCount {
		return fmt.Errorf("invalid department count: %w", err)
	}

	for range departmentCount {
		err := processDepartment()
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
