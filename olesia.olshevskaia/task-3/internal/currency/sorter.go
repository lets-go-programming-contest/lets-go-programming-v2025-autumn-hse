package currency

import (
	"sort"

	"github.com/Olesia.Ol/task-3/internal/model"
)

func Sort(currencies []model.Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].RateValue > currencies[j].RateValue
	})
}
