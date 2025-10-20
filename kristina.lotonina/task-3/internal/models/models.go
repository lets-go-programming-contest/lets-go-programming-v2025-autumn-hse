package models

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	ValueStr string  `xml:"Value"`
	Value    float64 `xml:"-"`
}

type OutputValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}