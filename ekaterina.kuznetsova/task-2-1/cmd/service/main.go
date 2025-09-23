package main

import (
	"fmt"
)

func lessOrEqual(maxTemperature, minTemperature *int, temperature int) bool {
	if temperature < *minTemperature || temperature < 15 {
		fmt.Println("-1")

		return true
	}

	if temperature < *maxTemperature {
		*maxTemperature = temperature
	}
	fmt.Println(*minTemperature)

	return false
}

func moreOrEqual(maxTemperature, minTemperature *int, temperature int) bool {
	if *maxTemperature < temperature || temperature > 30 {
		fmt.Println("-1")

		return true
	}

	if *minTemperature < temperature {
		*minTemperature = temperature
	}

	fmt.Println(*minTemperature)

	return false
}

func main() {
	var (
		numberDepartaments, numberEmployees, temperature int
		minTemperature, maxTemperature                   = 15, 30
		comparisonSign                                   string
		flagFail                                         bool
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
		flagFail = false

		for range numberEmployees {
			_, err := fmt.Scanf("%s %d\n", &comparisonSign, &temperature)
			if err != nil {
				return
			}

			switch comparisonSign {
			case "<=":
				if !flagFail {
					flagFail = lessOrEqual(&maxTemperature, &minTemperature, temperature)
				} else {
					fmt.Println("-1")
				}
			case ">=":
				if !flagFail {
					flagFail = moreOrEqual(&maxTemperature, &minTemperature, temperature)
				} else {
					fmt.Println("-1")
				}
			default:
				fmt.Println("Error compaison sign")
			}
		}
	}
}
