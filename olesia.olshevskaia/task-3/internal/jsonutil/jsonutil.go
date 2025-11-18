package jsonutil

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const dirPer = 0o755

func WriteOutput[T any](outputPath string, data T) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(outputPath), dirPer)
	if err != nil {
		return fmt.Errorf("failed to create output directory for %q: %w", outputPath, err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %q: %w", outputPath, err)
	}

	defer func() {
		if ferr := outputFile.Close(); ferr != nil {
			panic(fmt.Errorf("failed to close file %q: %w", outputPath, ferr))
		}
	}()

	_, err = outputFile.Write(bytes)
	if err != nil {
		return fmt.Errorf("failed to write to %q: %w", outputPath, err)
	}

	return nil
}
