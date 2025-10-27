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

	h := &myheap.MinHeap{}
	stdheap.Init(h)

	for i := 0; i < heapCount; i++ {
		var preferenceScore int

		if _, err := fmt.Fscan(input, &preferenceScore); err != nil {
			return
		}
		stdheap.Push(h, preferenceScore)
	}

	if _, err := fmt.Fscan(input, &preferredRank); err != nil {
		return
	}

	if preferredRank < 1 || preferredRank > heapCount {
		return
	}

	for i := 0; i < heapCount-preferredRank; i++ {
		stdheap.Pop(h)
	}

	v := stdheap.Pop(h)

	if v == nil {
		return
	}
	val, ok := v.(int)

	if !ok {
		return
	}

	fmt.Println(val)
}
