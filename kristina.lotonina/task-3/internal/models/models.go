package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

type CommaFloat float64

func (cf *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	val := strings.Replace(s, ",", ".", 1)
	f, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err != nil {
		return fmt.Errorf("parsing CommaFloat '%s': %w", s, err)
	}

	*cf = CommaFloat(f)
	return nil
}
