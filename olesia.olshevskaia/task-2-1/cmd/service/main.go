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

func applyCondition(sign string, temperature, minTemperature, maxTemperature int) (int, int) {
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

	return minTemperature, maxTemperature
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

	for range countDepartments {
		if _, err := fmt.Scanln(&countEmployee); err != nil {
			fmt.Println(err)

			return
		}

		minTemperature := Min
		maxTemperature := Max

		for range countEmployee {
			if _, err := fmt.Scanln(&sign, &temperature); err != nil {
				fmt.Println(err)

				return
			}

			minTemperature, maxTemperature = applyCondition(sign, temperature, minTemperature, maxTemperature)
			printOptimumTemperature(minTemperature, maxTemperature)
		}
	}
}
