package intheap

import (
	"container/heap"
	"errors"
)

var (
	ErrInvalidK          = errors.New("k must be positive")
	ErrNotEnoughElements = errors.New("not enough elements to find k-th maximum")
	ErrTypeAssertion     = errors.New("type assertion failed")
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

func KthMaximum(arr []int, kth int) (int, error) {
	if kth <= 0 {
		return 0, ErrInvalidK
	}

	if len(arr) < kth {
		return 0, ErrNotEnoughElements
	}

	IntHeap := InitIntHeap()
	for _, x := range arr {
		heap.Push(IntHeap, x)
	}

	var result int

	for i := range kth {
		val := heap.Pop(IntHeap)

		if i == kth-1 {
			rating, ok := val.(int)
			if !ok {
				return 0, ErrTypeAssertion
			}

			result = rating
		}
	}

	return result, nil
}
