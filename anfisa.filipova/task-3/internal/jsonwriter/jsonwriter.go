package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const fileOpenPermission = 0o755

func WriteFileJSON(value any, filePath string) error {
	dir := filepath.Dir(filePath)

	err := os.MkdirAll(dir, fileOpenPermission)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("Error close file: %v", err))
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	err = encoder.Encode(value)
	if err != nil {
		return fmt.Errorf("error encoding file: %w", err)
	}

	return nil
}
