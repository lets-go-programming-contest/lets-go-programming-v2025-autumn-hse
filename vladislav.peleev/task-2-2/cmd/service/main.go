package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
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
		panic("cannot pop from empty heap")
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

	for i := 0; i < dishCount; i++ {
		var dish int
		if _, err := fmt.Scan(&dish); err != nil {
			fmt.Println("failed to read dish:", err)
			return
		}
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

func findKthPreference(h *IntHeap, k int) (int, error) {
	if k > h.Len() || k <= 0 {
		return 0, errors.New("invalid preference order")
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	top := heap.Pop(h)
	num, ok := top.(int)
	if !ok {
		return 0, errors.New("heap returned non-integer value")
	}

	return num, nil
}
