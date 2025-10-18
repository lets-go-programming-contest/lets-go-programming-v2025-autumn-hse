package main

import (
	"fmt"

	"github.com/kef1rch1k/task-2-1/internal/temperature"
)

const (
	MaxTemp = 30
	MinTemp = 15
)

func main() {
	var number int

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("unable to read:", err)

		return
	}

	for range number {
		var count int

		_, err = fmt.Scan(&count)
		if err != nil {
			fmt.Println("unable to read:", err)

			return
		}

		values := temperature.NewValues(MaxTemp, MinTemp)

		for range count {
			var (
				operation string
				temp      int
			)

			_, err := fmt.Scanln(&operation, &temp)
			if err != nil {
				fmt.Println("unable to read:", err)

				return
			}

			result, err := values.FindTemp(operation, temp)
			if err != nil {
				fmt.Println("error:", err)

				return
			}

			fmt.Println(result)
		}
	}
}
