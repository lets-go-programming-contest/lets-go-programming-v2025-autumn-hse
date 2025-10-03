package main

import (
	"fmt"

	"github.com/JingolBong/task-2-1/internal/conditioner"
)

func main() {
	var numberOfDepartments int
	if _, err := fmt.Scanln(&numberOfDepartments); err != nil {
		fmt.Println("when scanning number of departments", err)

		return
	}

	for range numberOfDepartments {
		conditioner.TemperatureWantedDepartment()
	}
}
