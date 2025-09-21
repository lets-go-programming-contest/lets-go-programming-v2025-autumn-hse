package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var (
		dishCount, dish, rank int
	)

	_, err = fmt.Scanln(&dishCount)
	if err != nil {
		return
	}

	type IntHeap []int

	for range dishCount {
		_, err = fmt.Scan(&dish)
		if err != nil {
			return
		}

	}

	_, err = fmt.Scanln(&rank)
	if err != nil {
		return
	}

}
