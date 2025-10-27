package myheap

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if i >= len(*h) || i < 0 || j >= len(*h) || j < 0 {
		panic("index out of range")
	}

	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i >= len(*h) || i < 0 || j >= len(*h) || j < 0 {
		panic("index out of range")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		panic("cast to int error")
	}

	*h = append(*h, value)
}

func (h *IntHeap) Pop() any {
	if len(*h) == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
