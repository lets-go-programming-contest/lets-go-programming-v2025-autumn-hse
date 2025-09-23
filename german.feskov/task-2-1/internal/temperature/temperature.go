package temperature

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrBadRequestType = errors.New("uncorrect request type")
	ErrFailRead       = errors.New("fail to read")
	ErrFailWrite      = errors.New("fail to write")
)

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

func Read(in io.Reader, inRequests chan<- Request, count int) error {
	defer close(inRequests)

	for range count {
		var req Request
		if _, err := fmt.Fscanf(in, "\n%s %d", &req.Type, &req.Val); err != nil {
			return fmt.Errorf("%w: %w", ErrFailRead, err)
		}

		inRequests <- req
	}

	return nil
}

func Calculate(outRequests <-chan Request, inCalculated chan<- int) error {
	defer close(inCalculated)

	var (
		minT       = minTemperature
		maxT       = maxTemperature
		err  error = nil
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

func Write(out io.Writer, outCalculated <-chan int) error {
	for res := range outCalculated {
		if _, err := fmt.Fprintln(out, res); err != nil {
			return fmt.Errorf("%w: %w", ErrFailWrite, err)
		}
	}

	return nil
}

func calculate(minT, maxT int, req Request) (int, int, error) {
	switch req.Type {
	case minRequestType:
		return minT, minInt(maxT, req.Val), nil
	case maxRequestType:
		return maxInt(minT, req.Val), maxT, nil
	default:
		return minT, maxT, fmt.Errorf("%w %s", ErrBadRequestType, req.Type)
	}
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
