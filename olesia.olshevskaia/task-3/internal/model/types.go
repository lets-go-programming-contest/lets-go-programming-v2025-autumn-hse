package model

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
