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
				if value > upperBound {
					fmt.Println(-1)

					continue
				}
				lowerBound = max(value, lowerBound)

			default:
				fmt.Println("Wrong format")

				return
			}

			if value < lowerBound {
				fmt.Println(-1)

				continue
			}
			fmt.Println(lowerBound)
		}
	}
}

func main() {
	var departmentCount int
	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Reading error")

		return
	}
	printOptimalTemperature(departmentCount)
}
