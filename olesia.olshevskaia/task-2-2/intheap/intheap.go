package intheap

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	newX, ok := x.(int)
	if !ok {
		panic("invalid type, expected int")
	}

	*h = append(*h, newX)
}

func (h *IntHeap) Pop() any {
	old := *h
	length := len(*h)

	if length == 0 {
		return nil
	}

	x := old[length-1]
	*h = old[:length-1]

	return x
}
