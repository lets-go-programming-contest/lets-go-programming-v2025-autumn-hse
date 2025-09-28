package myheap

import "errors"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}

func (h *IntHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		panic(errors.New("cast error"))
	}

	*h = append(*h, value)
}

func (h *IntHeap) Pop() any {
	if len(*h) == 0 {
		panic(errors.New("data is empty"))
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
