package main

import (
	"container/heap"
	"fmt"

	"github.com/olesia.olshevsraia/task-2-2/intheap"
)

func main() {
	var numberDishes, desiredDish, count int

	myHeap := &intheap.IntHeap{}
	heap.Init(myHeap)

	if _, err := fmt.Scanln(&numberDishes); err != nil {
		fmt.Println("error reading the number of dishes", err)

		return
	}

	var dish int
	for range numberDishes {
		if _, err := fmt.Scan(&dish); err != nil {
			fmt.Println("error reading dishf", err)

			return
		}

		heap.Push(myHeap, dish)
	}

	if _, err := fmt.Scanln(&desiredDish); err != nil {
		fmt.Println("error reading priority", err)

		return
	}

	for range desiredDish - 1 {
		heap.Pop(myHeap)
	}

	x := heap.Pop(myHeap)
	count, typeCastingOk := x.(int)

	if !typeCastingOk {
		fmt.Println("invalid type from heap.Pop")

		return
	}

	if _, err := fmt.Println(count); err != nil {
		fmt.Println("error while displaying the response", err)

		return
	}
}
