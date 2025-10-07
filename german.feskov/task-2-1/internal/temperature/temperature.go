package temperature

import (
	"errors"
	"fmt"
)

var ErrBadRequestType = errors.New("uncorrect request type")

const (
	minTemperature = 15
	maxTemperature = 30
)

type RequestType string

const (
	minRequestType RequestType = "<="
	maxRequestType RequestType = ">="
)

type Request struct {
	Type RequestType
	Val  int
}

func Calculate(outRequests <-chan Request, inCalculated chan<- int) error {
	var (
		minT = minTemperature
		maxT = maxTemperature
		err  error
	)

	for req := range outRequests {
		minT, maxT, err = calculate(minT, maxT, req)
		if err != nil {
			return err
		}

		if minT > maxT {
			inCalculated <- -1

			continue
		}

		inCalculated <- minT
	}

	return err
}

func calculate(minT, maxT int, req Request) (int, int, error) {
	switch req.Type {
	case minRequestType:
		return minT, min(maxT, req.Val), nil
	case maxRequestType:
		return max(minT, req.Val), maxT, nil
	default:
		return minT, maxT, fmt.Errorf("%w: %q", ErrBadRequestType, req.Type)
	}
}
