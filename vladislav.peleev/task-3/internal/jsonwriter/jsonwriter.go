package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveJSON[T any](outputPath string, data T, dirPerm, filePerm os.FileMode) error {
	outputJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("cannot encode JSON: %w", err)
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("cannot create directory: %w", err)
	}

	if err := os.WriteFile(outputPath, outputJSON, filePerm); err != nil {
		return fmt.Errorf("cannot write output file: %w", err)
	}

	return nil
}
