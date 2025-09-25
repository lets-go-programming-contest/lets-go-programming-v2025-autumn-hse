package main

import (
	"fmt"
	"strconv"
)

func printOptimalTemperature(departmentCount int) {
	var (
		employeeCount int
		inputBound    string
	)
	for range departmentCount {
		upperBound := 30
		lowerBound := 15
		fmt.Scan(&employeeCount)
		for range employeeCount {
			fmt.Scan(&inputBound)
			mathSign := inputBound[:2]
			value, err := strconv.Atoi(inputBound[2:])
			if err != nil {
				fmt.Println("Error converting string to int")

				return
			}
			if mathSign == "<=" {
				upperBound = min(value, upperBound)
			}
			if mathSign == ">=" {
				if value > upperBound {
					fmt.Println(-1)

					continue
				}
				lowerBound = max(value, lowerBound)
			}
			if value < lowerBound {
				fmt.Println(-1)
			} else {
				fmt.Println(lowerBound)
			}
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
