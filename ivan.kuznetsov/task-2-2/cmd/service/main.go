package main

import (
	"container/heap"
	"fmt"

	intHeap "github.com/kuzid-17/task-2-2/internal/int-heap"
)

func main() {
	var dishesCount, ratingNumber, preferenceIndex, result int

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Printf("Invalid number of dishes: %v\n", err)

		return
	}

	ratings := &intHeap.Rating{}
	heap.Init(ratings)

	for range dishesCount {
		_, err = fmt.Scan(&ratingNumber)
		if err != nil {
			fmt.Printf("Invalid value of rating: %v\n", err)

			return
		}

		heap.Push(ratings, ratingNumber)
	}

	_, err = fmt.Scan(&preferenceIndex)
	if err != nil {
		fmt.Printf("Invalid number of preference: %v\n", err)

		return
	}

	for range preferenceIndex - 1 {
		heap.Pop(ratings)
	}
	result, ok := heap.Pop(ratings).(int)
	if !ok {
		panic("type conversion error")
	}
	fmt.Println(result)
}
