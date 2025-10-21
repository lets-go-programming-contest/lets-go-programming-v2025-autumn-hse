package main

import (
	"container/heap"
	"fmt"

	"github.com/JingolBong/task-2-2/internal/dishorder"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from panic: ", r)
		}
	}()

	var (
		containerOfDishes dishorder.PrefOrder
		numberOfDishes    int
	)

	heap.Init(&containerOfDishes)

	if _, err := fmt.Scan(&numberOfDishes); err != nil {
		fmt.Println("when scanning dish count: ", err)

		return
	}

	for range numberOfDishes {
		var dishPrefNumber int

		if _, err := fmt.Scan(&dishPrefNumber); err != nil {
			fmt.Println("when scanning preference order: ", err)

			return
		}

		heap.Push(&containerOfDishes, dishPrefNumber)
	}

	var preferedDishNumb int
	if _, err := fmt.Scanln(&preferedDishNumb); err != nil {
		fmt.Println("when scanning preferred dish number: ", err)

		return
	}

	if containerOfDishes.Len() < preferedDishNumb {
		fmt.Println("too big dish number")

		return
	}

	for range containerOfDishes.Len() - preferedDishNumb {
		heap.Pop(&containerOfDishes)
	}

	last := heap.Pop(&containerOfDishes)

	if last == nil {
		fmt.Println("no dishes left")

		return
	}

	fmt.Println(last)
}
