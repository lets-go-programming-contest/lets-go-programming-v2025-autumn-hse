package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Anfisa111/task-2-2/intheap"
)

var (
	errReading       = errors.New("reading error")
	errTypeAssertion = errors.New("error type assertion")
)

func getPreferredDish(dishCount int) (int, error) {
	var priority, choiceIdx int

	heapint := &intheap.IntHeap{}

	heap.Init(heapint)

	for range dishCount {
		if _, err := fmt.Scan(&priority); err != nil {
			return 0, errReading
		}

		heap.Push(heapint, priority)
	}

	if _, err := fmt.Scan(&choiceIdx); err != nil {
		return 0, errReading
	}

	for range choiceIdx - 1 {
		heap.Pop(heapint)
	}

	if val, ok := heap.Pop(heapint).(int); ok {
		return val, nil
	} else {
		return 0, errTypeAssertion
	}
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
