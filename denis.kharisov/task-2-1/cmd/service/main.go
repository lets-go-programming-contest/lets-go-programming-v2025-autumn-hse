package main

import (
	"fmt"
)

type operation string

const (
	defaultMaxTemperature             = 30
	defaultMinTemperature             = 15
	lessOrEqualOperation    operation = "<="
	greaterOrEqualOperation operation = ">="
)

type temperatureRange struct {
	maxTemperature int
	minTemperature int
}

func optimalTemperature(tr temperatureRange, op operation, temperature int) (temperatureRange, error) {
	switch op {
	case greaterOrEqualOperation:
		if temperature > tr.maxTemperature {
			return tr, fmt.Errorf("max temperature > temperature")
		} else {
			tr.minTemperature = temperature
		}
	case lessOrEqualOperation:
		if temperature < tr.minTemperature {
			return tr, fmt.Errorf("temperature < min temperature")
		} else {
			tr.maxTemperature = temperature
		}
	}
	if tr.maxTemperature < tr.minTemperature {
		return tr, fmt.Errorf("max temperature < min temperature")
	}
	return tr, nil
}

func main() {
	var (
		numberOfDepartments, numberOfEmployees, temperature  int
		op                  								 operation
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
		tr = temperatureRange{
			maxTemperature: defaultMaxTemperature,
			minTemperature: defaultMinTemperature,
		}
		)

		for employee := 1; employee <= numberOfEmployees; employee++ {
			_, err = fmt.Scanln(&op, &temperature)
			if err != nil {
				fmt.Println("Invalid operation and temperature")
			}
			tr, err := optimalTemperature(tr, op, temperature)
			if err != nil {
				errorFlag = true
			}
			if errorFlag {
				fmt.Println("-1")
				continue
			}
			fmt.Println(tr.minTemperature)
		}
		errorFlag = false
	}
}
