package main

import (
	"fmt"

	"github.com/kuzid-17/task-2-1/internal/temperature"
)

func main() {
	var (
		departmentsCount, employeesCount, temperatureLimitValue int
		limitSign                                               string
	)

	_, err := fmt.Scan(&departmentsCount)
	if err != nil {
		fmt.Printf("Invalid number of departments: %v\n", err)

		return
	}

	for range departmentsCount {
		_, err = fmt.Scan(&employeesCount)
		if err != nil {
			fmt.Printf("Invalid number of employees: %v\n", err)

			return
		}

		temperatureRange := temperature.InitialRange()

		for range employeesCount {
			_, err = fmt.Scan(&limitSign)
			if err != nil {
				fmt.Printf("Invalid limit format: %v\n", err)

				return
			}

			_, err = fmt.Scan(&temperatureLimitValue)
			if err != nil {
				fmt.Printf("Invalid temperature value: %v\n", err)

				return
			}

			temperatureRange = temperature.OptimalTemperature(limitSign, temperatureLimitValue, temperatureRange)
			fmt.Println(temperature.GetOptimalTemperature(temperatureRange))
		}
	}
}
