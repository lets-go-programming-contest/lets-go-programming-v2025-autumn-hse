package main

import (
	"fmt"
	"log"
)

const (
	minimumTemperature = 15
	maximumTemperature = 30
)

func main() {
	var (
		departmentCount, employeeCount, preferredTemperature, currentTemperature int
		inequalitySign                                                           string
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

		lowerBound := minimumTemperature
		upperBound := maximumTemperature
		currentTemperature = -1

		for range employeeCount {
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
				switch {
				case preferredTemperature > upperBound:
					currentTemperature = -1
				case preferredTemperature > lowerBound:
					lowerBound = preferredTemperature
					currentTemperature = lowerBound
				default:
					currentTemperature = lowerBound
				}
			case "<=":
				switch {
				case preferredTemperature < lowerBound:
					currentTemperature = -1
				case preferredTemperature < upperBound:
					upperBound = preferredTemperature
					currentTemperature = lowerBound
				default:
					currentTemperature = lowerBound
				}
			default:
				currentTemperature = -1
			}

			log.Println(currentTemperature)
		}
	}
}
