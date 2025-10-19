package jsonoutput

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Tapochek2894/task-3/internal/valute"
)

func WriteJSON(path string, valutes []valute.ValuteInfo) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		panic("failed to create directory")
	}

	file, err := os.Create(path)
	if err != nil {
		panic("failed to create output file")
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(valutes)
	if err != nil {
		return err
	}
	return nil
}
