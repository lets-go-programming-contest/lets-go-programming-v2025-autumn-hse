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

type CurrencyRates struct {
	Currencies []Currency `xml:"Valute"`
}

type CurrencyValue float64

type Currency struct {
	CodeNum   int           `json:"num_code"  xml:"NumCode"`
	CodeChar  string        `json:"char_code" xml:"CharCode"`
	RateValue CurrencyValue `json:"value"     xml:"Value"`
}

func (v *CurrencyValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("error decode xml: %w", err)
	}

	str = strings.ReplaceAll(str, ",", ".")

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("error parse float: %w", err)
	}

	*v = CurrencyValue(value)
	return nil
}
