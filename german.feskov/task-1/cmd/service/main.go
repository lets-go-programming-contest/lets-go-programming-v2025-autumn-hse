package main

import (
	"fmt"

	calc "github.com/6ermvh/calculator/internal/calculator"
)

func main() {
	var (
		a, b calc.Operand
		op   calc.Operation
	)

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err := fmt.Scan(&op); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	res, err := calc.Calculate(a, b, op)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
