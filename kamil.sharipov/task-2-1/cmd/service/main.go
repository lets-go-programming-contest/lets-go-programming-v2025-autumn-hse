package main

import (
	"fmt"

	"github.com/kamilSharipov/task-2-1/internal/temperature"
)

func main() {
	var numOfDepartments int

	_, err := fmt.Scanln(&numOfDepartments)
	if err != nil {
		fmt.Printf("Error reading number of departments: %v\n", err)

		return
	}

	var employees int
	for range numOfDepartments {
		_, err := fmt.Scanln(&employees)
		if err != nil {
			fmt.Printf("Error reading number of employees: %v\n", err)

			return
		}

		comfortTemperature := temperature.NewComfortTemperature()

		var (
			operator string
			temp     int
		)

		for range employees {
			_, err = fmt.Scanln(&operator, &temp)
			if err != nil {
				fmt.Printf("Error reading operator and temperature: %v\n", err)

				return
			}

			parsedOp, err := temperature.ParseOperator(operator)
			if err != nil {
				fmt.Printf("Error parsing operator: %v\n", err)

				return
			}

			if comfortTemperature.AddConstraint(parsedOp, temp) != nil {
				fmt.Println("Failed to add constraint", err)

				return
			}

			fmt.Println(comfortTemperature.Result())
		}
	}
}
