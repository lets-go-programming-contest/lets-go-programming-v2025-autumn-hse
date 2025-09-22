package main

import (
	"fmt"
)

func lessOrEqual(maxTemperature, minTemperature *int, temperature int) {
	if temperature < *minTemperature || temperature > *maxTemperature || temperature < 15 {
		fmt.Println("-1")

		return
	}

	if temperature < *maxTemperature {
		*maxTemperature = temperature
	}

	fmt.Println(*minTemperature)
}

func moreOrEqual(maxTemperature, minTemperature *int, temperature int) {
	if *maxTemperature < temperature || temperature < *minTemperature || temperature > 30 {
		fmt.Println("-1")

		return
	}

	if *minTemperature < temperature {
		*minTemperature = temperature
	}

	fmt.Println(*minTemperature)
}

func main() {
	var (
		numberDepartaments, numberEmployees, temperature int
		minTemperature, maxTemperature                   = 15, 30
		comparisonSign                                   string
	)

	_, err := fmt.Scanln(&numberDepartaments)
	if err != nil {
		return
	}

	for range numberDepartaments {
		_, err := fmt.Scanln(&numberEmployees)
		if err != nil {
			return
		}

		minTemperature = 15
		maxTemperature = 30

		for range numberEmployees {
			_, err := fmt.Scanf("%s %d\n", &comparisonSign, &temperature)
			if err != nil {
				return
			}

			switch comparisonSign {
			case "<=":
				lessOrEqual(&maxTemperature, &minTemperature, temperature)
			case ">=":
				moreOrEqual(&maxTemperature, &minTemperature, temperature)
			default:
				fmt.Println("Error compaison sign")
			}
		}
	}
}
