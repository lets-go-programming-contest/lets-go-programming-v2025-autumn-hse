package intheap

type Rating []int

func (rating *Rating) Len() int {
	return len(*rating)
}

func (rating *Rating) Less(index, jndex int) bool {
	length := rating.Len()
	if 0 > index || index >= length || 0 > jndex || jndex >= length {
		panic("index out of range")
	}

	return (*rating)[index] > (*rating)[jndex]
}

func (rating *Rating) Swap(index, jndex int) {
	length := rating.Len()
	if 0 > index || index >= length || 0 > jndex || jndex >= length {
		panic("index out of range")
	}

	(*rating)[index], (*rating)[jndex] = (*rating)[jndex], (*rating)[index]
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
