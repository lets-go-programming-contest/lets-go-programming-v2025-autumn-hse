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
			value, err = strconv.Atoi(inputBound[2:])
			if err != nil {
				fmt.Println("Error converting string to int")
			}
			switch mathSign {
			case "<=":
				if value < upperBound {
					upperBound = value
				}
			case ">=":
				if lowerBound < value {
					lowerBound = value
				}
				if value > upperBound {
					fmt.Println(-1)
					continue
				}
			default:
				fmt.Println("Wrong input format")
			}
			//fmt.Println(upperBound)
			//fmt.Println(lowerBound)
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
