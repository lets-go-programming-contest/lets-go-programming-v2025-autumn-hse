package temperaturecontroller

import "errors"

const (
	minTemperature = 15
	maxTemperature = 30
)

var ErrBadRequest = errors.New("error: bad request")

type Request struct {
	Type string
	Val  int
}

func GetOptimalTemperatures(requests []Request) ([]int, error) {
	result := make([]int, 0)

	var (
		minT = minTemperature
		maxT = maxTemperature
	)

	for _, req := range requests {
		switch req.Type {
		case ">=":
			minT = maxInt(minT, req.Val)
		case "<=":
			maxT = minInt(maxT, req.Val)
		default:
			return result, ErrBadRequest
		}

		if minT > maxT {
			result = append(result, -1)
		} else {
			result = append(result, minT)
		}
	}

	return result, nil
}

func minInt(a, b int) int { // можно вынести в отдельный пакет, но это усложнит, решил оставить
	if a < b {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
