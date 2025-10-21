package jsonoutput

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tapochek2894/task-3/internal/xmlinput"
)

const (
	dirMode  = 0o555
	fileMode = 0o666
)

func CreateValuteCursJSON(path string, valCurs xmlinput.ValuteCurs) error {
	outputJSON, err := json.MarshalIndent(valCurs, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling %s: %w", path, err)
	}

	err = WriteJSON(path, outputJSON)
	if err != nil {
		return fmt.Errorf("error writing %s: %w", path, err)
	}

	return nil
}

func WriteJSON(outputFile string, outputJSON []byte) error {
	err := os.MkdirAll(filepath.Dir(outputFile), dirMode)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	err = os.WriteFile(outputFile, outputJSON, fileMode)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}
