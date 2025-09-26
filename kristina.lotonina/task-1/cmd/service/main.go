package main

import (
	"fmt"
	"github.com/kef1rch1k/task-1/internal/calculator"
)

func main() {
	var (
		first, second int
		operation string
	)
	_, err := fmt.Scan(&first)
	if (err != nil) {
		fmt.Print("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&second)
	if (err != nil) {
		fmt.Print("Invalid second operand")
		return
	}
	_, err = fmt.Scan(&operation)
	if (err != nil) {
		fmt.Print("Invalid operation")
		return
	}
	calculator.Calculate(first, second, operation)
	return
}