package model

type Config struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
}

type CurrencyValue float64

//type Currency struct {
//	CodeNum   int           `json:"num_code"  xml:"NumCode"`
//	CodeChar  string        `json:"char_code" xml:"CharCode"`
//	RateValue CurrencyValue `json:"value"     xml:"Value"`
//}

type Currency struct {
	CodeNum   int     `json:"num_code"  xml:"-"`
	CodeChar  string  `json:"char_code" xml:"CharCode"`
	RateValue float64 `json:"value"     xml:"-"`
	HasValue  bool    `json:"-"         xml:"-"`
	RawNum    string  `xml:"NumCode"`
	RawValue  string  `xml:"Value"`
}
