package main

import (
	"fmt"

	"github.com/kef1rch1k/task-2-1/internal/temperature"
)

<<<<<<< HEAD
const (
	MaxTemp = 30
	MinTemp = 15
)

func main() {
	var number int
=======
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
>>>>>>> b738d5c9f7fb824b1236f4c6877627be159127ef

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("unable to read:", err)

		return
	}

	for range number {
<<<<<<< HEAD
		var count int

=======
>>>>>>> b738d5c9f7fb824b1236f4c6877627be159127ef
		_, err = fmt.Scan(&count)
		if err != nil {
			fmt.Println("unable to read:", err)

			return
		}

<<<<<<< HEAD
		values := temperature.NewValues(MaxTemp, MinTemp)
=======
		values := temperature.NewValues()
>>>>>>> b738d5c9f7fb824b1236f4c6877627be159127ef

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

<<<<<<< HEAD
			result, err := values.FindTemp(operation, temp)
=======
			result, err := FindTemp(&values, operation, temp)
>>>>>>> b738d5c9f7fb824b1236f4c6877627be159127ef
			if err != nil {
				fmt.Println("error:", err)

				return
			}

			fmt.Println(result)
		}
	}
}
