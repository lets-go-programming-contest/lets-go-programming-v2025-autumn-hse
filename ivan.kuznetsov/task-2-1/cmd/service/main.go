package main

import (
	"fmt"

	"github.com/kuzid-17/task-2-1/internal/temperature"
)

func main() {
	var departmentsCount int

	_, err := fmt.Scan(&departmentsCount)
	if err != nil {
		fmt.Printf("Invalid number of departments: %v\n", err)

		return
	}

	for range departmentsCount {
		var employeesCount int

		_, err = fmt.Scan(&employeesCount)
		if err != nil {
			fmt.Printf("Invalid number of employees: %v\n", err)

			return
		}

		temperatureRange := temperature.TemperatureRangeInit()

		var limitSign string

		for range employeesCount {
			_, err = fmt.Scan(&limitSign)
			if err != nil {
				fmt.Printf("Invalid limit format: %v\n", err)

				return
			}

			var temperatureLimitValue int

			_, err = fmt.Scan(&temperatureLimitValue)
			if err != nil {
				fmt.Printf("Invalid temperature value: %v\n", err)

				return
			}

			temperatureRange = temperature.OptimalTemperature(limitSign, temperatureLimitValue, temperatureRange)

			if temperatureRange == nil {
				fmt.Printf("Invalid comparison sign '%s'\nThe temperature range has not changed\n", limitSign)
			}

			fmt.Println(temperature.GetOptimalTemperature(temperatureRange))
		}
	}
}
