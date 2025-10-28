package valuteinfo

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Value float64

type Valute struct {
	NumCode  int    `xml:"NumCode" json:"num_code"`
	CharCode string `xml:"CharCode" json:"char_code"`
	Value    Value  `xml:"Value" json:"value"`
}

type ValuteCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func (v *Value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string

	if err := d.DecodeElement(&str, &start); err != nil {

		return fmt.Errorf("while decoding: %w", err)
	}
	str = strings.TrimSpace(str)
	value, err := strconv.ParseFloat(strings.ReplaceAll(str, ",", "."), 64)

	if err != nil {

		return fmt.Errorf("while replacing dots: %w", err)
	}

	*v = Value(value)

	return nil
}
