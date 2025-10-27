package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kef1rch1k/task-3/internal/checkdir"
)

func WriteToFile(data interface{}, filePath string) error {
	if err := checkdir.EnsureDir(filePath); err != nil {
		return fmt.Errorf("creating directory for output file: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("Failed to close file: %v", err))
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encoding JSON: %w", err)
	}

	return nil
}
