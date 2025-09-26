package main

import (
	"fmt"

	"github.com/JingolBong/task-2-1/internal/conditioner"
)

func main() {
	var numberOfDepartments int
	if _, err := fmt.Scanln(&numberOfDepartments); err != nil {
		return
	}
	for i := 0; i < numberOfDepartments; i++ {
		conditioner.TemperatureWantedDepartment()
	}
}
