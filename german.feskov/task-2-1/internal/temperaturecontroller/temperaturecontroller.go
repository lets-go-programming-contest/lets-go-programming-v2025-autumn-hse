package temperaturecontroller

import (
	"errors"
	"fmt"
)

var errBadRequestType = errors.New("error: bad request type")

const (
	minTemperature = 15
	maxTemperature = 30

	requestTypeLess  = "<="
	requestTypeGreat = ">="
)

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

	for ind, req := range requests {
		switch req.Type {
		case requestTypeGreat:
			minT = maxInt(minT, req.Val)
		case requestTypeLess:
			maxT = minInt(maxT, req.Val)
		default:
			return result, fmt.Errorf("%w: %s, in %d request", errBadRequestType, req.Type, ind)
		}

		if minT > maxT {
			result = append(result, -1)
		} else {
			result = append(result, minT)
		}
	}

	return result, nil
}

func minInt(a, b int) int {
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
