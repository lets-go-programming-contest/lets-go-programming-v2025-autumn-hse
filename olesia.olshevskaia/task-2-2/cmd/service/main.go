package main

import (
	"container/heap"
	"fmt"

	"github.com/olesia.olshevsraia/task-2-2/intheap"
)

func main() {
	myHeap := &intheap.IntHeap{}
	heap.Init(myHeap)

	var numberDishes int

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

	var desiredDish int
	if _, err := fmt.Scanln(&desiredDish); err != nil {
		fmt.Println("error reading priority", err)

		return
	}

	for range desiredDish - 1 {
		val := heap.Pop(myHeap)

		if val == nil {
			fmt.Println("Heap is empty")

			break
		}
	}

	value := heap.Pop(myHeap)

	if value == nil {
		fmt.Println("Heap is empty")

		return
	}

	count, typeCastingOk := value.(int)

	if !typeCastingOk {
		fmt.Println("invalid type from heap.Pop")

		return
	}

	fmt.Println(count)

}
