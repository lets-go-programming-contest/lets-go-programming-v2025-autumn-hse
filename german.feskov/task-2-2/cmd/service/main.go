package main

import (
	"container/heap"
	"fmt"

	"github.com/6ermvH/task-2-2/internal/heapmax"
)

func main() {
	var (
		dishCount int
		kInd      int
	)

	dishes := &heapmax.IntHeap{}
	heap.Init(dishes)

	if _, err := fmt.Scan(&dishCount); err != nil {
		fmt.Println(err)

		return
	}

	for range dishCount {
		var val int
		if _, err := fmt.Scan(&val); err != nil {
			fmt.Println(err)

			return
		}

		heap.Push(dishes, val)
	}

	if _, err := fmt.Scan(&kInd); err != nil {
		fmt.Println(err)

		return
	}

	for range kInd - 1 {
		heap.Pop(dishes)
	}

	fmt.Println(heap.Pop(dishes))
}
