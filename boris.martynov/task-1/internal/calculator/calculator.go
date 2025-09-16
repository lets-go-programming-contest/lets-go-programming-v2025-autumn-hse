package calc

import (
	"fmt"
)

func Add(a, b int64) int64 {
	return a + b
}

func Subt(a, b int64) int64 {
	return a - b
}

func Mul(a, b int64) int64 {
	return a * b
}

func Div(a, b int64) (int64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func UserInput() (int64, int64, string) {
	var a, b int64
	var operator string
	var n int
	var err error

	n, err = fmt.Scanln(&a)
	if err != nil || n != 1 {
		fmt.Println("Invalid first operand")
		return 0, 0, ""
	}

	n, err = fmt.Scanln(&b)
	if err != nil || n != 1 {
		fmt.Println("Invalid second operand")
		return 0, 0, ""
	}

	n, err = fmt.Scanln(&operator)
	if err != nil || n != 1 {
		fmt.Println("Invalid operation")
		return 0, 0, ""
	}

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Invalid operation")
		return 0, 0, ""
	}

	return a, b, operator
}
