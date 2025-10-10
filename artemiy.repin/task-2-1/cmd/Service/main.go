package main

import (
	"fmt"

	"github.com/Nevermind0911/task-2-1/internal/temp"
)

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("couldn't read number of departments")
	}

	for range departments {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println("couldn't read number of employees")
		}

		temperature := temp.UpdateTemperature()

		for range employees {
			var (
				val      int
				operator string
			)

			if _, err := fmt.Scan(&operator, &val); err != nil {
				fmt.Println("couldn't read temp")
			}

			temperature.UpdateInterval(operator, val)
			opt := temperature.GetOptimal()
			fmt.Println(opt)
		}
	}
}
