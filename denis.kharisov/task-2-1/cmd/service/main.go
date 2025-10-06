package main

import (
	"errors"
	"fmt"
)

type operation string

const (
	defaultMaxTemperature             = 30
	defaultMinTemperature             = 15
	lessOrEqualOperation    operation = "<="
	greaterOrEqualOperation operation = ">="
)

var (
	errTemperatureExceedMax = errors.New("temperature Exceed Max")
	errTemperatureBelowMin  = errors.New("temperature Below Min")
	errInvalidRange         = errors.New("max temperature < min temperature")
)

type temperatureRange struct {
	maxTemperature int
	minTemperature int
}

func optimalTemperature(tempRange temperatureRange, operType operation, temperature int) (temperatureRange, error) {
	switch operType {
	case greaterOrEqualOperation:
		if temperature > tempRange.maxTemperature {
			return tempRange, errTemperatureExceedMax
		} else {
			tempRange.minTemperature = max(temperature, tempRange.minTemperature)
		}
	case lessOrEqualOperation:
		if temperature < tempRange.minTemperature {
			return tempRange, errTemperatureBelowMin
		} else {
			tempRange.maxTemperature = min(temperature, tempRange.maxTemperature)
		}
	}

	if tempRange.maxTemperature < tempRange.minTemperature {
		return tempRange, errInvalidRange
	}

	return tempRange, nil
}

func main() {
	var (
		numberOfDepartments, numberOfEmployees, temperature int
		operationType                                       operation
	)

	_, err := fmt.Scanln(&numberOfDepartments)
	if err != nil {
		fmt.Println("Invalid number of departments")
	}

	for departmentNumber := 1; departmentNumber <= numberOfDepartments; departmentNumber++ {
		_, err = fmt.Scanln(&numberOfEmployees)
		if err != nil {
			fmt.Println("Invalid number of employees")
		}

		var (
			errorFlag = false
			tempRange = temperatureRange{
				maxTemperature: defaultMaxTemperature,
				minTemperature: defaultMinTemperature,
			}
		)

		for employee := 1; employee <= numberOfEmployees; employee++ {
			_, err = fmt.Scanln(&operationType, &temperature)
			if err != nil {
				fmt.Println("Invalid operation and temperature")
			}

			if !errorFlag {
				newRange, err := optimalTemperature(tempRange, operationType, temperature)
				if err != nil {
					errorFlag = true

					fmt.Println("-1")
				} else {
					tempRange = newRange
					fmt.Println(tempRange.minTemperature)
				}
			} else {
				fmt.Println("-1")
			}
		}
	}
}
