package currency

import (
	"os"

	"github.com/Olesia.Ol/task-3/internal/jsonutil"
	"github.com/Olesia.Ol/task-3/internal/model"
)

const (
	dirPerm  = os.ModePerm
	filePerm = 0o644
)

func WriteJSON(outputFile string, currencies []model.Currency) error {
	return jsonutil.ParseJSON(outputFile, currencies, dirPerm, filePerm)
}
