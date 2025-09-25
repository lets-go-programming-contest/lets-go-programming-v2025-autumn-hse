package main

import (
	"fmt"
	"strconv"
)

func getOptimalTemperature(departmentCount int) {
	var (
		employeeCount int
		inputBound    string
		value         int
	)
	for range departmentCount {
		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Reading error")

			return
		}
		upperBound := 30
		lowerBound := 15
		for range employeeCount {
			_, err = fmt.Scan(&inputBound)
			if err != nil {
				fmt.Println("Reading error")

				return
			}
			mathSign := inputBound[:2]
			if mathSign != "<=" && mathSign != ">=" {
				fmt.Println("Wrong input format")

				return
			}
			value, err = strconv.Atoi(inputBound[2:])
			if err != nil {
				fmt.Println("Error converting string to int")

				return
			}
			if mathSign == "<=" && value < upperBound {
				upperBound = value
			}
			if mathSign == ">=" && lowerBound < value {
				lowerBound = value
			}
			if mathSign == ">=" && value > upperBound {
				fmt.Println(-1)

				continue
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
	getOptimalTemperature(departmentCount)
}
