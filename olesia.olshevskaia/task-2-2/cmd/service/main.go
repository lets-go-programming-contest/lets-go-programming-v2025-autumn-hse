package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *intHeap) Swap(i, j int) {
	tmp := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = tmp
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func main() {
	var (
		numberDishes, desiredDish, dish, count int
	)

	myHeap := &intHeap{}
	heap.Init(myHeap)

	if _, err := fmt.Scanln(&numberDishes); err != nil {
		fmt.Println(err)

		return
	}

	for range numberDishes {
		if _, err := fmt.Scan(&dish); err != nil {
			fmt.Println(err)

			return
		}
		heap.Push(myHeap, dish)
	}

	if _, err := fmt.Scanln(&desiredDish); err != nil {
		fmt.Println(err)

		return
	}

	for i := 0; i < desiredDish; i++ {
		count = heap.Pop(myHeap).(int)
	}

	fmt.Println(count)
}
