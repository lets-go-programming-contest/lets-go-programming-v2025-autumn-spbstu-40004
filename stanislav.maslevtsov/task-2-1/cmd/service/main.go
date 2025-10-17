package main

import (
	"fmt"

	"github.com/jambii1/task-2-1/internal/tempupdater"
)

func main() {
	var (
		departmentNum uint
		employeeNum   uint
		cmpOperator   string
		newBorder     int
	)

	_, err := fmt.Scan(&departmentNum)
	if err != nil {
		fmt.Println(err)

		return
	}

	for range departmentNum {
		tempUpd := tempupdater.NewTempUpdater()

		_, err = fmt.Scan(&employeeNum)
		if err != nil {
			fmt.Println(err)

			return
		}

		for range employeeNum {
			_, err = fmt.Scan(&cmpOperator, &newBorder)
			if err != nil {
				fmt.Println(err)

				return
			}

			err := tempUpd.Update(cmpOperator, newBorder)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(tempUpd.GetCurrentTemp())
			}
		}
	}
}
