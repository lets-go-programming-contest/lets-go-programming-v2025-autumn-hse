package intmath

func LargerInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func SmallerInt(a, b int) int {
	if a > b {
		return b
	}

	return a
}
