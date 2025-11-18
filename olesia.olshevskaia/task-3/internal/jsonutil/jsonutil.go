package jsonutil

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteOutput[T any](outputPath string, data T, dirmode, filemode os.FileMode) error {
	outputJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), dirmode); err != nil {
		return fmt.Errorf("error creating directory for %q: %w", outputPath, err)
	}

	if err := os.WriteFile(outputPath, outputJSON, filemode); err != nil {
		return fmt.Errorf("error writing output file %q: %w", outputPath, err)
	}

	return nil
}
