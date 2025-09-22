package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (heap IntHeap) Len() int {
	return len(heap)
}

func (heap IntHeap) Less(i, j int) bool {
	return heap[i] > heap[j]
}

func (heap IntHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *IntHeap) Push(x interface{}) {
	*heap = append(*heap, x.(int))
}

func (heap *IntHeap) Pop() interface{} {
	oldHeap := *heap
	x := oldHeap[len(oldHeap)-1]
	*heap = oldHeap[0 : len(oldHeap)-1]
	return x
}

func main() {
	var (
		dishCount, dish, rank, ans int
	)

	_, err := fmt.Scanln(&dishCount)
	if err != nil {
		return
	}

	h := &IntHeap{}
	heap.Init(h)

	for range dishCount {
		_, err := fmt.Scan(&dish)
		if err != nil {
			return
		}
		heap.Push(h, dish)
	}

	_, err = fmt.Scanln(&rank)
	if err != nil {
		return
	}

	for range rank {
		ans = heap.Pop(h).(int)
	}

	_, err = fmt.Println(ans)
	if err != nil {
		return
	}
}
