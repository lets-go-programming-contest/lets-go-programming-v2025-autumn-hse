package jsonout

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const mode = 0o755

func Write(path string, val any) error {
	if err := os.MkdirAll(filepath.Dir(path), mode); err != nil {
		return fmt.Errorf("create %q with mode %o: %w", path, mode, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %q with mode %o: %w", path, mode, err)
	}

	defer func() {
		if ferr := file.Close(); ferr != nil {
			panic(fmt.Errorf("close file %q: %w", path, err))
		}
	}()

	enc := json.NewEncoder(file)
	if err := enc.Encode(&val); err != nil {
		return fmt.Errorf("encode to %q: %w", path, err)
	}

	return nil
}
