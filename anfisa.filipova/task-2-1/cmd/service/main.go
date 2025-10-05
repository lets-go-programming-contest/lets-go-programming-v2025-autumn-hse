package main

import (
	"fmt"
)

type TemperatureBound struct {
	UpperBound int
	LowerBound int
}

func (t *TemperatureBound) setUpperBound(newBound int) {
	t.UpperBound = newBound
}

func (t *TemperatureBound) setLowerBound(newBound int) {
	t.LowerBound = newBound
}

func (t TemperatureBound) getUpperBound() int {
	return t.UpperBound
}
func (t TemperatureBound) getLowerBound() int {
	return t.LowerBound
}
func printOptimalTemperature(departmentCount int) {
	var (
		employeeCount int
		mathSign      string
	)
	const (
		defaultUpperBound int = 30
		defaultLowerBound int = 15
	)

	for range departmentCount {
		bound := TemperatureBound{defaultUpperBound, defaultLowerBound}

		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("Reading error", err)

			return
		}

		for range employeeCount {
			var value int
			if _, err := fmt.Scan(&mathSign, &value); err != nil {
				fmt.Println("Reading error", err)

				return
			}

			switch mathSign {
			case "<=":
				bound.setUpperBound(min(value, bound.getUpperBound()))
			case ">=":
				bound.setLowerBound(max(value, bound.getLowerBound()))

			default:
				fmt.Println("Wrong format")

				return
			}

			if bound.getUpperBound() < bound.getLowerBound() {
				fmt.Println(-1)

				continue
			}

			fmt.Println(bound.getLowerBound())
		}
	}
}

func main() {
	var departmentCount int
	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("Reading error")

		return
	}

	printOptimalTemperature(departmentCount)
}
