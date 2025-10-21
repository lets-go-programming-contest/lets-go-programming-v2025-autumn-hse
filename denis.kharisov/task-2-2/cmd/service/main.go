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

	for range numberOfDishes {
		var dish int
		if _, err := fmt.Scan(&dish); err != nil {
			fmt.Println("Invalid dish number")

			return
		}

		heap.Push(heapOfDishes, dish)
	}

	if _, err := fmt.Scan(&preferredDishNumber); err != nil {
		fmt.Println("Invalid preferred dish number")
		return
	}

	for range preferredDishNumber - 1 {
		heap.Pop(heapOfDishes)
	}

	val := heap.Pop(heapOfDishes)

	if result, ok := val.(int); ok {
		fmt.Println(result)
	} else {
		fmt.Println("Favorite dish not found")
	}
}
