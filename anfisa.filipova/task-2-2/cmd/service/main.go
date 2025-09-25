package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Anfisa111/task-2-2/intheap"
)

var errorReading = errors.New("reading error")

func getPreferredDish(dishCount int) (int, error) {
	var priority, k int
	h := &intheap.IntHeap{}
	heap.Init(h)

	for range dishCount {
		if _, err := fmt.Scan(&priority); err != nil {
			return 0, errorReading
		}
		heap.Push(h, priority)
	}

	if _, err := fmt.Scan(&k); err != nil {
		return 0, errorReading
	}

	for range k - 1 {
		heap.Pop(h)
	}

	return heap.Pop(h).(int), nil
}

func main() {
	var dishCount int
	if _, err := fmt.Scan(&dishCount); err != nil {
		fmt.Println("Error reading", err)

		return
	}

	res, err := getPreferredDish(dishCount)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(res)
}
