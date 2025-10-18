package dishes

import "fmt"

type Heap []int

func (dishes *Heap) Len() int {
	return len(*dishes)
}

func (dishes *Heap) Less(first int, second int) bool {
	length := len(*dishes)
	if first < 0 || first >= length || second < 0 || second >= length {
		return false
	}

	return (*dishes)[first] > (*dishes)[second]
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
		panic(fmt.Sprintf("expected int, got %T", x))
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
