package main

import (
	"container/heap"
	"fmt"

	"github.com/kamilSharipov/task-2-2/internal/intheap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic occurred: %v\n", r)
		}
	}()

	var dishCount int

	_, err := fmt.Scanln(&dishCount)
	if err != nil {
		fmt.Printf("Error scanning dishCount: %v\n", err)

		return
	}

	heapInstance := intheap.InitIntHeap()

	var dish int
	for range dishCount {
		_, err := fmt.Scan(&dish)
		if err != nil {
			fmt.Printf("Error scanning dish: %v\n", err)

			return
		}

		heap.Push(heapInstance, dish)
	}

	var rank int

	_, err = fmt.Scanln(&rank)
	if err != nil {
		fmt.Printf("Error scanning rank: %v\n", err)

		return
	}

	if heapInstance.Len() < rank {
		fmt.Println("Not enough elements in heap")

		return
	}

	var ans int
	for i := 0; i < rank; i++ {
		result := heap.Pop(heapInstance)
		if i == rank-1 {
			if rating, ok := result.(int); ok {
				ans = rating
			} else {
				fmt.Println("Unexpected type in heap")

				return
			}
		}
	}

	fmt.Println(ans)
}
