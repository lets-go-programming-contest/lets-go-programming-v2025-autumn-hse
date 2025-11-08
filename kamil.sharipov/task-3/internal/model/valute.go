package model

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValuteValue float64

func (v *ValuteValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	s = strings.Replace(s, ",", ".", 1)

	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value %q: %w", s, err)
	}

	*v = ValuteValue(value)
	return nil
}

type Valute struct {
	NumCode  int         `xml:"NumCode" json:"num_code"`
	CharCode string      `xml:"CharCode" json:"char_code"`
	Value    ValuteValue `xml:"Value" json:"value"`
}
