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

type Value float64

type ValuteXML struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    Value  `xml:"Value"`
}

type ValuteCurs struct {
	Valutes []ValuteXML `xml:"Valute"`
}

func (v *Value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("error decode xml: %w", err)
	}

	str = strings.ReplaceAll(str, ",", ".")

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("error parse float  %w", err)
	}

	*v = Value(value)

	return nil
}

func ParseValuteCursXML(inputFile string) (ValuteCurs, error) {
	var valCurs ValuteCurs

	dataXML, err := os.ReadFile(inputFile)
	if err != nil {
		return valCurs, fmt.Errorf("error read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(dataXML))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&valCurs)
	if err != nil {
		return valCurs, fmt.Errorf("error parsing XML: %w", err)
	}

	return valCurs, nil
}
