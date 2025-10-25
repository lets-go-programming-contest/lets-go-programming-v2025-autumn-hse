package writerjson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ekaterina-101/task-3/internal/parsexml"
)

const (
	dirMode  = 0o755
	fileMode = 0o600
)

func WriteFileJSON(outputFile string, valCurs parsexml.ValuteCurs) error {
	outputJSON, err := json.MarshalIndent(valCurs.Valutes, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(outputFile), dirMode)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	err = os.WriteFile(outputFile, outputJSON, fileMode)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}
