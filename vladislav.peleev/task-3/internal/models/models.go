package models

import (
	"encoding/xml"
	"io"

	"golang.org/x/text/encoding/charmap"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Currency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

type CurrencyJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func GetCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}
	return input, nil
}
