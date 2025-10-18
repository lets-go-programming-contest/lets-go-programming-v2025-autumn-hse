package intheap

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		panic("Index out of bounds")
	}

	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		panic("Index out of bounds")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	rating, ok := x.(int)
	if !ok {
		panic("push into heap: expected int value")
	}

	*h = append(*h, rating)
}

func (h *IntHeap) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func InitIntHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)

	return h
}

func KthMaximum(arr []int, k int) (int, error) {
	if k <= 0 {
		return 0, fmt.Errorf("k must be positive")
	}

	if len(arr) < k {
		return 0, fmt.Errorf("not enough elements to find %d-th maximum", k)
	}

	h := InitIntHeap()
	for _, x := range arr {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := heap.Pop(h)
	if rating, ok := result.(int); !ok {
		return rating, nil
	}

	return 0, fmt.Errorf("type assertion failed")
}
