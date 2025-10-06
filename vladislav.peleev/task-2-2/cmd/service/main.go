package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	if num, ok := x.(int); ok {
		*h = append(*h, num)
	}
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	
	return x
}

func main() {
	var dishCount, preferenceOrder int
	_, _ = fmt.Scan(&dishCount)

	dishes := make([]int, dishCount)
	for idx := range dishes {
		_, _ = fmt.Scan(&dishes[idx])
	}

	_, _ = fmt.Scan(&preferenceOrder)

	result := findKthPreference(dishes, preferenceOrder)
	fmt.Println(result)
}

func findKthPreference(dishes []int, preferenceOrder int) int {
	intHeap := &IntHeap{}
	heap.Init(intHeap)

	for _, dish := range dishes {
		heap.Push(intHeap, dish)
	}

	for idx := 0; idx < preferenceOrder-1; idx++ {
		heap.Pop(intHeap)
	}

	poppedValue := heap.Pop(intHeap)
	if result, typeOk := poppedValue.(int); typeOk {
		return result
	}
	
	return 0
}
