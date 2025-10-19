package jsonoutput

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Tapochek2894/task-3/internal/valute"
)

func WriteJSON(path string, val []valute.ValuteInfo) error {
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		err := os.MkdirAll(dir, 0o755)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
