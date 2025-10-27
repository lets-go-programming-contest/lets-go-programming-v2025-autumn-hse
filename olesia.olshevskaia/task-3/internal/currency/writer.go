package currency

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Olesia.Ol/task-3/internal/model"
)

func ParseJSON[T any](outputFile string, data T, dirmode, filemode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(outputFile), dirmode); err != nil {
		return err
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			panic("Error closing JSON file: " + cerr.Error())
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func WriteJSON(outputFile string, currencies []model.Currency, dirmode, filemode os.FileMode) error {
	return ParseJSON(outputFile, currencies, dirmode, filemode)
}
