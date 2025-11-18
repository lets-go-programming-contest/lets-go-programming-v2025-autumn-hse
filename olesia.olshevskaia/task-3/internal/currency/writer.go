package currency

import (
	"fmt"

	"github.com/Olesia.Ol/task-3/internal/jsonutil"
	"github.com/Olesia.Ol/task-3/internal/model"
)

func WriteJSON(outputFile string, currencies []model.Currency) error {
	if err := jsonutil.ParseJSON(outputFile, currencies); err != nil {
		return fmt.Errorf("failed to write currencies to JSON file: %w", err)
	}
	return nil
}
