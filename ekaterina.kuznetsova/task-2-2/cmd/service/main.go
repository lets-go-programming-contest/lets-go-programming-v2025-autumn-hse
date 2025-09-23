package main

import (
	"container/heap"
	"fmt"

	"github.com/Ekaterina-101/task-2-2/internal/myheap"
)

func main() {
	var (
		numberDishes, sequenceNumber int
		h                            myheap.IntHeap
	)

	fmt.Scanln(&numberDishes)

	for range numberDishes {
		var rating int
		fmt.Scan(&rating)
		heap.Push(&h, rating)
	}

	fmt.Scanln(&sequenceNumber)
	for range sequenceNumber - 1 {
		heap.Pop(&h)
	}
	fmt.Println(heap.Pop(&h))
}
