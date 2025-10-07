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
	for dishNumber := 0; dishNumber < numberOfDishes; dishNumber++ {
		_, err = fmt.Scan(&dishes[dishNumber])
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
		heapOfDishes.Push(dish)
	}

	var result int
	for i := 0; i < preferredDishNumber; i++ {
		result = heapOfDishes.Pop().(int)
	}

	fmt.Println(&result)
}
