package main

import (
	"fmt"
)

func printOptimalTemperature(departmentCount int) {
	var (
		employeeCount int
		value         int
		mathSign      string
	)

	for range departmentCount {
		upperBound := 30
		lowerBound := 15

		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("Reading error", err)

			return
		}

		for range employeeCount {
			if _, err := fmt.Scan(&mathSign, &value); err != nil {
				fmt.Println("Reading error", err)

				return
			}

			switch mathSign {
			case "<=":
				upperBound = min(value, upperBound)
			case ">=":
				lowerBound = max(value, lowerBound)

			default:
				fmt.Println("Wrong format")

				return
			}

			if upperBound < lowerBound {
				fmt.Println(-1)

				continue
			}

			fmt.Println(lowerBound)
		}
	}
}

func main() {
	var departmentCount int
	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("Reading error")

		return
	}

	printOptimalTemperature(departmentCount)
}
