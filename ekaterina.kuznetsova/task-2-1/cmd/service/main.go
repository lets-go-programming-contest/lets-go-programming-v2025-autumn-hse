package main

import (
	"fmt"
)

var flagFail = false

func lessOrEqual(maxTemperature, minTemperature *int, temperature int) {
	if temperature < *minTemperature || temperature < 15 {
		if !flagFail {
			fmt.Println("-1")
		}
		flagFail = true

		return
	}

	if temperature < *maxTemperature {
		*maxTemperature = temperature
	}

	fmt.Println(*minTemperature)
}

func moreOrEqual(maxTemperature, minTemperature *int, temperature int) {
	if *maxTemperature < temperature || temperature > 30 {
		if !flagFail {
			fmt.Println("-1")
		}
		flagFail = true

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
		flagFail = false

		for range numberEmployees {
			_, err := fmt.Scanf("%s %d\n", &comparisonSign, &temperature)
			if err != nil {
				return
			}

			switch comparisonSign {
			case "<=":
				if !flagFail {
					lessOrEqual(&maxTemperature, &minTemperature, temperature)
				} else {
					fmt.Println("-1")
				}
			case ">=":
				if !flagFail {
					moreOrEqual(&maxTemperature, &minTemperature, temperature)
				} else {
					fmt.Println("-1")
				}
			default:
				fmt.Println("Error compaison sign")
			}
		}
	}
}
