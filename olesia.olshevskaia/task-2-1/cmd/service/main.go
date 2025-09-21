package main

import (
	"fmt"
)

const (
	Min = 15
	Max = 30
)

func printOptimumTemperature(minTemperature, maxTemperature int) {
	if minTemperature > maxTemperature {
		fmt.Println(-1)

		return
	}

	fmt.Println(minTemperature)
}

func main() {
	var (
		countDepartments, countEmployee, temperature int
		sign                                         string
	)

	if _, err := fmt.Scanln(&countDepartments); err != nil {
		fmt.Println(err)

		return
	}

	for range make([]struct{}, countDepartments) {
		if _, err := fmt.Scanln(&countEmployee); err != nil {
			fmt.Println(err)

			return
		}

		minTemperature := Min
		maxTemperature := Max

		for range make([]struct{}, countEmployee) {
			if _, err := fmt.Scanln(&sign, &temperature); err != nil {
				fmt.Println(err)

				return
			}

			switch sign {
			case ">=":
				if minTemperature < temperature {
					minTemperature = temperature
				}
			case "<=":
				if maxTemperature > temperature {
					maxTemperature = temperature
				}
			}

			printOptimumTemperature(minTemperature, maxTemperature)
		}
	}
}
