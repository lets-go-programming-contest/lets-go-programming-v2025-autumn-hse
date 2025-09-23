package heapmax

const (
	errPushCast      = "error: bad cast to int"
	errPopOutOfRange = ""
)

type IntHeap []int

//nolint:recvcheck // is linked type
func (h IntHeap) Len() int {
	return len(h)
}

//nolint:recvcheck
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

//nolint:recvcheck
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
