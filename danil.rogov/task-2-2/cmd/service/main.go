package main

import (
	"container/heap"
	"fmt"

	myheap "github.com/Tapochek2894/task-2/subtask-2/internal/heap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()

	var dishCount, dishRank int

	_, err := fmt.Scan(&dishCount)
	if err != nil {
		fmt.Println("Error reading dish count:", err)

		return
	}

	var intHeap myheap.Heap

	for range dishCount {
		_, err := fmt.Scan(&dishRank)
		if err != nil {
			fmt.Println("Error reading dish rank:", err)

			return
		}

		heap.Push(&intHeap, dishRank)
	}

	var preferredDish int

	_, err = fmt.Scan(&preferredDish)
	if err != nil {
		fmt.Println("Error reading preferred dish:", err)

		return
	}

	if preferredDish > intHeap.Len() || preferredDish <= 0 {
		fmt.Println("Invalid preferred dish")

		return
	}

	for range intHeap.Len() - preferredDish {
		heap.Pop(&intHeap)
	}

	value := heap.Pop(&intHeap)

	if value == nil {
		fmt.Println("Something is wrong")

		return
	}

	fmt.Println(value)
}
