package temp

func UpdateInterval(L int, R int, op string, val int) (int, int) {
	if op == "<=" {
		if val < R {
			R = val
		}
	} else if op == ">=" {
		if val > L {
			L = val
		}
	}
	return L, R
}

func GetOptimal(L int, R int) int {
	if L > R {
		return -1
	}

	if L < 15 {
		L = 15
	}
	if R > 30 {
		R = 30
	}
	return L
}
