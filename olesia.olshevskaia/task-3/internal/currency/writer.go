package currency

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func WriteJSON(currencies interface{}, outputPath string) {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic("Cannot create directory: " + err.Error())
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic("Cannot create JSON file: " + err.Error())
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic("Error closing JSON file: " + err.Error())
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(currencies); err != nil {
		panic("Cannot encode JSON: " + err.Error())
	}
}
