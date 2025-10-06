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
	n, _ := strconv.Atoi(scanner.Text())

	results := []string{}

	for i := 0; i < n; i++ {

		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())

		minTemp := 15
		maxTemp := 30
		valid := true
		departmentResults := []string{}

		for j := 0; j < k; j++ {
			scanner.Scan()
			line := scanner.Text()
			parts := strings.Fields(line)

			if !valid {
				departmentResults = append(departmentResults, "-1")
				continue
			}

			operator := parts[0]
			temp, _ := strconv.Atoi(parts[1])

			if operator == ">=" {
				if temp > minTemp {
					minTemp = temp
				}
			} else if operator == "<=" {
				if temp < maxTemp {
					maxTemp = temp
				}
			}

			if minTemp <= maxTemp {
				departmentResults = append(departmentResults, strconv.Itoa(minTemp))
			} else {
				departmentResults = append(departmentResults, "-1")
				valid = false
			}
		}

		results = append(results, departmentResults...)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
