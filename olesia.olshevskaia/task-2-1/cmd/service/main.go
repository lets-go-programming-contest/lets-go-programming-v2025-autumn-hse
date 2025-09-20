package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		countDepartments, countEmployee, temputer, maxTemputer, minTemputer int
		sign                                                                string
	)
	fmt.Scan(&countDepartments)
	for i := 0; i < countDepartments; i++ {
		fmt.Scan(&countEmployee)
		maxTemputer = 30
		minTemputer = 15

		for j := 0; j < countEmployee; j++ {
			var line string
			fmt.Scan(&line)
			sign = line[:2]
			tempStr := line[2:]
			temputer, _ = strconv.Atoi(tempStr)
			switch sign {
			case ">=":
				if minTemputer < temputer {
					minTemputer = temputer
				}
			case "<=":
				if maxTemputer > temputer {
					maxTemputer = temputer
				}
			}
			if minTemputer > maxTemputer {
				fmt.Println(-1)
			} else if minTemputer > 15 && maxTemputer < 30 {
				fmt.Println(minTemputer)
			} else if minTemputer > 15 {
				fmt.Println(minTemputer)
			} else {
				fmt.Println(maxTemputer)
			}
		}
	}
}
