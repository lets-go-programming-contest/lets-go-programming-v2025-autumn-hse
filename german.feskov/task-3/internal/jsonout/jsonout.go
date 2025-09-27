package jsonout

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const mode = 0o755

const errMsg = "while write JSON %q: %w"

func Write(path string, val any) (err error) {
	if err := os.MkdirAll(filepath.Dir(path), mode); err != nil {
		return fmt.Errorf(errMsg, path, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf(errMsg, path, err)
	}

	defer func() {
		if ferr := file.Close(); ferr != nil && err == nil {
			err = fmt.Errorf(errMsg, path, ferr)
		}
	}()

	data, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf(errMsg, path, err)
	}

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf(errMsg, path, err)
	}

	return nil
}
