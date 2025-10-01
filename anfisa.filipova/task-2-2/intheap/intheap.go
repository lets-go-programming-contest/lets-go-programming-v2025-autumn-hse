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
	} else {
		return
	}
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
