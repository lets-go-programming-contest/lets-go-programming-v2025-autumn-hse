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
	errInvalidRange     = errors.New("max temperature < min temperature")
	errUnknownOperation = errors.New("unknown operation")
)

type temperatureRange struct {
	maxTemperature int
	minTemperature int
}

func NewTemperatureRange(maxTemp int, minTemp int) *temperatureRange {
	return &temperatureRange{
		maxTemperature: maxTemp,
		minTemperature: minTemp,
	}
}

func (tempRange *temperatureRange) optimalTemperature(operType operation, temperature int) (int, error) {
	switch operType {
	case greaterOrEqualOperation:
		if temperature >= tempRange.minTemperature {
			tempRange.minTemperature = temperature
		}

	case lessOrEqualOperation:
		if temperature <= tempRange.maxTemperature {
			tempRange.maxTemperature = temperature
		}

	default:
		return 0, errUnknownOperation
	}

	if tempRange.maxTemperature >= tempRange.minTemperature {
		return tempRange.minTemperature, nil
	}

	return -1, errInvalidRange
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

			tempRange = NewTemperatureRange(defaultMaxTemperature, defaultMinTemperature)
		)

		for range numberOfEmployees {
			_, err = fmt.Scanln(&operationType, &temperature)
			if err != nil {
				fmt.Printf("Invalid operation and temperature: %v\n", err)

				return
			}

			if result, err := tempRange.optimalTemperature(operationType, temperature); err != nil {
				fmt.Println(-1)
			} else {
				fmt.Println(result)
			}
		}
	}
}
