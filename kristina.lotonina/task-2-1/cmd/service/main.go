package main

import (
	"fmt"

	"github.com/kef1rch1k/task-2-1/internal/temperature"
)

func main() {
	var number, count int

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("unable to read")

		return
	}

	for range number {
		_, err = fmt.Scan(&count)
		if err != nil {
			fmt.Println("unable to read")

			return
		}

		FindTemp(count)
	}
}
