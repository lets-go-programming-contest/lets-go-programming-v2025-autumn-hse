package parsexml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValueFloat float64

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    ValueFloat `json:"value"     xml:"Value"`
}

type ValuteCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func (v *ValueFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("error decode xml: %w", err)
	}

	str = strings.ReplaceAll(str, ",", ".")

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("error parse float  %w", err)
	}

	*v = ValueFloat(value)

	return nil
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
