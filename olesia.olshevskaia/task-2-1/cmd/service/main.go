package main

import (
	"errors"
	"fmt"
)

const (
	minTemperatureValue = 15
	maxTemperatureValue = 30
)

var ErrInvalidSign = errors.New("invalid sign")

type TemperatureRange struct {
	min int
	max int
}

func NewTemperatureRange(minValue, maxValue int) TemperatureRange {
	return TemperatureRange{min: minValue, max: maxValue}
}

func (tr *TemperatureRange) Update(sign string, temperature int) error {
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
		return fmt.Errorf("%w: %s", ErrInvalidSign, sign)
	}

	return nil
}

func (tr *TemperatureRange) Optimum() int {
	if tr.min > tr.max {
		return -1
	}

	return tr.min
}

func main() {
	var (
		countDepartments, returnVal int
		err                         error
	)

	returnVal, err = fmt.Scanln(&countDepartments)
	if err != nil {
		fmt.Printf("Error reading department count: %d values read, error: %v\n", returnVal, err)

		return
	}

	var countEmployee int

	for range countDepartments {
		if returnVal, err = fmt.Scanln(&countEmployee); err != nil {
			fmt.Printf("Error reading department employee count: %d values read, error: %v\n", returnVal, err)

			return
		}

		temperatureRange := NewTemperatureRange(minTemperatureValue, maxTemperatureValue)

		var (
			sign        string
			temperature int
		)

		for range countEmployee {
			if returnVal, err = fmt.Scanln(&sign, &temperature); err != nil {
				fmt.Printf("Error reading data: %d values read, error: %v\n", returnVal, err)

				return
			}

			if err := temperatureRange.Update(sign, temperature); err != nil {
				fmt.Println("incorrect temperature sign", err)

				return
			}

			fmt.Println(temperatureRange.Optimum())
		}
	}
}
