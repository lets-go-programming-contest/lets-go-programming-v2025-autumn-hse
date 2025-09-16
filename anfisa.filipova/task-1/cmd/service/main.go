package main

import (
	"calculator/pkg"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var operand1, operand2 string
	var err error
	var operation string
	var num1, num2 int
	fmt.Println("Enter first operand: ")
	fmt.Fscan(os.Stdin, &operand1)
	num1, err = strconv.Atoi(operand1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	fmt.Println("Enter second operand: ")
	fmt.Fscan(os.Stdin, &operand2)
	num2, err = strconv.Atoi(operand2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	fmt.Println("Enter operation: ")
	fmt.Fscan(os.Stdin, &operation)
	result, err := pkg.Calculate(num1, num2, operation)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result: ")
	fmt.Println(result)
}
