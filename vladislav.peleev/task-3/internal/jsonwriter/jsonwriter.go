package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/VlasfimosY/task-3/internal/models"
)

const dirPerm = 0o755

func SaveJSON(outputPath string, currencies []models.CurrencyJSON) error {
	dir := filepath.Dir(outputPath)

	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("cannot create directory: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("cannot create output file: %w", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to close file %s: %v\n", outputPath, cerr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("cannot encode JSON: %w", err)
	}

	return nil
}
