package writerjson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ekaterina-101/task-3/internal/parsexml"
)

func ParseJSON[T any](outputFile string, data T, dirmode, filemode os.FileMode) error {
	outputJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(outputFile), dirmode)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	err = os.WriteFile(outputFile, outputJSON, filemode)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}

func WriteFileJSON(outputFile string, valCurs parsexml.ValuteCurs, dirmode, filemode os.FileMode) error {
	err := ParseJSON(outputFile, valCurs.Valutes, dirmode, filemode)

	return err
}
