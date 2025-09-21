package main

import (
	"fmt"

	"github.com/6ermvh/calculator/internal/calculator"
)

func main() {
	var (
		a, b int
		op   string
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

	res, err := calculator.Calculate(a, b, op)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
