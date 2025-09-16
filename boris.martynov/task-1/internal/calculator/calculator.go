package calc

import (
 "fmt"
 "os"
)

func Add(a, b int64) int64 {
 return a + b
}

func Subt(a, b int64) int64 {
 return a - b
}

func Mul(a, b int64) int64 {
 return a * b
}

func Div(a, b int64) float64 {
 if b == 0 {
  fmt.Println("Division by zero")
  return
 }
 return float64(a) / float64(b)
}

func UserInput() (int64, int64, string) {
 var a, b int64
 var operator string
 var n int
 var err error

 fmt.Print("First number: ")
 n, err = fmt.Scanln(&a)
 if err != nil || n != 1 {
  fmt.Println("Invalid first operand")
  return
 }

 fmt.Print("Second number: ")
 n, err = fmt.Scanln(&b)
 if err != nil || n != 1 {
  fmt.Println("Invalid second operand")
  return
 }

 fmt.Print("Operation: ")
 n, err = fmt.Scanln(&operator)
 if err != nil || n != 1 {
  fmt.Println("Invalid operation")
  return
 }

 if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
  fmt.Println("Invalid operation")
  return
 }

 return a, b, operator
}
