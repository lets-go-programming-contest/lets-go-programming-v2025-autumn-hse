package currency

import (
	"fmt"

	"github.com/Olesia.Ol/task-3/internal/jsonutil"
	"github.com/Olesia.Ol/task-3/internal/model"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
)

func WriteJSON(outputFile string, currencies []model.Currency) error {
	if err := jsonutil.WriteOutput(outputFile, currencies, dirPerm, filePerm); err != nil {
		return fmt.Errorf("failed to write output to %q: %w", outputFile, err)
	}
	return nil
}
