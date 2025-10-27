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

		return
	}
}

func (t *Temperature) GetOptimal() int {
	if t.LeftBound <= t.RightBound {

		return t.LeftBound
	}

	return -1
}
