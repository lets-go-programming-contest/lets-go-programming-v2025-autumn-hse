package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/kamilSharipov/task-2-2/internal/intheap"
)

var (
	errKPositive        error = errors.New("k must be positive")
	errNotEnoughEls     error = errors.New("not enough elements to find k-th maximum")
	errHeapBecameEmpty  error = errors.New("heap became empty before k-th element")
	errHeapReturnedNill error = errors.New("heap returned nil")
	errTypeAssertFailed error = errors.New("type assertion failed")
)

func KthMaximum(arr []int, kth int) (int, error) {
	if kth <= 0 {
		return 0, errKPositive
	}

	if len(arr) < kth {
		return 0, errNotEnoughEls
	}

	IntHeap := intheap.InitIntHeap()
	for _, x := range arr {
		heap.Push(IntHeap, x)
	}

	var result int

	for range kth - 1 {
		val := heap.Pop(IntHeap)
		if val == nil {
			return 0, errHeapBecameEmpty
		}
	}

	val := heap.Pop(IntHeap)
	if val == nil {
		return 0, errHeapReturnedNill
	}

	result, ok := val.(int)
	if !ok {
		return 0, errTypeAssertFailed
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
