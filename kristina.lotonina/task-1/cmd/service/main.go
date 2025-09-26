package main

import (
	"fmt"

	"github.com/kef1rch1k/task-1/internal/calculator"
)

func main() {
	var (
		first, second int
		operation     string
	)
	_, err1 := fmt.Scan(&first)
	_, err2 := fmt.Scan(&second)
	_, err3 := fmt.Scan(&operation)
	if err1 != nil {
		fmt.Println("Invalid first operand")
	}
	if err2 != nil {
		fmt.Println("Invalid second operand")
	}
	if err3 != nil {
		fmt.Println("Invalid operation")
	}
	calculator.Calculate(first, second, operation)
}
