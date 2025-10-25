package main

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Anfisa111/task-2-2/internal/intheap"
)

var (
	errTypeAssertion    = errors.New("error type assertion")
	errInvalidChoiceIdx = errors.New("error invalid choiceIdx")
)

func readIndex() (int, error) {
	var choiceIdx int
	if _, err := fmt.Scan(&choiceIdx); err != nil {
		return 0, fmt.Errorf("cannot read choice index: %w", err)
	}

	return choiceIdx, nil
}

func readDishes() (intheap.IntHeap, error) {
	var dishCount int
	if _, err := fmt.Scan(&dishCount); err != nil {
		return nil, fmt.Errorf("cannot read count of dishes: %w", err)
	}

	heapint := &intheap.IntHeap{}

	heap.Init(heapint)

	for range dishCount {
		var priority int
		if _, err := fmt.Scan(&priority); err != nil {
			return nil, fmt.Errorf("cannot read priority: %w", err)
		}

		heap.Push(heapint, priority)
	}

	return *heapint, nil
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

	dishes, err := readDishes()
	if err != nil {
		fmt.Println("Failed to read dishes: ", err)

		return
	}

	choiceIdx, err := readIndex()
	if err != nil {
		fmt.Println("Failed to read choice index: ", err)

		return
	}

	res, err := getPreferredDish(dishes, choiceIdx)
	if err != nil {
		fmt.Println("Failed to get preferred dish ", err)

		return
	}

	fmt.Println(res)
}
