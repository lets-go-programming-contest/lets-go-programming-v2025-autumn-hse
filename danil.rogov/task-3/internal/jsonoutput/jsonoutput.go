package jsonoutput

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Tapochek2894/task-3/internal/valute"
)

func WriteJSON(path string, val []valute.ValuteInfo) error {
	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(val)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
