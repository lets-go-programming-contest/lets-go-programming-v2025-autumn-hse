package maxheap

type Maxheap []int

func (heap *Maxheap) Len() int {
	return len(*heap)
}

func (heap *Maxheap) Less(i, j int) bool {
	return (*heap)[i] > (*heap)[j]
}

func (heap *Maxheap) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *Maxheap) Push(x any) {
	value, flag := x.(int)
	if !flag {
		panic("MaxHeap.Push: expected int")
	}
	*heap = append(*heap, value)
}

func (heap *Maxheap) Pop() any {
	old := *heap
	lenOld := len(old)
	lastElem := old[lenOld-1]
	*heap = old[0 : lenOld-1]

	return lastElem
}
