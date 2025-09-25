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

func (h *IntHeap) Push(v any) {
	if val, ok := v.(int); ok {
		*h = append(*h, val)
	}
}

func (h *IntHeap) Pop() any {
	oldHeap := *h
	l := len(oldHeap)
	elem := oldHeap[l-1]
	*h = oldHeap[0 : l-1]

	return elem
}
