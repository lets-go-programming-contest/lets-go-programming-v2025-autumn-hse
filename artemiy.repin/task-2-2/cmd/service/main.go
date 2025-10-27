package main

import (
	"bufio"
	stdheap "container/heap"
	"fmt"
	"os"

	myheap "github.com/Nevermind0911/task-2-2/internal/heap"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	var (
		heapCount       int
		preferenceScore int
		preferredRank   int
	)

	if _, err := fmt.Fscan(input, &heapCount); err != nil {
		return
	}

	inputHeap := &myheap.MinHeap{}
	stdheap.Init(inputHeap)

	for range heapCount {
		if _, err := fmt.Scan(input, &preferenceScore); err != nil {
			return
		}

		stdheap.Push(inputHeap, preferenceScore)
	}

	if _, err := fmt.Fscan(input, &preferredRank); err != nil {
		return
	}

	if preferredRank < 1 || preferredRank > heapCount {
		return
	}

	for range heapCount - preferredRank {
		stdheap.Pop(inputHeap)
	}

	val, ok := stdheap.Pop(inputHeap).(int)
	if !ok {
		fmt.Println("error: ", ok)

		return
	}

	fmt.Println(val)
}
