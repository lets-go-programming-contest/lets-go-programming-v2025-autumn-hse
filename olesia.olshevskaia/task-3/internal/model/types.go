package model

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Currency struct {
	CodeNum   int     `xml:"NumCode"   json:"num_code"`
	CodeChar  string  `xml:"CharCode"  json:"char_code"`
	RateValue float64 `xml:"Value"     json:"value"`
	HasValue  bool    `json:"-"`
	RawNum    string  `xml:"-"`
	RawValue  string  `xml:"-"`
}
