package main

import (
	"fmt"

	"github.com/kamilSharipov/task-2-2/internal/intheap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic occurred: %v\n", r)
		}
	}()

	var dishCount int

	_, err := fmt.Scanln(&dishCount)
	if err != nil {
		fmt.Printf("Error scanning dishCount: %v\n", err)

		return
	}

	dishes := make([]int, dishCount)

	for i := range dishes {
		_, err := fmt.Scan(&dishes[i])
		if err != nil {
			fmt.Printf("Error scanning dish: %v\n", err)

			return
		}
	}

	var rank int

	_, err = fmt.Scanln(&rank)
	if err != nil {
		fmt.Printf("Error scanning rank: %v\n", err)

		return
	}

	ans, err := intheap.KthMaximum(dishes, rank)
	if err != nil {
		fmt.Printf("Error to find %d-th maximum: %v\n", rank, err)

		return
	}

	fmt.Println(ans)
}
