package main

import (
	"calculator/pkg/mycalc"
	"fmt"
)

func main() {
	var (
		num1, num2 int
		operation  string
	)
	_, err := fmt.Scan(&num1)

	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&num2)

	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, _ = fmt.Scan(&operation)

	result, err := mycalc.Calculate(num1, num2, operation)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
