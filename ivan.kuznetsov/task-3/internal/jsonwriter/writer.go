package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kuzid-17/task-3/internal/models"
)

const mode = 0o755

func WriteJSON(filename string, valutes []models.Valute) error {
	dir := filepath.Dir(filename)

	err := os.MkdirAll(dir, mode)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	output := make([]models.Valute, 0, len(valutes))

	for _, valute := range valutes {
		output = append(output, models.Valute{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    models.CommaFloat(valute.Value),
		})
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(output)
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}
