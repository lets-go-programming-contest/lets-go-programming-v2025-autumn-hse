package xmlin

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func Read(path string, val any) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open %q: %w", path, err)
	}

	defer func() {
		if ferr := file.Close(); ferr != nil {
			panic(fmt.Errorf("close file %q: %w", path, err))
		}
	}()

	dec := xml.NewDecoder(file)
	dec.CharsetReader = charset.NewReaderLabel

	if err := dec.Decode(val); err != nil {
		return fmt.Errorf("decode to %q: %w", path, err)
	}

	return nil
}
