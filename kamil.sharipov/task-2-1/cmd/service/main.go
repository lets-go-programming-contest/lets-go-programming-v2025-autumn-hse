package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/kamilSharipov/task-2-1/internal/temperature"
)

var (
	ErrReadingNumOfDepartments = errors.New("Error reading number of departments")
	ErrReadingNumOfEmployees   = errors.New("Error reading number of employees")
	ErrReadingOperatorTemp     = errors.New("Error reading operator and temperature")
	ErrParseOperator           = errors.New("Error parsing operator")
)

func main() {
	var numOfDepartments int
	_, err := fmt.Scanln(&numOfDepartments)
	if err != nil {
		fmt.Fprintln(os.Stderr, ErrReadingNumOfDepartments, err)
		return
	}

	var employees int
	for range numOfDepartments {
		_, err := fmt.Scanln(&employees)
		if err != nil {
			fmt.Fprintln(os.Stderr, ErrReadingNumOfEmployees, err)
			return
		}

		comfortTemperature := temperature.InitComfortTemperature()

		var (
			operator string
			temp     int
		)
		for range employees {
			_, err = fmt.Scanln(&operator, &temp)
			if err != nil {
				fmt.Fprintln(os.Stderr, ErrReadingOperatorTemp, err)
				return
			}

			op, err := temperature.ParseOperator(operator)
			if err != nil {
				fmt.Println(os.Stderr, ErrParseOperator, err)
				return
			} else {
				comfortTemperature.AddConstraint(op, temp)
			}

			fmt.Println(comfortTemperature.Result())
		}
	}
}
