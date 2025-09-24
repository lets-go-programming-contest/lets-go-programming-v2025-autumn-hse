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
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Push(x any) {
	newX, ok := x.(int)
	if !ok {
		fmt.Println("invalid type for Push")

		return
	}
	*h = append(*h, newX)
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
		ok                                     bool
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

	for range desiredDish {
		count, ok = heap.Pop(myHeap).(int)
		if !ok {
			fmt.Println("invalid type from heap.Pop")

			return
		}
	}

	if _, err := fmt.Println(count); err != nil {
		fmt.Println(err)

		return
	}
}
