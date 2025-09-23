package myheap

type Heap []int

func (h *Heap) Len() int {
	dereferencedHeap := *h

	return len(dereferencedHeap)
}

func (h *Heap) Less(i, j int) bool {
	dereferencedHeap := *h

	return dereferencedHeap[i] < dereferencedHeap[j]
}

func (h *Heap) Swap(i, j int) {
	dereferencedHeap := *h

	dereferencedHeap[i], dereferencedHeap[j] = dereferencedHeap[j], dereferencedHeap[i]
}

func (h *Heap) Push(x any) {
	castedElem, ok := x.(int)
	if ok {
		*h = append(*h, castedElem)
	}
}

func (h *Heap) Pop() any {
	oldHeap := *h
	value := oldHeap[len(oldHeap)-1]
	*h = oldHeap[0 : len(oldHeap)-1]

	return value
}
