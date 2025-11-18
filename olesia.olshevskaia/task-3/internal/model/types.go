package model

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
}

type Currency struct {
	CodeNum   int     `json:"num_code"  xml:"NumCode"`
	CodeChar  string  `json:"char_code" xml:"CharCode"`
	RateValue float64 `json:"value"     xml:"Value"`
}

type CurrencyValue float64

func (v *CurrencyValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("failed to decode: %w", err)
	}

	str = strings.TrimSpace(str)
	str = strings.Replace(str, ",", ".", 1)

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to parse %q: %w", str, err)
	}

	*v = CurrencyValue(value)
	return nil
}
