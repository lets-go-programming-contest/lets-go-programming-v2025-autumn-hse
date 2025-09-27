package dishorder

type PrefOrder []int

func (h PrefOrder) Len() int           { return len(h) }
func (h PrefOrder) Less(i, j int) bool { return h[i] > h[j] }
func (h PrefOrder) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PrefOrder) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *PrefOrder) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}
