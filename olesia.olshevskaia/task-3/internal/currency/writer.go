package currency

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func WriteJSON(currencies interface{}, outputPath string) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic("Error closing XML file: " + err.Error())
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return err
	}

	return nil
}
