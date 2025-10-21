package jsonoutput

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const mode = 0o777

func WriteJSON(path string, valutes any) error {
	err := os.MkdirAll(filepath.Dir(path), mode)
	if err != nil {
		return fmt.Errorf("error creating %s with mode %o: %w", path, mode, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %w", path, err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(fmt.Errorf("error closing file %s: %w", path, err))
		}
	}()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&valutes); err != nil {
		return fmt.Errorf("error encoding to %s: %w", path, err)
	}

	return nil
}
