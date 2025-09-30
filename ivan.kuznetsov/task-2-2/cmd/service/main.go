package main

import (
	"container/heap"
	"fmt"

	"github.com/kuzid-17/task-2-2/internal/heapinterface"
)

func main() {
	var dishesCount, ratingNumber, preferenceIndex, result int

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Println("Invalid number of dishes")

		return
	}

	ratings := &heapinterface.Rating{}
	heap.Init(ratings)

	for range dishesCount {
		_, err = fmt.Scan(&ratingNumber)
		if err != nil {
			fmt.Println("Invalid value of rating")

			return
		}

		heap.Push(ratings, ratingNumber)
	}

	_, err = fmt.Scan(&preferenceIndex)
	if err != nil {
		fmt.Println("Invalid number of preference")

		return
	}

	for range preferenceIndex {
		value, ok := heap.Pop(ratings).(int)
		if ok {
			result = value
		}
	}

	fmt.Println(result)
}
