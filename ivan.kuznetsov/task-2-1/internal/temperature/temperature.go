package temperature

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
