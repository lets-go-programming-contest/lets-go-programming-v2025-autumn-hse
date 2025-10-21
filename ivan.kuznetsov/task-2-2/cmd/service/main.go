package main

import (
	"container/heap"
	"fmt"

	"github.com/kuzid-17/task-2-2/internal/intheap"
)

func main() {
	var dishesCount int

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Printf("Invalid number of dishes: %v\n", err)

		return
	}

	ratings := &intheap.Rating{}
	heap.Init(ratings)

	var ratingNumber int
	for range dishesCount {
		_, err = fmt.Scan(&ratingNumber)
		if err != nil {
			fmt.Printf("Invalid value of rating: %v\n", err)

			return
		}

		heap.Push(ratings, ratingNumber)
	}

	var preferenceIndex int

	_, err = fmt.Scan(&preferenceIndex)
	if err != nil {
		fmt.Printf("Invalid number of preference: %v\n", err)

		return
	}

	for range preferenceIndex - 1 {
		heap.Pop(ratings)
	}

	ratingValue := heap.Pop(ratings)

	resultValue, ok := ratingValue.(int)
	if !ok {
		panic("rating value is not an integer: empty or type mismatch")
	}

	fmt.Println(resultValue)
}
