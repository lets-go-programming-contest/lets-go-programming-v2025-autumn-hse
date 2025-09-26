package main

import (
	"fmt"

	"github.com/JingolBong/task-2-1/internal/conditioner"
)

func main() {
	var numberOfDepartments int64
	fmt.Scanln(&numberOfDepartments)
	for i := int64(0); i < numberOfDepartments; i++ {
		conditioner.TemperatureWantedDepartment()
	}
}
