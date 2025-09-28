package main

import (
	"errors"
	"fmt"
)

const (
	minTemperatureConst = 15
	maxTemperatureConst = 30
)

var (
	ErrTemperatureTooBig   = errors.New("Temperature is too big")
	ErrTemperatureTooSmall = errors.New("Temperature is too small")
	ErrTemperatureFail     = errors.New("Temperature is fail")
	ErrInvalidComparison   = errors.New("Error compaison sign")
)

func lessOrEqual(maxTemperature, minTemperature *int, temperature int) error {
	if temperature < *minTemperature || temperature < minTemperatureConst {
		return ErrTemperatureTooSmall
	}

	if temperature < *maxTemperature {
		*maxTemperature = temperature
	}
	return nil
}

func moreOrEqual(maxTemperature, minTemperature *int, temperature int) error {
	if *maxTemperature < temperature || temperature > maxTemperatureConst {
		return ErrTemperatureTooBig
	}

	if *minTemperature < temperature {
		*minTemperature = temperature
	}

	return nil
}

func compareValues(maxTemperature, minTemperature *int, temperature int, comparisonSign string) error {

	switch comparisonSign {
	case "<=":
		return lessOrEqual(maxTemperature, minTemperature, temperature)
	case ">=":
		return moreOrEqual(maxTemperature, minTemperature, temperature)
	default:
		return ErrInvalidComparison
	}
}

func main() {
	var (
		numberDepartaments, numberEmployees, temperature int
		comparisonSign                                   string
	)

	_, err := fmt.Scanln(&numberDepartaments)
	if err != nil {
		fmt.Println("Error scan number of departaments:", err)

		return
	}

	for range numberDepartaments {
		_, err := fmt.Scanln(&numberEmployees)
		var errTemperature error
		if err != nil {
			fmt.Println("Error scan number of employees:", err)

			return
		}

		minTemperature := minTemperatureConst
		maxTemperature := maxTemperatureConst

		for range numberEmployees {
			_, err = fmt.Scanf("%s %d\n", &comparisonSign, &temperature)
			if err != nil {
				return
			}

			if errTemperature != nil {
				fmt.Println("-1")
			} else {
				errTemperature = compareValues(&maxTemperature, &minTemperature, temperature, comparisonSign)
				if errTemperature != nil {
					fmt.Println(-1)
				} else {
					fmt.Println(minTemperature)
				}
			}
		}
	}
}
