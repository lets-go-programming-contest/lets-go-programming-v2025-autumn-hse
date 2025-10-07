package main

import (
	"container/heap"
	"fmt"

	"github.com/6ermvH/task-2-2/internal/heapmax"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	var dishCount int

	dishes := &heapmax.IntHeap{}
	heap.Init(dishes)

	if _, err := fmt.Scan(&dishCount); err != nil {
		fmt.Printf("err: while scan dish count: %v\n", err)

		return
	}

	for range dishCount {
		var val int
		if _, err := fmt.Scan(&val); err != nil {
			fmt.Printf("err: while scan dish value: %v\n", err)

			return
		}

		heap.Push(dishes, val)
	}

	var kInd int
	if _, err := fmt.Scan(&kInd); err != nil {
		fmt.Printf("err: while scan K num: %v\n", err)

		return
	}

	for range kInd - 1 {
		if heap.Pop(dishes) == nil {
			fmt.Printf("err: heap is empty")

			return
		}
	}

	bestDish := heap.Pop(dishes)
	if bestDish == nil {
		fmt.Printf("err: heap is empty")

		return
	}

	fmt.Println(bestDish)
}
