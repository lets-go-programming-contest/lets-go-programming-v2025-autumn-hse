package main

import (
	"fmt"
	"strconv"
)

func printOptimalTemperature(departmentCount int) {
	var (
		employeeCount int
		inputBound    string
		value         int
	)
	for range departmentCount {
		upperBound := 30
		lowerBound := 15
		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Reading error", err)

			return
		}
		for range employeeCount {
			_, err = fmt.Scan(&inputBound)
			if err != nil {
				fmt.Println("Reading error", err)

				return
			}
			mathSign := inputBound[:2]
			value, err = strconv.Atoi(inputBound[3:])
			if err != nil {
				fmt.Println("Error converting string to int", err)

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
