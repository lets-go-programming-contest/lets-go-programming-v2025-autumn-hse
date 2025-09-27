package main

import (
	"container/heap"

	"fmt"

	"github.com/JingolBong/task-2-2/internal/dishorder"
)

func main() {
	var numberOfDishes, preferedDishNumb int
	var stringOfPref string

	fmt.Scanln(&numberOfDishes)
	fmt.Scanln(&stringOfPref)
	fmt.Scanln(&preferedDishNumb)

	containerOfDishes := &dishorder.PrefOrder{}
	heap.Init(containerOfDishes)
	containerOfDishes.AddFromString(stringOfPref)

	var dishWanted int
	for i := 0; i < preferedDishNumb; i++ {
		dishWanted = heap.Pop(containerOfDishes).(int)
	}

	fmt.Println(dishWanted)
}
