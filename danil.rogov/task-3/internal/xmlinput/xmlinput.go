package xmlinput

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func ReadXML(path string, target any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading %s file: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&target)
	if err != nil {
		return fmt.Errorf("error decoding %s: %w", path, err)
	}

	return nil
}
