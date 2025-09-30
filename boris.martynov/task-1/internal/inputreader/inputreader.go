package inputreader

import (
	"errors"

	"fmt"
)

var (
	errInvalidFirstOperand  = errors.New("Invalid first operand")
	errInvalidSecondOperand = errors.New("Invalid second operand")
)

func ReadNumber(name string) (int64, error) {
	var n int64
	_, err := fmt.Scanln(&n)
	if err != nil {
		if name == "first operand" {
			return 0, errInvalidFirstOperand
		}
		return 0, errInvalidSecondOperand
	}
	return n, nil
}

func ReadOperator() (string, error) {
	var (
		op                  string
		errInvalidOperation = errors.New("Invalid operation")
	)
	_, err := fmt.Scanln(&op)
	if err != nil {
		return "", errInvalidOperation
	}
	return op, nil
}
