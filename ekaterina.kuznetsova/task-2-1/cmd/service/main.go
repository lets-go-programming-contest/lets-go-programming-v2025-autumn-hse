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
	ErrTemperatureTooBig   = errors.New("temperature is too big")
	ErrTemperatureTooSmall = errors.New("temperature is too small")
	ErrTemperatureFail     = errors.New("temperature is fail")
	ErrInvalidComparison   = errors.New("еrror compaison sign")
)

type TemperatureValidator struct {
	minTemperature int
	maxTemperature int
}

func (v *TemperatureValidator) lessOrEqual(temperature int) error {
	if temperature < v.minTemperature || temperature < minTemperatureConst {
		return ErrTemperatureTooSmall
	}

	if temperature < v.maxTemperature {
		v.maxTemperature = temperature
	}

	return nil
}

func (v *TemperatureValidator) moreOrEqual(temperature int) error {
	if v.maxTemperature < temperature || temperature > maxTemperatureConst {
		return ErrTemperatureTooBig
	}

	if v.minTemperature < temperature {
		v.minTemperature = temperature
	}

	return nil
}

func (v *TemperatureValidator) compareValues(temperature int, comparisonSign string) error {
	switch comparisonSign {
	case "<=":
		return v.lessOrEqual(temperature)
	case ">=":
		return v.moreOrEqual(temperature)
	default:
		return ErrInvalidComparison
	}
}

func main() {
	var (
		numberDepartaments, numberEmployees, temperature int
		comparisonSign                                   string
		errTemperature 									 error
	)

	_, err := fmt.Scanln(&numberDepartaments)
	if err != nil {
		fmt.Println("Error scan number of departaments:", err)

		return
	}

	for range numberDepartaments {
		_, err := fmt.Scanln(&numberEmployees)
		if err != nil {
			fmt.Println("Error scan number of employees:", err)

			return
		}

		validator := TemperatureValidator{
			minTemperature: minTemperatureConst,
			maxTemperature: maxTemperatureConst,
		}

		for range numberEmployees {
			_, err = fmt.Scanf("%s %d\n", &comparisonSign, &temperature)
			if err != nil {
				fmt.Println("Error scan comparison sign and temperature:", err)

				return
			}

			if errTemperature != nil {
				fmt.Println("-1")
			} else {
				errTemperature = validator.compareValues(temperature, comparisonSign)
				if errTemperature != nil {
					fmt.Println(-1)
				} else {
					fmt.Println(validator.minTemperature)
				}
			}
		}
	}
}
