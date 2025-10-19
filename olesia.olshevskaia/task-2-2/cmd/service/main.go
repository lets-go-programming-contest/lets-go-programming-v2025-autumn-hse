package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/olesia.olshevsraia/task-2-2/internal/intheap"
)

var (
	ErrDesiredDishTooBig = errors.New("desired dish is greater than the number of dishes in the heap")
	ErrHeapEmpty         = errors.New("heap is empty")
	ErrInvalidType       = errors.New("invalid type from heap.Pop")
)

func popDesiredDish(heapInt intheap.IntHeap, desired int) (int, error) {
	if desired > (&heapInt).Len() {
		return 0, ErrDesiredDishTooBig
	}

	for range desired - 1 {
		heap.Pop(&heapInt)
	}

	val := heap.Pop(&heapInt)
	count, ok := val.(int)

	if !ok {
		return 0, ErrInvalidType
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

	result, err := popDesiredDish(*myHeap, desiredDish)
	if err != nil {
		fmt.Println("failed to get dish with given id:", err)

		return
	}

	fmt.Println(result)
}
