package currency

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func WriteJSON(currencies interface{}, outputPath string) {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic("I can't create a directory: " + err.Error())
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic("I can't create a JSON file: " + err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		panic("Write error JSON: " + err.Error())
	}
}
