package valute

import (
	"cmp"
	"encoding/xml"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Value float64

type Valute struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Value  `json:"value"     xml:"Value"`
}

type Valutes struct {
	Valutes []Valute `xml:"Valute"`
}

func (v *Valutes) Sort(reverse bool) {
	switch reverse {
	case true:
		slices.SortFunc(v.Valutes, func(a, b Valute) int {
			return cmp.Compare(b.Value, a.Value)
		})
	case false:
		slices.SortFunc(v.Valutes, func(a, b Valute) int {
			return -cmp.Compare(b.Value, a.Value)
		})
	}
}

func (v *Value) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var stringValue string

	err := decoder.DecodeElement(&stringValue, &start)
	if err != nil {
		return fmt.Errorf("error decoding xml: %w", err)
	}

	stringValue = strings.ReplaceAll(stringValue, ",", ".")

	value, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return fmt.Errorf("error parsing float: %w", err)
	}

	*v = Value(value)

	return nil
}
