package main

import (
	"errors"
	"fmt"
	"slices"

	"github.com/JingolBong/task-2-2/internal/dishorder"
)

var errFailedToScan = errors.New("Invalid input")

func main() {
	var numberOfDishes int

	containerOfDishes := &dishorder.PrefOrder{}

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

	slices.Sort(*containerOfDishes)
	fmt.Println((*containerOfDishes)[len(*containerOfDishes)-preferedDishNumb])
}
