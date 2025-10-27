package parsexml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/Ekaterina-101/task-3/internal/castomparsexml"
	"golang.org/x/net/html/charset"
)

type Valute struct {
	NumCode  int                       `json:"num_code"  xml:"NumCode"`
	CharCode string                    `json:"char_code" xml:"CharCode"`
	Value    castomparsexml.ValueFloat `json:"value"     xml:"Value"`
}

type ValuteCurs struct {
	Valutes []Valute `xml:"Valute"`
}

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

func ParseValuteCursXML(inputFile string) (ValuteCurs, error) {
	valCurs, err := ParseXML[ValuteCurs](inputFile)
	if err != nil {
		return ValuteCurs{}, err
	}

	return valCurs, nil
}
