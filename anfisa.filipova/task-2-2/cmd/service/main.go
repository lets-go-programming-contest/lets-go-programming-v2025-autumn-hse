package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Anfisa111/task-2-2/internal/intheap"
)

var (
	errReading          = errors.New("reading error")
	errTypeAssertion    = errors.New("error type assertion")
	errInvalidChoiceIdx = errors.New("error invalid choiceIdx")
)

func readInput() (intheap.IntHeap, int, error) {
	var dishCount int
	if _, err := fmt.Scan(&dishCount); err != nil {
		return nil, 0, errReading
	}

	heapint := &intheap.IntHeap{}

	heap.Init(heapint)

	for range dishCount {
		var priority int
		if _, err := fmt.Scan(&priority); err != nil {
			return nil, 0, errReading
		}

		heap.Push(heapint, priority)
	}

	var choiceIdx int
	if _, err := fmt.Scan(&choiceIdx); err != nil {
		return nil, 0, errReading
	}

	return *heapint, choiceIdx, nil
}

func getPreferredDish(heapint intheap.IntHeap, choiceIdx int) (int, error) {
	heapSize := len(heapint)
	if heapSize < choiceIdx {
		return 0, errInvalidChoiceIdx
	}

	for range choiceIdx - 1 {
		heap.Pop(&heapint)
	}

	val, ok := heap.Pop(&heapint).(int)
	if !ok {
		return 0, errTypeAssertion
	}

	return val, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	heapint, choiceIdx, err := readInput()
	if err != nil {
		fmt.Println("Error: ", err)

		return
	}

	res, err := getPreferredDish(heapint, choiceIdx)
	if err != nil {
		fmt.Println("Error: ", err)

		return
	}

	fmt.Println(res)
}
