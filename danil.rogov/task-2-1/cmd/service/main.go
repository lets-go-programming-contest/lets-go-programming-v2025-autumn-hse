package main

import (
	"fmt"
	"log"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

const (
	minimumTemperature = 15
	maximumTemperature = 30
)

func main() {
	var (
		departmentCount, employeeCount, preferredTemperature int
		inequalitySign                                       string
	)

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		log.Fatal(err)
	}

	for range departmentCount {
		_, err = fmt.Scan(&employeeCount)
		if err != nil {
			log.Fatal(err)
		}

		var (
			lowerBound = minimumTemperature
			upperBound = maximumTemperature
		)

		for range employeeCount {
			_, err = fmt.Scan(&inequalitySign, &preferredTemperature)
			if err != nil {
				log.Fatal(err)
			}

			if preferredTemperature < minimumTemperature || preferredTemperature > maximumTemperature {
				fmt.Println(-1)

				continue
			}

			switch inequalitySign {
			case ">=":
				lowerBound = maxInt(lowerBound, preferredTemperature)
			case "<=":
				upperBound = minInt(upperBound, preferredTemperature)
			}
			if lowerBound > upperBound {
				fmt.Println(-1)
			} else {
				fmt.Println(lowerBound)
			}
		}
		fmt.Println()
	}
}
