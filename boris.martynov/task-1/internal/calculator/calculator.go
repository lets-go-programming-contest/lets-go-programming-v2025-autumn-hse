package calculator

import (
	"fmt"
)

func Div(a, b int64) (int64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

func ReadNumber(name string) (int64, error) {
	var n int64
	read, err := fmt.Scanln(&n)
	if err != nil || read != 1 {
		return 0, fmt.Errorf("Invalid %s", name)
	}
	return n, nil
}

func ReadOperator() (string, error) {
	var op string
	read, err := fmt.Scanln(&op)
	if err != nil || read != 1 {
		return "", fmt.Errorf("Invalid operation")
	}
	return op, nil
}
