package main

import (
	"calculator/pkg"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var operand1, operand2 string
	var operation string
	var num1, num2 int
	//fmt.Println("Enter first operand: ")
	_, err := fmt.Fscan(os.Stdin, &operand1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading error: %v\n", err)
		os.Exit(1)
	}
	num1, err = strconv.Atoi(operand1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	//fmt.Println("Enter second operand: ")
	_, err = fmt.Fscan(os.Stdin, &operand2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading error: %v\n", err)
		os.Exit(1)
	}
	num2, err = strconv.Atoi(operand2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	//fmt.Println("Enter operation: ")
	_, err = fmt.Fscan(os.Stdin, &operation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading error: %v\n", err)
		os.Exit(1)
	}
	result, err := pkg.Calculate(num1, num2, operation)

	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Result: ")
	fmt.Println(result)
}
