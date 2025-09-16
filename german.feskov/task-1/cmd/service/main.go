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
		fmt.Println(fmt.Errorf("Invalid first operand"))
		os.Exit(0)
	}
	if err := B.Read(); err != nil {
		fmt.Println(fmt.Errorf("Invalid second operand"))
		os.Exit(0)
	}
	if err := Op.Read(); err != nil {
		fmt.Println(err)
		os.Exit(0)
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
			fmt.Println(err)
			os.Exit(0)
		}
	}

	fmt.Println(Res.Value)
}
