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

	fmt.Scanln(&numberDepartaments)
	for i := 0; i < numberDepartaments; i++ {
		fmt.Scanln(&numberEmployees)
		for j := 0; j < numberEmployees; j++ {
			fmt.Scanf("%s %d", &comparisonSign, &temperature)
			switch comparisonSign {
			case "<=":
				if temperature < minTemperature {
					fmt.Println("-1")
					break
				}
				if temperature <= maxTemperature {
					maxTemperature = temperature
				}
				fmt.Println(minTemperature)
			case ">=":
				if maxTemperature < temperature {
					fmt.Println("-1")
					break
				}
				if minTemperature <= temperature {
					minTemperature = temperature
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
