package temp

const (
	MinBound = 15
	MaxBound = 30
)

func UpdateInterval(left int, right int, op string, val int) (int, int) {
	if op == "<=" {
		if val < right {
			right = val
		}
	} else if op == ">=" {
		if val > left {
			left = val
		}
	}
	return left, right
}

func GetOptimal(left int, right int) int {
	if left > right {
		return -1
	}

	if left < MinBound {
		left = MinBound
	}
	if right > MaxBound {
		right = MaxBound
	}

	if left > right {
		return -1
	}
	return left
}
