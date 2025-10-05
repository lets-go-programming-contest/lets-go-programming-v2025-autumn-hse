package dishes

type Heap []int

func (dishes Heap) Len() int {
	return len(dishes)
}

func (dishes Heap) Less(i int, j int) bool {
	return dishes[i] > dishes[j]
}

func (dishes Heap) Swap(i int, j int) {
	dishes[i], dishes[j] = dishes[j], dishes[i]
}

func (dishes *Heap) Push(x interface{}) {
	*dishes = append(*dishes, x.(int))
}

func (dishes *Heap) Pop() interface{} {
	old := *dishes
	n := len(old)
	x := old[n-1]
	*dishes = old[0 : n-1]
	return x
}
