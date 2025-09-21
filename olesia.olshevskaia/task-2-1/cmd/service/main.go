package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Min = 15
	Max = 30
)

func printOptimumTemperature(minTemperature, maxTemperature int) {
	switch {
	case minTemperature > maxTemperature:
		fmt.Println(-1)
	case minTemperature > Min && maxTemperature < Max:
		fmt.Println(minTemperature)
	case minTemperature > Min:
		fmt.Println(minTemperature)
	default:
		fmt.Println(maxTemperature)
	}
}

func main() {
	var (
		countDepartments, countEmployee, minTemperature, maxTemperature int
		sign                                                            string
	)

	if _, err := fmt.Scanln(&countDepartments); err != nil {
		fmt.Println(err)

		return
	}

	for range make([]struct{}, countDepartments) {
		if _, err := fmt.Scanln(&countEmployee); err != nil {
			fmt.Println(err)

			return
		}

		minTemperature = Min
		maxTemperature = Max

		for range make([]struct{}, countEmployee) {
			var line string
			if _, err := fmt.Scanln(&line); err != nil {
				fmt.Println(err)

				return
			}

			sign = line[:2]
			tempStr := strings.TrimSpace(line[2:])

			temperature, err := strconv.Atoi(tempStr)
			if err != nil {
				fmt.Println(err)

				return
			}

			switch sign {
			case ">=":
				if minTemperature < temperature {
					minTemperature = temperature
				}
			case "<=":
				if maxTemperature > temperature {
					maxTemperature = temperature
				}
			}

			printOptimumTemperature(minTemperature, maxTemperature)
		}
	}
}
