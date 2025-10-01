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
		processor := temperature.NewTemperatureProcessor()
		result, err := processor.ProcessDepartment(os.Stdin)

		if err != nil {
			for i := range result {
				fmt.Println(result[i])
			}
		} else {
			fmt.Println(err)
		}
	}
}
