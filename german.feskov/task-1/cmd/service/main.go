package main

import (
	calc "calculator/internal/calculator"
	"fmt"
	"os"
)

func main() {

	var A, B, Res calc.Operand
	var Op calc.Operation

	if err := A.Read(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("Invalid first operand"))
		os.Exit(1)
	}
	if err := B.Read(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("Invalid second operand"))
		os.Exit(1)
	}
	if err := Op.Read(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch Op.Value {
	case '+':
		Res, _ = calc.Addition(A, B)
	case '-':
		Res, _ = calc.Subtraction(A, B)
	case '*':
		Res, _ = calc.Multiplication(A, B)
	case '/':
		var err error
		Res, err = calc.Division(A, B)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	fmt.Println(Res.Value)

}
