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

	_, err := fmt.Scanln(&numberDishes)
	if err != nil {
		fmt.Println("Error scan number of dishes:", err)

		return
	}

	for range numberDishes {
		var rating int
		fmt.Scan(&rating)
		heap.Push(&h, rating)
	}

	_, err = fmt.Scanln(&sequenceNumber)
	if err != nil {
		fmt.Println("Error scan sequence number:", err)

		return
	}

	for range sequenceNumber - 1 {
		heap.Pop(&h)
	}
	fmt.Println(heap.Pop(&h))
}
