package currency

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Olesia.Ol/task-3/internal/model"
)

const (
	dirPerm  = os.ModePerm
	filePerm = 0o644
)

func parseJSON[T any](outputFile string, data T, dirmode, filemode os.FileMode) error {
	outputJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(outputFile), dirmode); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	if err := os.WriteFile(outputFile, outputJSON, filemode); err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}

func WriteJSON(outputFile string, currencies []model.Currency) error {
	return parseJSON(outputFile, currencies, dirPerm, filePerm)
}
