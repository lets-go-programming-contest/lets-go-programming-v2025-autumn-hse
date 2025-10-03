package main

import (
	"fmt"
	"slices"

	"github.com/JingolBong/task-2-2/internal/dishorder"
)

func main() {
	containerOfDishes := &dishorder.PrefOrder{}

	var numberOfDishes int

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

		containerOfDishes.Push(dishPrefNumber)
	}

	var preferedDishNumb int
	if _, err := fmt.Scanln(&preferedDishNumb); err != nil {
		fmt.Println("when scanning prefered dish number: ", err)

		return
	}

	dishedChoosed := len(*containerOfDishes) - preferedDishNumb
	slices.Sort(*containerOfDishes)
	fmt.Println((*containerOfDishes)[dishedChoosed])
}
