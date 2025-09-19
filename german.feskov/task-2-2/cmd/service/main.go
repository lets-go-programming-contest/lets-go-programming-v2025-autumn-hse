package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	v, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, v)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	var (
		dishCount int
		dishes    = &IntHeap{}
		kInd      int
	)

	heap.Init(dishes)

	if _, err := fmt.Scan(&dishCount); err != nil {
		return
	}

	for range dishCount {
		var val int
		if _, err := fmt.Scan(&val); err != nil {
			return
		}

		heap.Push(dishes, val)
	}

	if _, err := fmt.Scan(&kInd); err != nil {
		return
	}

	for range kInd - 1 {
		heap.Pop(dishes)
	}

	fmt.Println(heap.Pop(dishes))
}
