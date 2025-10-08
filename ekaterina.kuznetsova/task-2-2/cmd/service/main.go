package main

import (
	"container/heap"
	"fmt"

	"github.com/Ekaterina-101/task-2-2/internal/myheap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic in main: %v\n", r)
		}
	}()

	var numberDishes int

	_, err := fmt.Scanln(&numberDishes)
	if err != nil {
		fmt.Println("Error scan number of dishes:", err)

		return
	}

	var heapRatings myheap.IntHeap

	for range numberDishes {
		var rating int

		_, err = fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Error scan rating:", err)

			return
		}

		heap.Push(&heapRatings, rating)
	}

	var sequenceNumber int

	_, err = fmt.Scanln(&sequenceNumber)
	if err != nil {
		fmt.Println("Error scan sequence number:", err)

		return
	}

	if sequenceNumber > heapRatings.Len() {
		fmt.Println("Too few Dishes")

		return
	}

	for range sequenceNumber - 1 {
		if nil == heap.Pop(&heapRatings) {
			fmt.Println("heap is empty")

			return
		}
	}

	fmt.Println(heap.Pop(&heapRatings))
}
