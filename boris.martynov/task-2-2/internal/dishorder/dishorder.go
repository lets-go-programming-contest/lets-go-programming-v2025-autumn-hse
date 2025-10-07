package dishorder

type PrefOrder []int

func (h *PrefOrder) Len() int {
	return len(*h)
}

func (h *PrefOrder) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *PrefOrder) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *PrefOrder) Push(value any) {
	pushedInt, ok := value.(int)

	if !ok {
		panic("Pushed not int")
	}

	*h = append(*h, pushedInt)
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
