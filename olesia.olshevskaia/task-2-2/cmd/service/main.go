package main

import (
	"container/heap"
	"fmt"

	"github.com/olesia.olshevsraia/task-2-2/intheap"
)

func popDesiredDish(h *intheap.IntHeap, desired int) (int, error) {
	if desired > h.Len() {
		return 0, fmt.Errorf("desired dish is greater than the number of dishes in the heap")
	}

	for range desired - 1 {
		val := heap.Pop(h)
		if val == nil {
			return 0, fmt.Errorf("Heap is empty")
		}
	}

	val := heap.Pop(h)
	if val == nil {
		return 0, fmt.Errorf("heap is empty")
	}

	count, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("invalid type from heap.Pop")
	}

	return count, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic occurred: %v\n", r)
		}
	}()

	myHeap := &intheap.IntHeap{}
	heap.Init(myHeap)

	var numberDishes int
	if _, err := fmt.Scanln(&numberDishes); err != nil {
		fmt.Println("error reading the number of dishes", err)

		return
	}

	var dish int
	for range numberDishes {
		if _, err := fmt.Scan(&dish); err != nil {
			fmt.Println("error reading dishf", err)

			return
		}

		heap.Push(myHeap, dish)
	}

	var desiredDish int
	if _, err := fmt.Scanln(&desiredDish); err != nil {
		fmt.Println("error reading priority", err)

		return
	}

	result, err := popDesiredDish(myHeap, desiredDish)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
