package heapmax

import "errors"

var (
	errPushCast      = errors.New("bad cast to int")
	errPopOutOfRange = errors.New("out of range")
)

//nolint:recvcheck // is linked type
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	v, ok := x.(int)
	if !ok {
		panic(errPushCast)
	}

	*h = append(*h, v)
}

func (h *IntHeap) Pop() any {
	if len(*h) == 0 {
		panic(errPopOutOfRange)
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
