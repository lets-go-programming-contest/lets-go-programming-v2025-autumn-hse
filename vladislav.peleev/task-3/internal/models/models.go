package models

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	InputFile  string      `yaml:"input-file"`
	OutputFile string      `yaml:"output-file"`
	DirPerm    os.FileMode `yaml:"dir-perm"`
	FilePerm   os.FileMode `yaml:"file-perm"`
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

type DecimalFloat float64

func (df *DecimalFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var valueStr string
	if err := d.DecodeElement(&valueStr, &start); err != nil {
		return fmt.Errorf("failed to decode Value element: %w", err)
	}

	valueStr = strings.ReplaceAll(valueStr, ",", ".")

	parsedValue, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("cannot parse %q as float: %w", valueStr, err)
	}

	*df = DecimalFloat(parsedValue)

	return nil
}

func (df DecimalFloat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.4f", float64(df))), nil
}

type Currency struct {
	NumCode  int          `json:"num_code"  xml:"NumCode"`
	CharCode string       `json:"char_code" xml:"CharCode"`
	Value    DecimalFloat `json:"value"     xml:"Value"`
}
