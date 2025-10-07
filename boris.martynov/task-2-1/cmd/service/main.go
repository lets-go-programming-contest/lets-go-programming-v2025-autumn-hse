package main

import (
	"fmt"

	"github.com/JingolBong/task-2-1/internal/conditioner"
)

func main() {
	var numberOfDepartments int
	if _, err := fmt.Scanln(&numberOfDepartments); err != nil {
		fmt.Println("when scanning number of departments: ", err)

		return
	}

	for range numberOfDepartments {
		var departmentCapacity int
		if _, err := fmt.Scanln(&departmentCapacity); err != nil {
			fmt.Println("when scanning capacity of department: ", err)

			return
		}

		temperatureRange := conditioner.NewTemperature(conditioner.MinTemp, conditioner.MaxTemp)

		for range departmentCapacity {
			var (
				sign string
				temp int
			)

			if _, err := fmt.Scanln(&sign, &temp); err != nil {
				fmt.Println("error scanning temperature wanted: ", err)

				return
			}

			bestTemp, err := temperatureRange.TempWantedByEmployee(sign, temp)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(bestTemp)
			}
		}
	}
}
