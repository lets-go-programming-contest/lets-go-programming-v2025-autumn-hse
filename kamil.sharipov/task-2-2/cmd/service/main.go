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

			return
		}
	}()

	var dishCount, dish, rank, ans int

	_, err := fmt.Scanln(&dishCount)
	if err != nil {
		return
	}

	heapInstance := intheap.InitIntHeap()

	for range dishCount {
		_, err := fmt.Scan(&dish)
		if err != nil {
			fmt.Printf("Error scanning dish: %v\n", err)

			return
		}

		heap.Push(heapInstance, dish)
	}

	_, err = fmt.Scanln(&rank)
	if err != nil {
		fmt.Printf("Error scanning rank: %v\n", err)

		return
	}

	for range rank {
		result := heap.Pop(heapInstance)
		if rating, ok := result.(int); ok {
			ans = rating
		} else {
			return
		}
	}

	_, err = fmt.Println(ans)
	if err != nil {
		return
	}
}
