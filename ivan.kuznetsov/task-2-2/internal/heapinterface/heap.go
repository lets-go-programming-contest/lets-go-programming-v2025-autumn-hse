package heapinterface

type Rating []int

func (r Rating) Len() int {
	return len(r)
}

func (r Rating) Less(i, j int) bool {
	return r[i] > r[j]
}

func (r Rating) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r *Rating) Push(x interface{}) {
	*r = append(*r, x.(int))
}

func (r *Rating) Pop() interface{} {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}
