package temp

const (
	MinBound = 15
	MaxBound = 30
)

func UpdateInterval(left int, right int, operator string, value int) (int, int) {
	switch operator {
		case "<=":
			if value < right {
				right = value
			}
		case ">=":
			if value > left {
				left = value
			}
		default:

			return MinBound, MaxBound
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
