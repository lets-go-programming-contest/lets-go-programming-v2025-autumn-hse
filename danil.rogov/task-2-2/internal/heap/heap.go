package myheap

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	if i < 0 || i > h.Len() || j < 0 || j > h.Len() {
		panic("invalid index")
	}

	return (*h)[i] < (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	if i < 0 || i > h.Len() || j < 0 || j > h.Len() {
		panic("invalid index")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x any) {
	castedElem, ok := x.(int)
	if !ok {
		panic("element is not int")
	}

	*h = append(*h, castedElem)
}

func (h *Heap) Pop() any {
	if h.Len() == 0 || h == nil {
		return nil
	}

	oldHeap := *h
	value := oldHeap[len(oldHeap)-1]
	*h = oldHeap[0 : len(oldHeap)-1]

	return value
}
