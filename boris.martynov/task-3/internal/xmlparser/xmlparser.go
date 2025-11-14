package xmlparser

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func XMLParse(inputxml string, data any) error {
	xmlFile, err := os.Open(inputxml)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}

	defer func() {
		if err := xmlFile.Close(); err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	reader, err := charset.NewReader(xmlFile, "")
	if err != nil {
		return fmt.Errorf("charset reader: %w", err)
	}

	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	return decoder.Decode(data)
}
