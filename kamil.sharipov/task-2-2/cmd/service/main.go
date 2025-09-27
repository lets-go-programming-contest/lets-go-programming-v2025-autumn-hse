package main

import (
	"container/heap"
	"fmt"

	"github.com/kamilSharipov/task-2-2/internal/intHeap"
)

func main() {
	var dishCount, dish, rank, ans int

	_, err := fmt.Scanln(&dishCount)
	if err != nil {
		return
	}

	heapInstance := intHeap.InitIntHeap()

	for range dishCount {
		_, err := fmt.Scan(&dish)
		if err != nil {
			return
		}

		heap.Push(heapInstance, dish)
	}

	_, err = fmt.Scanln(&rank)
	if err != nil {
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
