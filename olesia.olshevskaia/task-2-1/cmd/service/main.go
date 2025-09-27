package main

import (
	"fmt"
)

const (
	Min                      = 15
	Max                      = 30
	expectedValuesPerInput   = 2
	expectedcountDepartments = 1
)

type TemperatureRange struct {
	min int
	max int
}

func (tr *TemperatureRange) Update(sign string, temperature int) {
	switch sign {
	case ">=":
		if tr.min < temperature {
			tr.min = temperature
		}
	case "<=":
		if tr.max > temperature {
			tr.max = temperature
		}
	}
}

func (tr *TemperatureRange) Optimum() int {
	if tr.min > tr.max {
		return -1
	}
	return tr.min
}

func NewTemperatureRange() TemperatureRange {
	return TemperatureRange{min: Min, max: Max}
}

func main() {
	var (
		countDepartments, countEmployee, temperature int
		sign                                         string
	)

	if n, err := fmt.Scanln(&countDepartments); err != nil {
		fmt.Printf("Ошибка при чтении количества департаментов: считано %d значений, ошибка: %v\n", n, err)

		return
	} else if n != expectedcountDepartments {
		fmt.Printf("Ожидалось 1 значение для countDepartments, считано %d\n", n)

		return
	}

	for range countDepartments {
		if n, err := fmt.Scanln(&countEmployee); err != nil {
			fmt.Printf("Ошибка при чтении количества сотрудников в департаменте: считано %d значений, ошибка: %v\n", n, err)

			return
		}

		tr := NewTemperatureRange()

		for range countEmployee {
			if n, err := fmt.Scanln(&sign, &temperature); err != nil {
				fmt.Printf("Ошибка при чтении данных: считано %d значений, ошибка: %v\n", n, err)

				return
			} else if n != expectedValuesPerInput {
				fmt.Printf("Ожидалось 2 значения (sign и temperature), считано %d\n", n)

				return
			}

			tr.Update(sign, temperature)
			fmt.Println(tr.Optimum())
		}
	}
}
