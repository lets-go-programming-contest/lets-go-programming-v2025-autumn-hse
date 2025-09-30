package main

import (
	"errors"
	"fmt"
	"slices"

	"github.com/JingolBong/task-2-2/internal/dishorder"
)

var errFailedToScan = errors.New("invalid input")

func main() {
	containerOfDishes := &dishorder.PrefOrder{}

	var numberOfDishes int

	if _, err := fmt.Scan(&numberOfDishes); err != nil {
		fmt.Println(errFailedToScan)

		return
	}

	for range numberOfDishes {
		var dishPrefNumber int

		if _, err := fmt.Scan(&dishPrefNumber); err != nil {
			fmt.Println(errFailedToScan)

			return
		}

		containerOfDishes.Push(dishPrefNumber)
	}

	var preferedDishNumb int
	if _, err := fmt.Scanln(&preferedDishNumb); err != nil {
		fmt.Println(errFailedToScan)

		return
	}

	dishedChoosed := len(*containerOfDishes) - preferedDishNumb
	slices.Sort(*containerOfDishes)
	fmt.Println((*containerOfDishes)[dishedChoosed])
}
