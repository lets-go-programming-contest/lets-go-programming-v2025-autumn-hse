package utils

import (
	"os"
	"path/filepath"
)

func EnsureDir(filePath string) error {
	dir := filepath.Dir(filePath)
	return os.MkdirAll(dir, os.ModePerm)
}
