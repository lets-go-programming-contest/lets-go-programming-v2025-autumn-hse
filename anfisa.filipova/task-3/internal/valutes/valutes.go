package valutes

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type FloatValue struct {
	Value float64
}

func (value *FloatValue) MarshalJSON() ([]byte, error) {
	floatValue := value.Value
	return json.Marshal(floatValue)
}

func (value *FloatValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var strValue string

	err := decoder.DecodeElement(&strValue, &start)
	if err != nil {
		return fmt.Errorf("error decoding string: %w", err)
	}

	strValue = strings.ReplaceAll(strValue, ",", ".")

	floatValue, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return fmt.Errorf("error converting string to float: %w", err)
	}

	value.Value = floatValue
	return nil
}

type Valute struct {
	NumCode  int        `xml:"NumCode" json:"num_code"`
	CharCode string     `xml:"CharCode" json:"char_code"`
	Value    FloatValue `xml:"Value" json:"value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

func SortValutes(valutes []Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value.Value > valutes[j].Value.Value
	})
}
