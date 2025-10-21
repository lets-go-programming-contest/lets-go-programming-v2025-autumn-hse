package jsonoutput

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tapochek2894/task-3/internal/valute"
)

const mode = 0o755

func WriteJSON(path string, valutes valute.Valutes) error {
	err := os.MkdirAll(filepath.Dir(path), mode)
	if err != nil {
		return fmt.Errorf("error creating directory %s with mode %o: %w", path, mode, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file %s with mode %o: %w", path, mode, err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(fmt.Errorf("error closing file %s: %w", path, err))
		}
	}()

	enc := json.NewEncoder(file)
	if err := enc.Encode(&valutes); err != nil {
		return fmt.Errorf("error encoding to file %s: %w", path, err)
	}

	return nil
}
