package xmlreader

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/emersion/go-message/charset"
)

func ReadFileXML(filepath string, value any) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(file))
	decoder.CharsetReader = charset.Reader

	err = decoder.Decode(value)
	if err != nil {
		return fmt.Errorf("error decoding file: %w", err)
	}

	return nil
}
