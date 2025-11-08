package model

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValuteValue float64

func (v *ValuteValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("failed to decode XML element: %w", err)
	}

	str = strings.Replace(str, ",", ".", 1)

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value %q: %w", str, err)
	}

	*v = ValuteValue(value)

	return nil
}

type Valute struct {
	NumCode  int         `json:"num_code"  xml:"NumCode"`
	CharCode string      `json:"char_code" xml:"CharCode"`
	Value    ValuteValue `json:"value"     xml:"Value"`
}
