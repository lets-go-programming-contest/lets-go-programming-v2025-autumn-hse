package main

import (
	"errors"
	"fmt"
)

const (
	MinTempLimit = 15
	MaxTempLimit = 30
)

var errInvalidOperator = errors.New("invalid operator")

type Department struct {
	minTemp int
	maxTemp int
}

func NewDepartment(minTemp, maxTemp int) *Department {
	return &Department{
		minTemp: minTemp,
		maxTemp: maxTemp,
	}
}

func (d *Department) ProcessConstraint(operator string, temperature int) (int, error) {
	if !d.isValid() {
		return -1, nil
	}

	switch operator {
	case ">=":
		if temperature > d.minTemp {
			d.minTemp = temperature
		}
	case "<=":
		if temperature < d.maxTemp {
			d.maxTemp = temperature
		}
	default:
		return -1, fmt.Errorf("%w: %s", errInvalidOperator, operator)
	}

	if d.minTemp <= d.maxTemp {
		return d.minTemp, nil
	}

	return -1, nil
}

func (d *Department) isValid() bool {
	return d.minTemp <= d.maxTemp
}

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Invalid department number:", err)

		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Invalid employee number:", err)

			return
		}

		dept := NewDepartment(MinTempLimit, MaxTempLimit)

		for range employeesCount {
			var operator string

			if _, err := fmt.Scan(&operator); err != nil {
				fmt.Println("Invalid operator:", err)

				return
			}

			var temperature int

			if _, err := fmt.Scan(&temperature); err != nil {
				fmt.Println("Invalid temperature:", err)

				return
			}

			result, err := dept.ProcessConstraint(operator, temperature)
			if err != nil {
				fmt.Println("Invalid process:", err)

				return
			}

			fmt.Println(result)
		}
	}
}
