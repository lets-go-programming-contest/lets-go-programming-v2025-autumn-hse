package temperature

func OptimalTemperature(limit string, temperatureValue int, numbers []int) []int {
	var result []int

	for _, temp := range numbers {
		switch limit {
		case "<=":
			if temp <= temperatureValue {
				result = append(result, temp)
			}
		case ">=":
			if temp >= temperatureValue {
				result = append(result, temp)
			}
		}
	}

	return result
}
