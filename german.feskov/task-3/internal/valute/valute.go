package valute

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValCursXML struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    Decimal `json:"value"     xml:"Value"`
}

type Decimal float64

func (d *Decimal) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := dec.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("decode element Decimal: %w", err)
	}

	if str == "" {
		*d = Decimal(0)

		return nil
	}

	str = strings.ReplaceAll(str, ",", ".")

	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("parse xml Decimal %q: %w", str, err)
	}

	*d = Decimal(val)

	return nil
}
