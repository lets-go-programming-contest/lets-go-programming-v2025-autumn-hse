package main

import (
	"fmt"
	"log"
)

const (
	minimumTemperature int = 15
	maximumTemperature int = 30
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
	for i := 0; i < departmentCount; i++ {
		_, err = fmt.Scan(&employeeCount)
		if err != nil {
			log.Fatal(err)
		}

		lowerBound := minimumTemperature
		upperBound := maximumTemperature
		currentTemperature := -1

		for j := 0; j < employeeCount; j++ {
			_, err = fmt.Scan(&inequalitySign, &preferredTemperature)
			if err != nil {
				log.Fatal(err)
			}

			if preferredTemperature < minimumTemperature || preferredTemperature > maximumTemperature {
				currentTemperature = -1
				log.Println(currentTemperature)
				continue
			}

			switch inequalitySign {
			case ">=":
				if preferredTemperature > upperBound {
					currentTemperature = -1
				} else if preferredTemperature > lowerBound {
					lowerBound = preferredTemperature
					currentTemperature = lowerBound
				} else {
					currentTemperature = lowerBound
				}
			case "<=":
				if preferredTemperature < lowerBound {
					currentTemperature = -1
				} else if preferredTemperature < upperBound {
					upperBound = preferredTemperature
					currentTemperature = lowerBound
				} else {
					currentTemperature = lowerBound
				}
			default:
				currentTemperature = -1
			}
			log.Println(currentTemperature)
		}
	}
}
