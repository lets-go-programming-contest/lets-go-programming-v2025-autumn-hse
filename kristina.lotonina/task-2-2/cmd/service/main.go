package main

import (
	"container/heap"
	"fmt"

	"github.com/kef1rch1k/task-2-2/internal/dishes"
)

func main() {
	var number int
	if _, err := fmt.Scan(&number); err != nil {
		fmt.Println("unable to read number:", err)

		return
	}

	dishesContainer := &dishes.Heap{}
	heap.Init(dishesContainer)

	for range number {
		var preferences int
		if _, err := fmt.Scan(&preferences); err != nil {
			fmt.Println("unable to read preference:", err)

			return
		}

		heap.Push(dishesContainer, preferences)
	}

	var neededPreference int

	_, err := fmt.Scan(&neededPreference)
	if err != nil {
		fmt.Println("unable to read needed preference:", err)

		return
	}

	var result int

	if neededPreference > dishesContainer.Len() {
		fmt.Println("incorrect preference number")

		return
	}

	for range neededPreference {
		value := heap.Pop(dishesContainer)
		if intValue, ok := value.(int); ok {
			result = intValue
		} else {
			fmt.Printf("expected int, %v :", value)
		}
	}

	fmt.Println(result)
}
