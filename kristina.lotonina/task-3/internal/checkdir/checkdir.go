package checkdir

import (
	"fmt"
	"os"
	"path/filepath"
)

func EnsureDir(filePath string) error {
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("creating directory %s: %w", dir, err)
	}

	return nil
}
