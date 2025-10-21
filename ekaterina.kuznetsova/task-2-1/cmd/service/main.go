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
	ErrCompaisonSign       = errors.New("compaison sign is not match")
)

type TemperatureValidator struct {
	minTemperature int
	maxTemperature int
}

func (v *TemperatureValidator) lessOrEqual(temperature int) {
	if temperature < v.maxTemperature {
		v.maxTemperature = temperature
	}
}

func (v *TemperatureValidator) moreOrEqual(temperature int) {
	if v.minTemperature < temperature {
		v.minTemperature = temperature
	}
}

func (v *TemperatureValidator) compareValues(temperature int, comparisonSign string) error {
	switch comparisonSign {
	case "<=":
		v.lessOrEqual(temperature)
	case ">=":
		v.moreOrEqual(temperature)
	default:
		return ErrCompaisonSign
	}

	return nil
}

func (v *TemperatureValidator) getOptimum() int {
	if v.minTemperature > v.maxTemperature {
		return -1
	}

	return v.minTemperature
}

func main() {
	var numberDepartaments int

	_, err := fmt.Scanln(&numberDepartaments)
	if err != nil {
		fmt.Println("Error scan number of departaments:", err)

		return
	}

	var (
		numberEmployees, temperature int
		comparisonSign               string
	)

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

			err = validator.compareValues(temperature, comparisonSign)
			if err != nil {
				fmt.Println("Error:", err)

				return
			}

			fmt.Println(validator.getOptimum())
		}
	}
}
