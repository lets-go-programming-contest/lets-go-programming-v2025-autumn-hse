package main

import (
	"container/heap"
	"fmt"

	"github.com/Ekaterina-101/task-2-2/internal/myheap"
)

func main() {
	var (
		numberDishes, sequenceNumber int
		heapRatings                  myheap.IntHeap
	)

	_, err := fmt.Scanln(&numberDishes)
	if err != nil {
		fmt.Println("Error scan number of dishes:", err)

		return
	}

	for range numberDishes {
		var rating int

		_, err = fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Error scan rating:", err)

			return
		}

		heap.Push(&heapRatings, rating)
	}

	_, err = fmt.Scanln(&sequenceNumber)
	if err != nil {
		fmt.Println("Error scan sequence number:", err)

		return
	}

	for range sequenceNumber - 1 {
		heap.Pop(&heapRatings)
	}

	fmt.Println(heap.Pop(&heapRatings))
}
