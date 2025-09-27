package main

import (
	"fmt"
)

const (
	Min = 15
	Max = 30
)

func applyCondition(sign string, temperature, minTemperature, maxTemperature int) (int, int) {
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

	return minTemperature, maxTemperature
}

func printOptimumTemperature(minTemperature, maxTemperature int) {
	if minTemperature > maxTemperature {
		fmt.Println(-1)

		return
	}

	fmt.Println(minTemperature)
}

func main() {
	var (
		countDepartments, countEmployee, temperature int
		sign                                         string
	)

	if n, err := fmt.Scanln(&countDepartments); err != nil {
		fmt.Printf("Ошибка при чтении количества департаментов: считано %d значений, ошибка: %v\n", n, err)

		return
	} else if n != 1 {
		fmt.Printf("Ожидалось 1 значение для countDepartments, считано %d\n", n)

		return
	}

	for range countDepartments {
		if _, err := fmt.Scanln(&countEmployee); err != nil {
			fmt.Println(err)

			return
		}

		minTemperature := Min
		maxTemperature := Max

		for range countEmployee {
			if n, err := fmt.Scanln(&sign, &temperature); err != nil {
				fmt.Printf("Ошибка при чтении данных: считано %d значений, ошибка: %v\n", n, err)

				return
			} else if n != 2 {
				fmt.Printf("Ожидалось 2 значения (sign и temperature), считано %d\n", n)

				return
			}

			minTemperature, maxTemperature = applyCondition(sign, temperature, minTemperature, maxTemperature)
			printOptimumTemperature(minTemperature, maxTemperature)
		}
	}
}
