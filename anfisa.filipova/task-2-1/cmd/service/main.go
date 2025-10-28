package main

import (
	"errors"
	"fmt"
)

const (
	defaultUpperBound int = 30
	defaultLowerBound int = 15
)

type TemperatureManager struct {
	UpperBound int
	LowerBound int
}

var (
	errSignFormat = errors.New("wrong math sign format ")
	errNotExist   = errors.New("optimal temperature does not existed")
)

func (t *TemperatureManager) updateBounds(mathSign string, value int) error {
	switch mathSign {
	case "<=":
		t.UpperBound = min(value, t.UpperBound)

	case ">=":
		t.LowerBound = max(value, t.LowerBound)

	default:
		return errSignFormat
	}

	return nil
}

func (t *TemperatureManager) getOptimalTemperature() (int, error) {
	if t.UpperBound < t.LowerBound {
		return 0, errNotExist
	}

	return t.LowerBound, nil
}

func main() {
	var departmentCount int
	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("Failed to read count of departments: ", err)

		return
	}

	for range departmentCount {
		bound := TemperatureManager{defaultUpperBound, defaultLowerBound}

		var employeeCount int
		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("Failed to read count of employees: ", err)

			return
		}

		for range employeeCount {
			var (
				mathSign string
				value    int
			)

			if _, err := fmt.Scan(&mathSign, &value); err != nil {
				fmt.Println("Failed to read bound: ", err)

				return
			}

			if err := bound.updateBounds(mathSign, value); err != nil {
				fmt.Println("Error: ", err)

				return
			}

			res, err := bound.getOptimalTemperature()
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(res)
		}
	}
}
