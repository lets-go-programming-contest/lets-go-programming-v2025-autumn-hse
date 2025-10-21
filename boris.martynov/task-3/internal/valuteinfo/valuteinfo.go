package valuteinfo

type Valute struct {
	NumCode  int     `xml:"NumCode" json:"NumCode"`
	CharCode string  `xml:"CharCode" json:"CharCode"`
	Value    float64 `xml:"Value" json:"Value"`
}

type ValuteCurs struct {
	Valutes []Valute `xml:"Valute"`
}
