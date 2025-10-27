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
		heapCount     int
		preferredRank int
	)

	if _, err := fmt.Fscan(input, &heapCount); err != nil {
		return
	}

	minHeap := &myheap.MinHeap{}
	stdheap.Init(minHeap)

	for range heapCount {
		var preferenceScore int

		if _, err := fmt.Fscan(input, &preferenceScore); err != nil {
			return
		}

		stdheap.Push(minHeap, preferenceScore)
	}

	if _, err := fmt.Fscan(input, &preferredRank); err != nil {
		return
	}

	if preferredRank < 1 || preferredRank > heapCount {
		return
	}

	for range heapCount - preferredRank {
		stdheap.Pop(minHeap)
	}

	value := stdheap.Pop(minHeap)

	if value == nil {
		return
	}

	val, ok := value.(int)
	if !ok {
		return
	}

	fmt.Println(val)
}
