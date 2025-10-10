package main

import (
	"fmt"

	"github.com/Nevermind0911/task-2-1/internal/temp"
)

const (
	minBound = 15
	maxBound = 30
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("couldn't read number of departments")
	}

	for range n {
		var k int
		if _, err := fmt.Scan(&k); err != nil {
			fmt.Println("couldn't read number of employees")
		}
		leftBound, rightBound := minBound, maxBound
		for range k {
			var (
				val int
				op  string
			)
			if _, err := fmt.Scan(&op, &val); err != nil {
				fmt.Println("couldn't read temp")
			}

			leftBound, rightBound = temp.UpdateInterval(leftBound, rightBound, op, val)
			opt := temp.GetOptimal(leftBound, rightBound)
			fmt.Println(opt)
		}
	}
}
