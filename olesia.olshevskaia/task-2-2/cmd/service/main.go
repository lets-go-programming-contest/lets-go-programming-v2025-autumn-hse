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
	newX, isInt := x.(int)
	if !isInt {
		panic("intHeap.Push: invalid type, expected int")
	}

	*h = append(*h, newX)
}

func (h *intHeap) Pop() any {
	old := *h
	length := len(*h)

	if length == 0 {
		fmt.Println("intHeap is empty")

		return nil
	}

	x := old[length-1]
	*h = old[:length-1]

	return x
}

func main() {
	var (
		numberDishes, desiredDish, dish, count int
		typeCastingOk                          bool
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
		x := heap.Pop(myHeap)

		count, typeCastingOk = x.(int)

		if !typeCastingOk {
			fmt.Println("invalid type from heap.Pop")

			return
		}
	}

	if _, err := fmt.Println(count); err != nil {
		fmt.Println(err)

		return
	}
}
