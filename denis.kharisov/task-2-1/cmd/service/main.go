package main

import (
	"fmt"
	"errors"
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

func optimalTemperature(tr temperatureRange, operationType operation, temperature int) (temperatureRange, error) {
	switch operationType {
	case greaterOrEqualOperation:
		if temperature > tr.maxTemperature {
			return tr, errTemperatureExceedMax
		} else {
			tr.minTemperature = temperature
		}
	case lessOrEqualOperation:
		if temperature < tr.minTemperature {
			return tr, errTemperatureBelowMin
		} else {
			tr.maxTemperature = temperature
		}
	}
	if tr.maxTemperature < tr.minTemperature {
		return tr, errInvalidRange
	}
	return tr, nil
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
			tempRange, err := optimalTemperature(tempRange, operationType, temperature)
			if err != nil {
				errorFlag = true
			}
			if errorFlag {
				fmt.Println("-1")
				continue
			}
			fmt.Println(tempRange.minTemperature)
		}
	}
}
