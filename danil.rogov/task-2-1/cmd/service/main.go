package main

import (
	"fmt"
	"os"

	"github.com/Tapochek2894/task-2/subtask-1/internal/temperature"
)

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Error reading department count:", err)

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err := fmt.Fscan(os.Stdin, &employeeCount)
		if err != nil {
			fmt.Println("Error reading employee count:", err)

			return
		}

		processor := temperature.NewTemperatureProcessor()
		err = processor.ProcessDepartment(employeeCount, os.Stdin, os.Stdout)

		if err != nil {
			fmt.Println("Error during processing department:", err)
		}
	}
}
