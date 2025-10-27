package heap

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(idxI, idxJ int) bool {
	n := len(*h)

	if idxI < 0 || idxJ < 0 || idxI >= n || idxJ >= n {
		return false
	}

	return (*h)[idxI] < (*h)[idxJ]
}

func (h *MinHeap) Swap(idxI, idxJ int) {
	n := len(*h)

	if idxI < 0 || idxJ < 0 || idxI >= n || idxJ >= n || idxI == idxJ {
		return
	}

	(*h)[idxI], (*h)[idxJ] = (*h)[idxJ], (*h)[idxI]
}

func (h *MinHeap) Push(value any) {
	v, ok := value.(int)
	if !ok {
		return
	}

	*h = append(*h, v)
}

func (h *MinHeap) Pop() any {
	length := len(*h)
	if length == 0 {
		return nil
	}

	x := (*h)[length-1]
	*h = (*h)[:length-1]

	return x
}
