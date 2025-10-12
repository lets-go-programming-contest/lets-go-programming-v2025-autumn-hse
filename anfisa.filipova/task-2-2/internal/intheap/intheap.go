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
	val, ok := v.(int)
	if !ok {
		panic("Type assertion failed")
	}

	*h = append(*h, val)
}

func (h *IntHeap) Pop() any {
	oldHeap := *h
	length := len(oldHeap)

	if length == 0 {
		return nil
	}

	elem := oldHeap[length-1]
	*h = oldHeap[0 : length-1]

	return elem
}
