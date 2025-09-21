package main

import (
	"fmt"
)

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

		for range numberEmployees {
			_, err := fmt.Scanf("%s %d", &comparisonSign, &temperature)
			if err != nil {
				return
			}

			switch comparisonSign {
			case "<=":
				if temperature <= maxTemperature {
					maxTemperature = temperature
				}
				if temperature <= minTemperature {
					fmt.Println("-1")
					break
				}
				fmt.Println(minTemperature)
			case ">=":
				if minTemperature <= temperature {
					minTemperature = temperature
				}
				if maxTemperature <= temperature {
					fmt.Println("-1")
					break
				}
				fmt.Println(minTemperature)
			default:
				fmt.Println("Error compaison sign")
			}
		}
		minTemperature = 15
		maxTemperature = 30
	}
}
