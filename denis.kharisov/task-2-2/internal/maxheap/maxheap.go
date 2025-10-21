package maxheap

type Maxheap []int

func (heap *Maxheap) Len() int {
	return len(*heap)
}

func (heap *Maxheap) Less(i, j int) bool {
	if i >= len(*heap) || j >= len(*heap) {
		panic("Index out of range")
	}

	return (*heap)[i] > (*heap)[j]
}

func (heap *Maxheap) Swap(i, j int) {
	if i >= len(*heap) || j >= len(*heap) {
		panic("Index out of range")
	}

	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *Maxheap) Push(x any) {
	if val, ok := x.(int); !ok {
		panic("MaxHeap.Push: expected int")
	} else {
		*heap = append(*heap, val)
	}
}

func (heap *Maxheap) Pop() any {
	if heap.Len() == 0 {
		return nil
	}

	old := *heap
	lenOld := len(old)
	lastElem := old[lenOld-1]
	*heap = old[0 : lenOld-1]

	return lastElem
}
