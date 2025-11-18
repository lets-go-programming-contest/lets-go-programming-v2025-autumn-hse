package jsonutil

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	DirPerm  = 0o755
	FilePerm = 0o644
)

func ParseJSON[T any](outputPath string, data T) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, DirPerm); err != nil {
		return fmt.Errorf("failed to create output directory %q: %w", dir, err)
	}

	if err := os.WriteFile(outputPath, bytes, FilePerm); err != nil {
		return fmt.Errorf("failed to write to %q: %w", outputPath, err)
	}

	return nil
}
