package dishorder

import (
	"container/heap"
	"strconv"
	"strings"
)

type PrefOrder []int

func (h PrefOrder) Len() int           { return len(h) }
func (h PrefOrder) Less(i, j int) bool { return h[i] > h[j] }
func (h PrefOrder) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PrefOrder) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *PrefOrder) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *PrefOrder) AddFromString(s string) {
	for _, p := range strings.Fields(s) {
		n, _ := strconv.Atoi(p)
		heap.Push(h, n)
	}
}
