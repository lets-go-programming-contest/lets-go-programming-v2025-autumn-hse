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

	var (
		dishCount, dishRank, preferredDish int
		intHeap                            myheap.Heap
	)

	_, err := fmt.Scan(&dishCount)
	if err != nil {
		fmt.Println("Error reading dish count:", err)

		return
	}

	for range dishCount {
		_, err := fmt.Scan(&dishRank)
		if err != nil {
			fmt.Println("Error reading dish rank:", err)

			return
		}

		heap.Push(&intHeap, dishRank)
	}

	_, err = fmt.Scan(&preferredDish)
	if err != nil {
		fmt.Println("Error reading preferred dish:", err)

		return
	}

	if preferredDish > intHeap.Len() {
		fmt.Println("Too big preffered dish")

		return
	}

	for range intHeap.Len() - preferredDish {
		heap.Pop(&intHeap)
	}

	fmt.Println(heap.Pop(&intHeap))
}
