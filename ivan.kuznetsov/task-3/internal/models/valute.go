package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type CommaFloat float64

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func (cf *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("decode element failed: %w", err)
	}

	str = strings.Replace(str, ",", ".", 1)

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("parse float failed: %w", err)
	}

	*cf = CommaFloat(value)

	return nil
}
