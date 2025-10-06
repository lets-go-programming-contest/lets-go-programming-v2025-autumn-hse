package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	departmentsCount, _ := strconv.Atoi(scanner.Text())

	results := []string{}

	for range departmentsCount {
		scanner.Scan()
		employeesCount, _ := strconv.Atoi(scanner.Text())

		minTemperature := 15
		maxTemperature := 30
		isValid := true
		departmentResults := []string{}

		for range employeesCount {
			scanner.Scan()
			line := scanner.Text()
			parts := strings.Fields(line)
			
			if !isValid {
				departmentResults = append(departmentResults, "-1")
				continue
			}

			operator := parts[0]
			temperature, _ := strconv.Atoi(parts[1])

			if operator == ">=" {
				if temperature > minTemperature {
					minTemperature = temperature
				}
			} else if operator == "<=" {
				if temperature < maxTemperature {
					maxTemperature = temperature
				}
			}

			if minTemperature <= maxTemperature {
				departmentResults = append(departmentResults, strconv.Itoa(minTemperature))
			} else {
				departmentResults = append(departmentResults, "-1")
				isValid = false
			}
		}

		results = append(results, departmentResults...)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
