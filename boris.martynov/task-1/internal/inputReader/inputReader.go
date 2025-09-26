package inputReader

import (
	"errors"

	"fmt"
)

func ReadNumber(name string) (int64, error) {
	var n int64
	_, err := fmt.Scanln(&n)
	if err != nil {
		if name == "first operand" {
			return 0, errors.New("Invalid first operand")
		} else {
			return 0, errors.New("Invalid second operand")
		}
	}
	return n, nil
}

func ReadOperator() (string, error) {
	var op string
	_, err := fmt.Scanln(&op)
	if err != nil {
		return "", errors.New("Invalid operation")
	}
	return op, nil
}
