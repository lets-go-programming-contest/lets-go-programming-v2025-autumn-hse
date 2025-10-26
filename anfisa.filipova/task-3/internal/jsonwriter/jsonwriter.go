package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Anfisa111/task-3/internal/valutes"
)

const fileOpenPermission = 0o755

func WriteCurrenciesJSON(currencies []valutes.Valute, filePath string) error {
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

	err = encoder.Encode(currencies)
	if err != nil {
		return fmt.Errorf("error encoding file: %w", err)
	}

	return nil
}
