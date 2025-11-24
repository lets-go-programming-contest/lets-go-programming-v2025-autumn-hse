package main

import (
	"bufio"
	stdheap "container/heap"
	"fmt"
	"os"

	myheap "github.com/Nevermind0911/task-2-2/internal/heap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	input := bufio.NewReader(os.Stdin)

	var heapCount int

	if _, err := fmt.Fscan(input, &heapCount); err != nil {
		fmt.Println("cannot read heap count")

		return
	}

	if heapCount <= 0 {
		fmt.Println("heap count must be > 0")

		return
	}

	minHeap := &myheap.MinHeap{}
	stdheap.Init(minHeap)

	for range heapCount {
		var preferenceScore int

		if _, err := fmt.Fscan(input, &preferenceScore); err != nil {
			fmt.Println("cannot read preference score")

			return
		}

		stdheap.Push(minHeap, preferenceScore)
	}

	var preferredRank int

	if _, err := fmt.Fscan(input, &preferredRank); err != nil {
		fmt.Println("cannot read preferred rank")

		return
	}

	if preferredRank < 1 || preferredRank > heapCount {
		fmt.Println("preferred rank is out of range")

		return
	}

	for range heapCount - preferredRank {
		if minHeap.Len() == 0 {
			fmt.Println("heap is empty, cannot pop")

			return
		}

		stdheap.Pop(minHeap)
	}

	if minHeap.Len() == 0 {
		fmt.Println("heap is empty before final pop")

		return
	}

	value := stdheap.Pop(minHeap)
	val := value.(int)

	fmt.Println(val)
}
