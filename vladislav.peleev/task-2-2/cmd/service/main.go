package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	ErrInvalidPreferenceOrder = errors.New("invalid preference order")
	ErrHeapReturnedNonInt     = errors.New("heap returned non-integer value")
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if i < 0 || j < 0 || len(*h) < i || len(*h) < j {
		panic("index out of range")
	}

	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	num, ok := x.(int)
	if !ok {
		panic("invalid type pushed to heap, expected int")
	}

	*h = append(*h, num)
}

func (h *IntHeap) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]

	*h = old[:n-1]

	return x
}

func main() {
	var dishCount int

	if _, err := fmt.Scan(&dishCount); err != nil {
		fmt.Println("failed to read dish count:", err)

		return
	}

	intHeap := &IntHeap{}
	heap.Init(intHeap)

	dishes := make([]int, dishCount)
	for index := range dishes {
		if _, err := fmt.Scan(&dishes[index]); err != nil {
			fmt.Println("failed to read dish:", err)

			return
		}
	}

	for _, dish := range dishes {
		heap.Push(intHeap, dish)
	}

	var preferenceOrder int

	if _, err := fmt.Scan(&preferenceOrder); err != nil {
		fmt.Println("failed to read preference order:", err)

		return
	}

	result, err := findKthPreference(intHeap, preferenceOrder)
	if err != nil {
		fmt.Println("error:", err)

		return
	}

	fmt.Println(result)
}

func findKthPreference(dishHeap *IntHeap, k int) (int, error) {
	if k <= 0 || k > dishHeap.Len() {
		return 0, ErrInvalidPreferenceOrder
	}

	for range make([]struct{}, k-1) {
		heap.Pop(dishHeap)
	}

	top := heap.Pop(dishHeap)

	num, ok := top.(int)
	if !ok {
		return 0, ErrHeapReturnedNonInt
	}

	return num, nil
}
