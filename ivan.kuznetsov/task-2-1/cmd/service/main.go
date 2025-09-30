package main

import (
	"fmt"

	"github.com/kuzid-17/task-2-1/internal/temperature"
)

func main() {
	var (
		departmentsCount, employeesCount, temperatureValue int
		limit                                              string
	)

	_, err := fmt.Scan(&departmentsCount)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departmentsCount {
		_, err = fmt.Scan(&employeesCount)
		if err != nil {
			fmt.Println("Invalid number of employees")

			return
		}

		values := []int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
		for range employeesCount {
			_, err = fmt.Scan(&limit)
			if err != nil {
				fmt.Println("Invalid limit format")

				return
			}

			_, err = fmt.Scan(&temperatureValue)
			if err != nil {
				fmt.Println("Invalid temperature value")

				return
			}

			values = temperature.OptimalTemperature(limit, temperatureValue, values)
			if len(values) > 0 {
				fmt.Println(values[0])
			} else {
				fmt.Println(-1)
			}
		}
	}
}
