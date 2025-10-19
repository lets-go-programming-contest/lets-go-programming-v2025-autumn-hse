package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	MinTempLimit = 15
	MaxTempLimit = 30
)

var ErrInvalidOperator = errors.New("invalid operator")

type Department struct {
	minTemp int
	maxTemp int
	isValid bool
}

func NewDepartment() *Department {
	return &Department{
		minTemp: MinTempLimit,
		maxTemp: MaxTempLimit,
		isValid: true,
	}
}

func (d *Department) ProcessConstraint(operator string, temperature int) (string, error) {
	if !d.isValid {
		return "-1", nil
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
		d.isValid = false
		return "-1", fmt.Errorf("%w: %s", ErrInvalidOperator, operator)
	}

	if d.minTemp <= d.maxTemp {
		return strconv.Itoa(d.minTemp), nil
	}

	d.isValid = false

	return "-1", nil
}

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Invalid department number:", err)
		return
	}

	for i := 0; i < departmentsCount; i++ {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Invalid employee number:", err)
			return
		}

		dept := NewDepartment()

		for j := 0; j < employeesCount; j++ {
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

