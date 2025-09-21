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
				if temperature <= minTemperature {
					_, err := fmt.Println("-1")
					if err != nil {
						return
					}
					break
				}
				if temperature <= maxTemperature {
					maxTemperature = temperature
				}
				_, err := fmt.Println(minTemperature)
				if err != nil {
					return
				}
			case ">=":
				if maxTemperature <= temperature {
					_, err := fmt.Println("-1")
					if err != nil {
						return
					}
					break
				}
				if minTemperature <= temperature {
					minTemperature = temperature
				}
				_, err := fmt.Println(minTemperature)
				if err != nil {
					return
				}
			default:
				_, err := fmt.Println("Error compaison sign")
				if err != nil {
					return
				}
			}
		}
		minTemperature = 15
		maxTemperature = 30
	}
}
