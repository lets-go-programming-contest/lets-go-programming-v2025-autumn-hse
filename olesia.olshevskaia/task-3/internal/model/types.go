package model

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type CurrencyXML struct {
	CodeNum   string `xml:"NumCode"`
	CodeChar  string `xml:"CharCode"`
	RateValue string `xml:"Value"`
}

type CurrenciesXML struct {
	Currencies []CurrencyXML `xml:"Valute"`
}

type Currency struct {
	CodeNum   int     `json:"num_code"`
	CodeChar  string  `json:"char_code"`
	RateValue float64 `json:"value"`
	HasValue  bool    `json:"-"`
}
