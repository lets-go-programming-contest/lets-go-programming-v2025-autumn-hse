package intheap

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

//nolint:recvcheck
func (h *IntHeap) Push(x interface{}) {
	rating, ok := x.(int)
	if !ok {
		panic("IntHeap.Push: expected int")
	}

	*h = append(*h, rating)
}

//nolint:recvcheck
func (h *IntHeap) Pop() interface{} {
	if h.Len() == 0 {
		panic("IntHeap.Pop: cannot pop from empty heap")
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
