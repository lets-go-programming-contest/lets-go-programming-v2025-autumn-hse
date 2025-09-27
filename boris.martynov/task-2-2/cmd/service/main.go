package main

import (
	"fmt"
	"slices"

	"github.com/JingolBong/task-2-2/internal/dishorder"
)

func main() {
	var numberOfDishes, preferedDishNumb int

	containerOfDishes := &dishorder.PrefOrder{}

	if _, err := fmt.Scan(&numberOfDishes); err != nil {
		return
	}

	for range numberOfDishes {
		var dishPrefNumber int

		if _, err := fmt.Scan(&dishPrefNumber); err != nil {
			return
		}

		containerOfDishes.Push(dishPrefNumber)
	}

	if _, err := fmt.Scanln(&preferedDishNumb); err != nil {
		return
	}

	slices.Sort(*containerOfDishes)
	fmt.Println((*containerOfDishes)[len(*containerOfDishes)-preferedDishNumb])
}
