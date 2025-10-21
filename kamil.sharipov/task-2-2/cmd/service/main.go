package main

import (
	"container/heap"
	"fmt"

	"github.com/kamilSharipov/task-2-2/internal/intheap"
)

func KthMaximum(arr []int, kth int) (int, error) {
	if kth <= 0 {
		return 0, fmt.Errorf("k must be positive")
	}

	if len(arr) < kth {
		return 0, fmt.Errorf("not enough elements to find k-th maximum")
	}

	IntHeap := intheap.InitIntHeap()
	for _, x := range arr {
		heap.Push(IntHeap, x)
	}

	var result int

	for i := 0; i < kth-1; i++ {
		val := heap.Pop(IntHeap)
		if val == nil {
			return 0, fmt.Errorf("heap became empty before k-th element")
		}
	}

	val := heap.Pop(IntHeap)
	if val == nil {
		return 0, fmt.Errorf("heap returned nil")
	}

	result, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("type assertion failed")
	}

	return result, nil
}

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

	dishes := make([]int, dishCount)

	for i := range dishes {
		_, err := fmt.Scan(&dishes[i])
		if err != nil {
			fmt.Printf("Error scanning dish: %v\n", err)

			return
		}
	}

	var rank int

	_, err = fmt.Scanln(&rank)
	if err != nil {
		fmt.Printf("Error scanning rank: %v\n", err)

		return
	}

	ans, err := KthMaximum(dishes, rank)
	if err != nil {
		fmt.Printf("Error to find %d-th maximum: %v\n", rank, err)

		return
	}

	fmt.Println(ans)
}
