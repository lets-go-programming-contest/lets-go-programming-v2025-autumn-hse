package heapmax

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) } // lint warn if has (*IntHeap) and (IntHeap) methods
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	v, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, v)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1] // slice [0:0] is empty slice!!!

	return x
}
