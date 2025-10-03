package dishorder

type PrefOrder []int //nolint:recvcheck

func (h PrefOrder) Len() int {
	return len(h)
}

func (h PrefOrder) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h PrefOrder) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PrefOrder) Push(x any) {
	if pushedInt, ok := x.(int); ok {
		*h = append(*h, pushedInt)
	} else {
		panic("Pushed not int")
	}
}

func (h *PrefOrder) Pop() any {
	if len(*h) == 0 {
		return nil
	}

	old := *h
	neededIndex := len(old) - 1
	x := old[neededIndex]

	*h = old[0:neededIndex]

	return x
}
