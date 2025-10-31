package main

import (
	"fmt"

	"github.com/cherkasoov/lets-go-programming-v2025-autumn-spbstu-40004/task-2-1/internal/temperature"
)

func handleDepartmentRequests(employeeCount int, conditions []string, targetTemps []int) error {
    tr := temperature.NewTemperatureRange(temperature.MinTemperature, temperature.MaxTemperature)

    for i := 0; i < employeeCount; i++ {
        err := tr.Update(targetTemps[i], conditions[i])
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

    var departmentCount int

    _, err := fmt.Scanln(&departmentCount)
    if err != nil || departmentCount < minDepartmentCount || departmentCount > maxDepartmentCount {
        fmt.Println("Invalid department count")
        return
    }

    for i := 0; i < departmentCount; i++ {
        var employeeCount int

        _, err = fmt.Scanln(&employeeCount)
        if err != nil || employeeCount < minEmployeeCount || employeeCount > maxEmployeeCount {
            fmt.Println("Invalid employee count")
            return
        }

        conditions := make([]string, employeeCount)
        targetTemps := make([]int, employeeCount)

        for j := 0; j < employeeCount; j++ {
            _, err := fmt.Scanln(&conditions[j], &targetTemps[j])
            if err != nil || targetTemps[j] < temperature.MinTemperature || targetTemps[j] > temperature.MaxTemperature {
                fmt.Println("Invalid employee input")
                return
            }
        }

        err := handleDepartmentRequests(employeeCount, conditions, targetTemps)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}
