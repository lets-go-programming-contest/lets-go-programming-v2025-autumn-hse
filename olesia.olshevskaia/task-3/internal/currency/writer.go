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
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return err
	}

	return nil
}
