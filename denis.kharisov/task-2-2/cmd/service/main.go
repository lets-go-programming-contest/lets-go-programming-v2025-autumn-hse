package main

import (
	"container/heap"
	"fmt"

	"github.com/denisK-H/task-2-2/internal/maxheap"
)

func main() {
	var numberOfDishes, preferredDishNumber int
	_, err := fmt.Scan(&numberOfDishes)

	if err != nil {
		fmt.Println("Invalid number of dishes")
	}

	dishes := make([]int, numberOfDishes)

	for i := range dishes {
		_, err = fmt.Scan(&dishes[i])
		if err != nil {
			fmt.Println("Invalid dish number")
		}
	}

	_, err = fmt.Scan(&preferredDishNumber)

	if err != nil {
		fmt.Println("Invalid preferred dish number")
	}

	heapOfDishes := &maxheap.Maxheap{}
	heap.Init(heapOfDishes)

	for _, dish := range dishes {
		heap.Push(heapOfDishes, dish)
	}

	var result int

	for range preferredDishNumber {
		result = heap.Pop(heapOfDishes).(int)
	}

	fmt.Println(result)
}
