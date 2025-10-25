package temp

const (
	MinBound = 15
	MaxBound = 30
)

type Temperature struct {
	LeftBound  int
	RightBound int
}

func UpdateTemperature() Temperature {
	temperature := Temperature{
		LeftBound:  MinBound,
		RightBound: MaxBound,
	}

	return temperature
}

func (t *Temperature) UpdateInterval(operator string, value int) {
	switch operator {
	case "<=":
		if value < t.RightBound {
			t.RightBound = value
		}

	case ">=":
		if value > t.LeftBound {
			t.LeftBound = value
		}

	default:
		t.LeftBound = MinBound
		t.RightBound = MaxBound
	}
}

func (t *Temperature) GetOptimal() int {
	left := t.LeftBound
	right := t.RightBound

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
