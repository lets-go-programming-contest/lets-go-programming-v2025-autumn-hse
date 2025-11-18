package currency

import (
	"github.com/Olesia.Ol/task-3/internal/jsonutil"
	"github.com/Olesia.Ol/task-3/internal/model"
)

func WriteJSON(outputFile string, currencies []model.Currency) error {
	return jsonutil.ParseJSON(outputFile, currencies, jsonutil.DirPerm, jsonutil.FilePerm)
}
