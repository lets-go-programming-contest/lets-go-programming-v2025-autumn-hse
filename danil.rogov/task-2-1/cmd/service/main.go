package main

import (
	"fmt"
)

const (
	minimumTemperature int = 15
	maximumTemperature int = 30
)

func main() {
	var (
		departmentCount, employeeCount, preferredTemperature int
		inequalitySign                                       string
		lowerBound                                           int = minimumTemperature
		upperBound                                           int = maximumTemperature
	)
	fmt.Scan(&departmentCount)
	for range departmentCount {
		fmt.Scan(&employeeCount)
		var currentTemperature int
		for range employeeCount {
			fmt.Scan(&inequalitySign, &preferredTemperature)
			if preferredTemperature > maximumTemperature || preferredTemperature < minimumTemperature {
				currentTemperature = -1
			}
			switch inequalitySign {
			case ">=":
				if preferredTemperature <= upperBound {
					currentTemperature = preferredTemperature
					lowerBound = preferredTemperature
				} else {
					currentTemperature = -1
				}
			case "<=":
				if preferredTemperature >= lowerBound {
					currentTemperature = lowerBound
					upperBound = preferredTemperature
				} else {
					currentTemperature = -1
				}
			}
			fmt.Println(currentTemperature)
		}
	}
}
