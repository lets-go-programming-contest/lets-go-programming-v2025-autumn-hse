package valuteinfo

type value float64

type Valute struct {
	NumCode  int    `xml:"NumCode" json:"NumCode"`
	CharCode string `xml:"CharCode" json:"CharCode"`
	Value    value  `xml:"Value" json:"Value"`
}

type ValuteCurs struct {
	Valutes []Valute `xml:"Valute"`
}
