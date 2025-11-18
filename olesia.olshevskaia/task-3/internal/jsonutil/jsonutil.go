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

func ParseJSON[T any](outputFile string, data T, dirmode, filemode os.FileMode) error {
	absPath, err := filepath.Abs(outputFile)
	if err != nil {
		return fmt.Errorf("cannot get absolute path for %q: %w", outputFile, err)
	}

	dir := filepath.Dir(absPath)

	if err := os.MkdirAll(dir, dirmode); err != nil {
		return fmt.Errorf("cannot create directory %q: %w", dir, err)
	}

	outputJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	if err := os.WriteFile(absPath, outputJSON, filemode); err != nil {
		return fmt.Errorf("error writing output file %q: %w", absPath, err)
	}

	return nil
}
