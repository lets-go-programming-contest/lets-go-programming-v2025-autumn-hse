package jsonoutput

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tapochek2894/task-3/internal/valute"
)

const (
	dirMode  = 0o555
	fileMode = 0o666
)

func WriteJSON(path string, valutes valute.Valutes) error {
	err := os.MkdirAll(filepath.Dir(path), dirMode)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %w", path, err)
	}

	output, err := json.MarshalIndent(valutes, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.WriteFile(path, output, fileMode)
	if err != nil {
		return fmt.Errorf("error writing %s file: %w", path, err)
	}

	return nil
}
