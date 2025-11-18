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

	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("failed to create output directory for %q: %w", outputPath, err)
	}

	if err := os.WriteFile(outputPath, bytes, 0o644); err != nil {
		return fmt.Errorf("failed to write to %q: %w", outputPath, err)
	}

	return nil
}
