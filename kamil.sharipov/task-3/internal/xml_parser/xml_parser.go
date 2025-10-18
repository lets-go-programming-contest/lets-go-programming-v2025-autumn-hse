package xml

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"

	"golang.org/x/net/html/charset"
)

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

var errNoValutes = errors.New("XML contains no valutes")

func ParseXML(xmlData []byte) ([]Valute, error) {
	xmlDecoder := xml.NewDecoder(bytes.NewReader(xmlData))
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	var temp struct {
		Valutes []Valute `xml:"Valute"`
	}

	if err := xmlDecoder.Decode(&temp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	if len(temp.Valutes) == 0 {
		return nil, errNoValutes
	}

	return temp.Valutes, nil
}
