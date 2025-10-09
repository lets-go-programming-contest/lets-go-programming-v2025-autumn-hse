package main

import (
	"fmt"
)

const (
	minTemperatureValue      = 15
	maxTemperatureValue      = 30
	expectedValuesPerInput   = 2
	expectedcountDepartments = 1
)

type TemperatureRange struct {
	min int
	max int
}

func (tr *TemperatureRange) Update(sign string, temperature int) {
	switch sign {
	case ">=":
		if tr.min < temperature {
			tr.min = temperature
		}
	case "<=":
		if tr.max > temperature {
			tr.max = temperature
		}
	default:
		fmt.Printf("Incorrect temperature string input\n")
	}
}

func (tr *TemperatureRange) Optimum() int {
	if tr.min > tr.max {
		return -1
	}

	return tr.min
}

func NewTemperatureRange() TemperatureRange {
	return TemperatureRange{min: minTemperatureValue, max: maxTemperatureValue}
}

func main() {

	var countDepartments int

	returnVal, err := fmt.Scanln(&countDepartments)
	if err != nil {
		fmt.Printf("Error reading department count: %d values read, error: %v\n", returnVal, err)

		return
	}

	if returnVal != expectedcountDepartments {
		fmt.Printf("Expected 1 value for countDepartments, read %d\n", returnVal)

		return
	}

	var countEmployee int

	for range countDepartments {
		if returnVal, err = fmt.Scanln(&countEmployee); err != nil {
			fmt.Printf("Error reading department employee count: %d values read, error: %v\n", returnVal, err)

			return
		}

		temperatureRange := NewTemperatureRange()
		var (
			sign        string
			temperature int
		)

		for range countEmployee {
			if returnVal, err = fmt.Scanln(&sign, &temperature); err != nil {
				fmt.Printf("Error reading data: %d values read, error: %v\n", returnVal, err)

				return
			} else if returnVal != expectedValuesPerInput {
				fmt.Printf("Expected 2 values (sign and temperature), read %d\n", returnVal)

				return
			}

			temperatureRange.Update(sign, temperature)
			fmt.Println(temperatureRange.Optimum())
		}
	}
}
