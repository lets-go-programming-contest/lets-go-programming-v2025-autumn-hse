package dishes

import "fmt"

type Heap []int

func (dishes *Heap) Len() int {
	return len(*dishes)
}

func (dishes *Heap) Less(i int, j int) bool {
	return (*dishes)[i] > (*dishes)[j]
}

func (dishes *Heap) Swap(first int, second int) {
	length := len(*dishes)
	if first < 0 || first >= length || second < 0 || second >= length {
		return
	}

	(*dishes)[first], (*dishes)[second] = (*dishes)[second], (*dishes)[first]
}

func (dishes *Heap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		panic(fmt.Sprintf("expected int, %v", x))
	}

	*dishes = append(*dishes, value)
}

func (dishes *Heap) Pop() interface{} {
	if old := *dishes; len(old) == 0 {
		return nil
	}

	old := *dishes
	x := old[(len(old))-1]
	*dishes = old[0 : (len(old))-1]

	return x
}
