package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/JingolBong/task-3/internal/valuteinfo"
)

const dirPerm = 0o755

func Jsonwrite(valuteCurs valuteinfo.ValuteCurs, outputFile string) error {
	directory := filepath.Dir(outputFile)

	if err := os.MkdirAll(directory, dirPerm); err != nil {
		return fmt.Errorf("failed to make dir: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to make file: %w", err)
	}

	jsonData, err := json.MarshalIndent(valuteCurs.Valutes, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if _, err := file.Write(jsonData); err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	return nil
}
