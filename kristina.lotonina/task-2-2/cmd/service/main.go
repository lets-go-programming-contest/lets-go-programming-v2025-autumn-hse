package main

import (
	"container/heap"
	"fmt"
	"slices"

	"github.com/kef1rch1k/task-2-2/internal/dishes"
)

func main() {
	var number int
	_, err := fmt.Scan(&number)

	dishesContainer := &dishes.Heap{}
	heap.Init(dishesContainer)

	if err != nil {
		fmt.Print("unable to read number :", err)

		return
	}

	for range number {
		var preferences int
		if _, err = fmt.Scan(&preferences); err != nil {
			fmt.Print("unable to read preference :", err)

			return
		}

		heap.Push(dishesContainer, preferences)
	}

	var neededPreference int

	_, err = fmt.Scan(&neededPreference)
	if err != nil {
		fmt.Print("unable to read needed preference :", err)

		return
	}

	chosenDish := len(*dishesContainer) - neededPreference
	slices.Sort(*dishesContainer)

	fmt.Println((*dishesContainer)[chosenDish])
}
