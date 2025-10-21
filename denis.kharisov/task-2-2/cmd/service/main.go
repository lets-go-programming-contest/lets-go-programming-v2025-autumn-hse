package main

import (
	"container/heap"
	"fmt"

	"github.com/denisK-H/task-2-2/internal/maxheap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Runtime error:", r)
		}
	}()

	var numberOfDishes, preferredDishNumber int

	_, err := fmt.Scan(&numberOfDishes)
	if err != nil {
		fmt.Println("Invalid number of dishes")

		return
	}

	heapOfDishes := &maxheap.Maxheap{}
	heap.Init(heapOfDishes)

	for i := 0; i < numberOfDishes; i++ {
		var dish int
		if _, err := fmt.Scan(&dish); err != nil {
			fmt.Println("Invalid dish number")

			return
		}

		heap.Push(heapOfDishes, dish)
	}

	var result int

	for range preferredDishNumber - 1 {
		heap.Pop(heapOfDishes)
	}

	val := heap.Pop(heapOfDishes)
	if val == nil {
		fmt.Println("Favorite dish not found")
	}

	result = val.(int)
	fmt.Println(result)
}
