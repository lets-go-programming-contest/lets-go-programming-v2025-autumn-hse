package main

import (
	"errors"
	"fmt"

	"github.com/kamilSharipov/task-2-1/internal/temperature"
)

var (
	ErrReadingNumOfDepartments = errors.New("error reading number of departments")
	ErrReadingNumOfEmployees   = errors.New("error reading number of employees")
	ErrReadingOperatorTemp     = errors.New("error reading operator and temperature")
	ErrParseOperator           = errors.New("error parsing operator")
)

func main() {
	var numOfDepartments int
	
	_, err := fmt.Scanln(&numOfDepartments)
	if err != nil {
		fmt.Println(ErrReadingNumOfDepartments, err)

		return
	}

	var employees int
	for range numOfDepartments {
		_, err := fmt.Scanln(&employees)
		if err != nil {
			fmt.Println(ErrReadingNumOfEmployees, err)

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
				fmt.Println(ErrReadingOperatorTemp, err)

				return
			}

			parsedOp, err := temperature.ParseOperator(operator)
			if err != nil {
				fmt.Println(ErrParseOperator, err)

				return
			} else {
				comfortTemperature.AddConstraint(parsedOp, temp)
			}

			fmt.Println(comfortTemperature.Result())
		}
	}
}
