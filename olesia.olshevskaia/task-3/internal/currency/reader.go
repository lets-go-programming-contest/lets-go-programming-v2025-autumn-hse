package currency

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func Read[T any](path string, xmlTag string) ([]T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open XML file %q: %w", path, err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			panic(fmt.Sprintf("failed to close file: %v", cerr))
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var xmlData struct {
		Data []T `xml:"Valute"`
	}

	if err := decoder.Decode(&xmlData); err != nil {
		return nil, fmt.Errorf("failed to decode XML file %q: %w", path, err)
	}

	return xmlData.Data, nil
}
