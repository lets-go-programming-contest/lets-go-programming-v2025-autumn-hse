package main

import (
	"fmt"
)

const (
	MinTempLimit = 15
	MaxTempLimit = 30
)

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
		return "-1", fmt.Errorf("Invalid operand: %s", operator)
	}

	if d.minTemp <= d.maxTemp {
		return fmt.Sprintf("%d", d.minTemp), nil
	} else {
		d.isValid = false
		return "-1", nil
	}
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Invalid dep number:", err)
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
			var temperature int

			if _, err := fmt.Scan(&operator, &temperature); err != nil {
				fmt.Println("Invalid scan:", err)
				return
			}

			result, err := dept.ProcessConstraint(operator, temperature)
			if err != nil {
				fmt.Println("Invalid processing:", err)
				return
			}

			fmt.Println(result)
		}
	}
}
