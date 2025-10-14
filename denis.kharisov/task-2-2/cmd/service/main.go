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

		return
	}

	dishes := make([]int, numberOfDishes)

	for i := range dishes {
		_, err = fmt.Scan(&dishes[i])
		if err != nil {
			fmt.Println("Invalid dish number")

			return
		}
	}

	_, err = fmt.Scan(&preferredDishNumber)
	if err != nil {
		fmt.Println("Invalid preferred dish number")

		return
	}

	heapOfDishes := &maxheap.Maxheap{}
	heap.Init(heapOfDishes)

	for _, dish := range dishes {
		heap.Push(heapOfDishes, dish)
	}

	var result int

	for range preferredDishNumber {
		if res, flag := heap.Pop(heapOfDishes).(int); flag {
			result = res
		} else {
			panic("unexpected type from heap")
		}
	}

	fmt.Println(result)
}
