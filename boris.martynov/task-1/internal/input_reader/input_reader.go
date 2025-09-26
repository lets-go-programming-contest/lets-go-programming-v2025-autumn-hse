package input_reader

import (
	"errors"

	"fmt"
)

var (
	ErrDivisionByZero   = errors.New("Division by zero")
	ErrInvalidNumber    = errors.New("Invalid number")
	ErrInvalidOperation = errors.New("Invalid operation")
)

func ReadNumber(name string) (int64, error) {
	var n int64
	read, err := fmt.Scanln(&n)
	if err != nil || read != 1 {
		return 0, fmt.Errorf("%w: %s", ErrInvalidNumber, name)
	}
	return n, nil
}

func ReadOperator() (string, error) {
	var op string
	read, err := fmt.Scanln(&op)
	if err != nil || read != 1 {
		return "", ErrInvalidOperation
	}
	return op, nil
}
