package main

import "fmt"

const (
	minStartingTemp = 15
	maxStartingTemp = 30
)

func main() {
	var (
		numOfDepartments, employees, temp      int
		minComfortableTemp, maxComfortableTemp int
		operator                               string
	)

	_, err := fmt.Scanln(&numOfDepartments)
	if err != nil {
		return
	}

	for range numOfDepartments {
		_, err := fmt.Scanln(&employees)
		if err != nil {
			return
		}

		minComfortableTemp = minStartingTemp
		maxComfortableTemp = maxStartingTemp
		for range employees {
			_, err = fmt.Scanln(&operator, &temp)
			if err != nil {
				return
			}
			switch operator {
			case "<=":
				maxComfortableTemp = min(maxComfortableTemp, temp)
			case ">=":
				minComfortableTemp = max(minComfortableTemp, temp)
			}

			if minComfortableTemp > maxComfortableTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minComfortableTemp)
			}
		}
	}
}
