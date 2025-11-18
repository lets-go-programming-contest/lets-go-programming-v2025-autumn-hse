package jsonutil

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	DirPerm  = os.ModePerm
	FilePerm = 0o644
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
