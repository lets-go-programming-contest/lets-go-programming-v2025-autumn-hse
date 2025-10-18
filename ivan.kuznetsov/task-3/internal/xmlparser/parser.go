package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type CommaFloat float64

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
		panic(err.Error())
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(err.Error())
	}

	return &valCurs
}
