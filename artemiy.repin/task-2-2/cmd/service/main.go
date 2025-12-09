package main

import (
	stdheap "container/heap"
	"fmt"

	myheap "github.com/Nevermind0911/task-2-2/internal/heap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	minHeap, heapCount, preferredRank, ok := prepareHeap()
	if !ok {
		return
	}

	if !popUntilRank(minHeap, heapCount, preferredRank) {
		return
	}

	printTop(minHeap)
}

func prepareHeap() (*myheap.MinHeap, int, int, bool) {
	var heapCount int

	if _, err := fmt.Scan(&heapCount); err != nil {
		fmt.Println("cannot read heap count")

		return nil, 0, 0, false
	}

	if heapCount <= 0 {
		fmt.Println("heap count must be > 0")

		return nil, 0, 0, false
	}

	minHeap := &myheap.MinHeap{}
	stdheap.Init(minHeap)

	for range heapCount {
		var preferenceScore int

		if _, err := fmt.Scan(&preferenceScore); err != nil {
			fmt.Println("cannot read preference score")

			return nil, 0, 0, false
		}

		stdheap.Push(minHeap, preferenceScore)
	}

	var preferredRank int

	if _, err := fmt.Scan(&preferredRank); err != nil {
		fmt.Println("cannot read preferred rank")

		return nil, 0, 0, false
	}

	if preferredRank < 1 || preferredRank > heapCount {
		fmt.Println("preferred rank is out of range")

		return nil, 0, 0, false
	}

	return minHeap, heapCount, preferredRank, true
}

func popUntilRank(minHeap *myheap.MinHeap, heapCount, preferredRank int) bool {
	popsCount := heapCount - preferredRank

	for range popsCount {
		if minHeap.Len() == 0 {
			fmt.Println("heap is empty, cannot pop")

			return false
		}

		stdheap.Pop(minHeap)
	}

	return true
}

func printTop(minHeap *myheap.MinHeap) {
	if minHeap.Len() == 0 {
		fmt.Println("heap is empty before final pop")

		return
	}

	value := stdheap.Pop(minHeap)
	if value == nil {
		fmt.Println("got nil from heap")

		return
	}

	val, ok := value.(int)
	if !ok {
		fmt.Println("heap value is not int")

		return
	}

	fmt.Println(val)
}
