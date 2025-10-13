package main

import (
	"fmt"

	"github.com/kef1rch1k/task-2-1/internal/temperature"
)

func FindTemp(values *temperature.Value, operation string, temp int) (int, error) {
	err := values.UpdateValues(operation, temp)
	if err != nil {
		return 0, fmt.Errorf("failed to update temperature values: %w", err)
	}

	if values.Lower <= values.Higher {
		return values.Lower, nil
	}

	return -1, nil
}

func main() {
	var number, count int

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("unable to read:", err)

		return
	}

	for range number {
		_, err = fmt.Scan(&count)
		if err != nil {
			fmt.Println("unable to read:", err)

			return
		}

		values := temperature.NewValues()

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

			result, err := FindTemp(&values, operation, temp)
			if err != nil {
				fmt.Println("error:", err)

				return
			}

			fmt.Println(result)
		}
	}
}
