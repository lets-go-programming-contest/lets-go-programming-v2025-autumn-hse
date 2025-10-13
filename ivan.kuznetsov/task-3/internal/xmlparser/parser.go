package xmlparser

import (
	"encoding/xml"
	"os"
	"strconv"
	"strings"
)

type CommaFloat float64

func (cf *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	s = strings.Replace(s, ",", ".", 1)

	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*cf = CommaFloat(value)

	return nil
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string     `xml:"NumCode"`
	CharCode string     `xml:"CharCode"`
	Value    CommaFloat `xml:"Value"`
}

func ParseXML(filename string) *ValCurs {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("failed to read XML file")
	}

	var valCurs ValCurs

	err = xml.Unmarshal(data, &valCurs)
	if err != nil {
		panic("failed to parse XML")
	}

	return &valCurs
}
