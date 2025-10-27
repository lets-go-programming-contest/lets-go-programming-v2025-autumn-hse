package model

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Currency struct {
	CodeNum   int     `json:"num_code"`
	CodeChar  string  `xml:"CharCode" json:"char_code"`
	RateValue float64 `json:"value"`
	HasValue  bool    `json:"-"`
	RawNum    string  `xml:"NumCode"`
	RawValue  string  `xml:"Value"`
}

func (c *Currency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var aux struct {
		CodeNum   string `xml:"NumCode"`
		CodeChar  string `xml:"CharCode"`
		RateValue string `xml:"Value"`
	}

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	if strings.TrimSpace(aux.CodeNum) != "" {
		if n, err := strconv.Atoi(aux.CodeNum); err == nil {
			c.CodeNum = n
		}
	}

	if strings.TrimSpace(aux.RateValue) != "" {
		str := strings.ReplaceAll(aux.RateValue, ",", ".")
		if v, err := strconv.ParseFloat(str, 64); err == nil {
			c.RateValue = v
			c.HasValue = true
		}
	}

	c.CodeChar = aux.CodeChar
	return nil
}
