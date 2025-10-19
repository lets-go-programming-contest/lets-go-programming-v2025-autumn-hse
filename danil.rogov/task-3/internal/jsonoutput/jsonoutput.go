package jsonoutput

import (
	"encoding/json"
	"os"

	"github.com/Tapochek2894/task-3/internal/valute"
)

func WriteJSON(path string, val []valute.ValuteInfo) error {
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
