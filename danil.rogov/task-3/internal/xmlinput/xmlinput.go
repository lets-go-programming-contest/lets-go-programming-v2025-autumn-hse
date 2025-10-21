package xmlinput

import (
	"bytes"
	"cmp"
	"encoding/xml"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Value float64

type Valute struct {
	NumCode  int
	CharCode string
	Value    Value
}

type ValuteCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func (v *ValuteCurs) Sort(reverse bool) {
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

func (v *Value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string

	err := d.DecodeElement(&str, &start)
	if err != nil {
		return fmt.Errorf("error decoding xml: %w", err)
	}

	str = strings.ReplaceAll(str, ",", ".")

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("error parsing float %w", err)
	}

	*v = Value(value)

	return nil
}

func ReadXML(path string) (ValuteCurs, error) {
	var valCurs ValuteCurs

	data, err := os.ReadFile(path)
	if err != nil {
		return valCurs, fmt.Errorf("error read %s file: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&valCurs)
	if err != nil {
		return valCurs, fmt.Errorf("error parsing %s: %w", path, err)
	}

	return valCurs, nil
}
