package model

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValueFloat float64

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

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    ValueFloat `json:"value"     xml:"Value"`
}

type ValuteCurs struct {
	Valutes []Valute `xml:"Valute"`
}
