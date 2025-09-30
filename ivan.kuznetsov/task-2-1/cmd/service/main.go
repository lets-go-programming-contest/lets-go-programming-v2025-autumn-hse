package main

import "fmt"

func OptimalTemperature(limit string, T int, numbers []int) []int {
	var result []int
	for _, temp := range numbers {
		if limit == "<=" {
			if temp <= T {
				result = append(result, temp)
			}
		} else if limit == ">=" {
			if temp >= T {
				result = append(result, temp)
			}
		}
	}
	return result
}

func main() {
	var (
		N, K, T int
		limit   string
	)

	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Invalid number of departments")
		return
	}

	for range N {
		_, err = fmt.Scan(&K)
		if err != nil {
			fmt.Println("Invalid number of employees")
			return
		}

		values := []int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
		for range K {
			_, err = fmt.Scan(&limit)
			if err != nil {
				fmt.Println("Invalid limit format")
				return
			}

			_, err = fmt.Scan(&T)
			if err != nil {
				fmt.Println("Invalid temperature value")
				return
			}

			values = OptimalTemperature(limit, T, values)
			if len(values) > 0 {
				fmt.Println(values[0])
			} else {
				fmt.Println(-1)
			}
		}
	}
}
