package parsexml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func ParseXML[T any](inputFile string) (T, error) {
	var result T

	dataXML, err := os.ReadFile(inputFile)
	if err != nil {
		return result, fmt.Errorf("failed to read XML: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(dataXML))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&result)
	if err != nil {
		return result, fmt.Errorf("error parsing XML: %w", err)
	}

	return result, nil
}
