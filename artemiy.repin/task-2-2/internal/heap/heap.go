package heap

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	length := len(*h)

	if i < 0 || j < 0 || i >= length || j >= length {
		return false
	}

	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	length := len(*h)

	if i < 0 || j < 0 || i >= length || j >= length || i == j {
		return
	}
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(value any) {
	value_, ok := value.(int)

	if !ok {
		panic("Need to push int")
	}

	*h = append(*h, value_)
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
