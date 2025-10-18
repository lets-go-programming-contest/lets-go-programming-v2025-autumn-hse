package xml

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func (v *Valute) GetValue() (float64, error) {
	s := strings.Replace(v.Value, ",", ".", 1)

	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse float %q: %w", s, err)
	}

	return value, nil
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
