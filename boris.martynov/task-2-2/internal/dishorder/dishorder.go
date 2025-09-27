package dishorder

type PrefOrder []int

func (h PrefOrder) Len() int           { return len(h) }           //nolint
func (h PrefOrder) Less(i, j int) bool { return h[i] > h[j] }      //nolint
func (h PrefOrder) Swap(i, j int)      { h[i], h[j] = h[j], h[i] } //nolint

func (h *PrefOrder) Push(x any) {
	if pushedInt, noerr := x.(int); noerr {
		*h = append(*h, pushedInt)
	} else {
		panic("Pushed not int")
	}
}

func (h *PrefOrder) Pop() any {
	if len(*h) == 0 {
		panic("Pop emty")
	}

	old := *h
	x := old[len(old)-1]

	*h = old[0 : len(old)-1]

	return x
}
