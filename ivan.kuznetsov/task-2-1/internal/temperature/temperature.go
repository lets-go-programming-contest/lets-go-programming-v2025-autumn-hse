package temperature

func OptimalTemperature(sign string, value int, numbers []int) []int {
	var result []int

	for _, temp := range numbers {
		switch sign {
		case "<=":
			if temp <= value {
				result = append(result, temp)
			}
		case ">=":
			if temp >= value {
				result = append(result, temp)
			}
		}
	}

	return result
}
