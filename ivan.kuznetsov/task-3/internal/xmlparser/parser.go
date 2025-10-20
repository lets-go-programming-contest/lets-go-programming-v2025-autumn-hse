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

type OutputValute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

type ValCurs struct {
	Valutes []OutputValute `xml:"Valute"`
}

func ParseXML(filename string) *ValCurs {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &valCurs
}
