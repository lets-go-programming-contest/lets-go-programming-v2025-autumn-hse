package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/6ermvH/task-2-2/internal/heapmax"
)

var errEmptyHeap = errors.New("heap is empty")

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
		fmt.Printf("while scan dish count: %v", err)

		return
	}

	for range dishCount {
		var val int
		if _, err := fmt.Scan(&val); err != nil {
			fmt.Printf("while scan dish value: %v", err)

			return
		}

		heap.Push(dishes, val)
	}

	var kInd int
	if _, err := fmt.Scan(&kInd); err != nil {
		fmt.Printf("while scan K num: %v", err)

		return
	}

	for range kInd - 1 {
		if heap.Pop(dishes) == nil {
			fmt.Printf("while pop in heap: %v", errEmptyHeap)

			return
		}
	}

	bestDish := heap.Pop(dishes)
	if bestDish == nil {
		fmt.Printf("while pop in heap: %v", errEmptyHeap)

		return
	}

	fmt.Println(heap.Pop(dishes))
}
