package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const DefaultDirPerm = 0o755

func SaveJSON(outputPath string, data any, dirPerm os.FileMode) error {
	dir := filepath.Dir(outputPath)

	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("cannot create directory: %w", err)
	}

	var (
		file *os.File
		err  error
	)

	file, err = os.Create(outputPath)
	defer func() {
		if file != nil {
			if cerr := file.Close(); cerr != nil {
				fmt.Fprintf(os.Stderr, "warning: failed to close file %s: %v\n", outputPath, cerr)
			}
		}
	}()

	if err != nil {
		return fmt.Errorf("cannot create output file: %w", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err = encoder.Encode(data); err != nil {
		return fmt.Errorf("cannot encode JSON: %w", err)
	}

	return nil
}
