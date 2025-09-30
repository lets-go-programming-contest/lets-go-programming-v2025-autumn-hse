package main

import (
	"errors"
	"fmt"

	"github.com/JingolBong/task-2-1/internal/conditioner"
)

var errFailedToScan = errors.New("invalid input")

func main() {
	var numberOfDepartments int
	if _, err := fmt.Scanln(&numberOfDepartments); err != nil {
		fmt.Println(errFailedToScan)

		return
	}

	for range numberOfDepartments {
		conditioner.TemperatureWantedDepartment()
	}
}
