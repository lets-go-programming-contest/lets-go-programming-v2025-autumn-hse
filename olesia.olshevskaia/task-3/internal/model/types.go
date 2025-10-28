package model

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Currency struct {
	CodeNum   int     `json:"num_code" xml:"-"`
	CodeChar  string  `json:"char_code" xml:"CharCode"`
	RateValue float64 `json:"value" xml:"-"`
	HasValue  bool    `json:"-" xml:"-"`
	RawNum    string  `xml:"NumCode"`
	RawValue  string  `xml:"Value"`
}
