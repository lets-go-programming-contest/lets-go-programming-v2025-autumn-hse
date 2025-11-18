package jsonutil

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteOutput[T any](outputPath string, data T) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	dir := filepath.Dir(outputPath)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create output directory for %q: %w", outputPath, err)
		}
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %q: %w", outputPath, err)
	}
	defer outputFile.Close()

	_, err = outputFile.Write(bytes)
	if err != nil {
		return fmt.Errorf("failed to write to %q: %w", outputPath, err)
	}

	return nil
}
