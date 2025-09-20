package main

import (
	"fmt"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}

	return a
}

const (
	minimumTemperature = 15
	maximumTemperature = 30
	errorTemperature   = -1
)

func main() {
	var departmentCount, employeeCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Error reading input:", err)

		return
	}

	for range departmentCount {
		_, err = fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Error reading input:", err)

			return
		}

		lowerBound := minimumTemperature
		upperBound := maximumTemperature

		for range employeeCount {
			var (
				preferredTemperature int
				inequalitySign       string
			)

			_, err = fmt.Scan(&inequalitySign, &preferredTemperature)
			if err != nil {
				fmt.Println("Error reading input:", err)

				return
			}

			if preferredTemperature < minimumTemperature || preferredTemperature > maximumTemperature {
				fmt.Println(errorTemperature)

				continue
			}

			switch inequalitySign {
			case ">=":
				lowerBound = maxInt(lowerBound, preferredTemperature)
			case "<=":
				upperBound = minInt(upperBound, preferredTemperature)
			}

			result := lowerBound
			if lowerBound > upperBound {
				result = errorTemperature
			}

			fmt.Println(result)
		}

		fmt.Println()
	}
}
