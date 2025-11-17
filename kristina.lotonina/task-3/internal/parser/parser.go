package parser

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func ParseXML(inputxml string, data any) error {
	file, err := os.Open(inputxml)
	if err != nil {
		return fmt.Errorf("opening XML file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("failed to close file: %v", err))
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(data); err != nil {
		return fmt.Errorf("decoding XML: %w", err)
	}

	return nil
}
