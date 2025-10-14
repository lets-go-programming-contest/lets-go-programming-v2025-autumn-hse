package intheap

type Rating []int

func (rating *Rating) Len() int {
	return len(*rating)
}

func (rating *Rating) Less(i, j int) bool {
	length := rating.Len()
	if 0 > i || i >= length || 0 > j || j >= length {
		panic("index out of range")
	}

	return (*rating)[i] > (*rating)[j]
}

func (rating *Rating) Swap(i, j int) {
	length := rating.Len()
	if 0 > i || i >= length || 0 > j || j >= length {
		panic("index out of range")
	}

	(*rating)[i], (*rating)[j] = (*rating)[j], (*rating)[i]
}

func (rating *Rating) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		panic("type conversion error")
	}

	*rating = append(*rating, value)
}

func (rating *Rating) Pop() interface{} {
	old := *rating
	length := len(old)

	if length == 0 {
		return nil
	}

	value := old[length-1]
	*rating = old[0 : length-1]

	return value
}
