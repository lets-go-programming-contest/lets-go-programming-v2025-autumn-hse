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

	for index := 0; index < heapCount; index++ {
		if _, err := fmt.Fscan(input, &preferenceScore); err != nil {
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

	for removedCount := 0; removedCount < heapCount-preferredRank; removedCount++ {
		stdheap.Pop(inputHeap)
	}

	selectedScore := stdheap.Pop(inputHeap).(int)
	fmt.Println(selectedScore)
}
