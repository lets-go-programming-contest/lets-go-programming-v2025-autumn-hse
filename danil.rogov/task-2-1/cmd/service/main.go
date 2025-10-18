package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Tapochek2894/task-2/subtask-1/internal/temperature"
)

func readPreference(reader io.Reader) (temperature.Preference, error) {
	var pref temperature.Preference

	_, err := fmt.Fscan(reader, &pref.Sign, &pref.Temperature)
	if err != nil {
		return temperature.Preference{}, fmt.Errorf("failed to read preference: %w", err)
	}

	return pref, nil
}

func printResult(result int, writer io.Writer) error {
	_, err := fmt.Fprintln(writer, result)
	if err != nil {
		return fmt.Errorf("failed to write result: %w", err)
	}

	return nil
}

func processDepartment(employeeCount int, reader io.Reader, writer io.Writer) error {
	processor := temperature.NewTemperatureProcessor()

	for range employeeCount {
		pref, err := readPreference(reader)
		if err != nil {
			return fmt.Errorf("read preference: %w", err)
		}

		result, err := processor.AddPreference(pref)
		if err != nil {
			return fmt.Errorf("add preference: %w", err)
		}

		err = printResult(result, writer)
		if err != nil {
			return fmt.Errorf("print result: %w", err)
		}
	}

	return nil
}

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Error reading department count:", err)

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Error reading employee count:", err)

			return
		}

		err = processDepartment(employeeCount, os.Stdin, os.Stdout)
		if err != nil {
			fmt.Println("Error during processing department:", err)

			return
		}
	}
}
