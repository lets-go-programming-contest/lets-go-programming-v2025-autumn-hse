package calculator

import (
	"fmt"
)

var (
	operators = map[byte]struct{}{'+': {},
		'-': {},
		'*': {},
		'/': {}}
)

type Operand struct {
	Value int
}

func (o *Operand) Read() error {
	_, err := fmt.Scan(&o.Value)
	if err != nil {
		return fmt.Errorf("Invalid operand")
	}
	return nil
}

type Operation struct {
	Value byte
}

func (o *Operation) Read() error {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		return fmt.Errorf("Invalid Operation")
	}
	if len(str) != 1 {
		return fmt.Errorf("Invalid Operation")
	}
	o.Value = str[0]
	if _, ok := operators[o.Value]; !ok {
		return fmt.Errorf("Invalid Operation")
	}
	return nil
}

/// Operations

func Addition(a, b Operand) (Operand, error) {
	result := Operand{Value: a.Value + b.Value}
	return result, nil
}

func Subtraction(a, b Operand) (Operand, error) {
	result := Operand{Value: a.Value - b.Value}
	return result, nil
}

func Multiplication(a, b Operand) (Operand, error) {
	result := Operand{Value: a.Value * b.Value}
	return result, nil
}

func Division(a, b Operand) (Operand, error) {
	if b.Value == 0 {
		return b, fmt.Errorf("Division by zero")
	}
	result := Operand{Value: a.Value / b.Value}
	return result, nil
}

func Calculate(a, b Operand, op Operation) (Operand, error) {
	var result Operand
	var err error
	switch op.Value {
	case '+':
		result, err = Addition(a, b)
	case '-':
		result, err = Subtraction(a, b)
	case '*':
		result, err = Multiplication(a, b)
	case '/':
		result, err = Division(a, b)
	}
	return result, err
}
