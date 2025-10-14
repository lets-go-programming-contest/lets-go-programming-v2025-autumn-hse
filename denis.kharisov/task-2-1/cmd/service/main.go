package main

import (
	"errors"
	"fmt"
)

type operation string

const (
	lessOrEqualOperation    operation = "<="
	greaterOrEqualOperation operation = ">="
)

const (
	defaultMaxTemperature = 30
	defaultMinTemperature = 15
)

var (
	errTemperatureExceedMax = errors.New("temperature Exceed Max")
	errTemperatureBelowMin  = errors.New("temperature Below Min")
	errInvalidRange         = errors.New("max temperature < min temperature")
	errUnknownOperation     = errors.New("unknown operation")
)

type temperatureRange struct {
	maxTemperature int
	minTemperature int
	invalid        bool
}

func (tempRange *temperatureRange) optimalTemperature(operType operation, temperature int) error {
	if tempRange.invalid {
		return errInvalidRange
	}

	switch operType {
	case greaterOrEqualOperation:
		if temperature > tempRange.maxTemperature {
			tempRange.invalid = true
			return errTemperatureExceedMax
		}
		tempRange.minTemperature = max(temperature, tempRange.minTemperature)
	case lessOrEqualOperation:
		if temperature < tempRange.minTemperature {
			tempRange.invalid = true
			return errTemperatureBelowMin
		}
		tempRange.maxTemperature = min(temperature, tempRange.maxTemperature)

	default:
		return errUnknownOperation
	}

	if tempRange.maxTemperature < tempRange.minTemperature {
		tempRange.invalid = true
		return errInvalidRange
	}

	return nil
}

func main() {
	var numberOfDepartments int

	_, err := fmt.Scanln(&numberOfDepartments)
	if err != nil {
		fmt.Printf("Invalid number of departments: %v\n", err)

		return
	}

	var numberOfEmployees int
	for range numberOfDepartments {
		_, err = fmt.Scanln(&numberOfEmployees)
		if err != nil {
			fmt.Printf("Invalid number of employees: %v\n", err)

			return
		}

		var (
			temperature   int
			operationType operation

			tempRange = temperatureRange{
				maxTemperature: defaultMaxTemperature,
				minTemperature: defaultMinTemperature,
				invalid:        false,
			}
		)

		for range numberOfEmployees {
			_, err = fmt.Scanln(&operationType, &temperature)
			if err != nil {
				fmt.Printf("Invalid operation and temperature: %v\n", err)

				return
			}

			if err := tempRange.optimalTemperature(operationType, temperature); err != nil {
				fmt.Println("-1")
			} else {
				fmt.Println(tempRange.minTemperature)
			}
		}
	}
}
