package currency

import (
	"github.com/Olesia.Ol/task-3/internal/jsonutil"
	"github.com/Olesia.Ol/task-3/internal/model"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
)

func WriteJSON(outputFile string, currencies []model.Currency) error {
	return jsonutil.WriteOutput(outputFile, currencies, dirPerm, filePerm)
}
