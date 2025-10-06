package valute

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	errMsgConvertXmlJson = "while convering XML to JSON: %w"
)

type ValCursXML struct {
	Valutes []ValuteXML `xml:"Valute"`
}
type ValuteXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValuteJson struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ConverteXMLtoJSON(valutes []ValuteXML) ([]ValuteJson, error) {
	result := make([]ValuteJson, 0, len(valutes))

	for _, val := range valutes {
		strFloat := strings.ReplaceAll(val.Value, ",", ".")

		num, err := strconv.ParseFloat(strFloat, 64)
		if err != nil {
			return result, fmt.Errorf(errMsgConvertXmlJson, err)
		}

		numCode, err := strconv.Atoi(val.NumCode)
		if err != nil {
			return result, fmt.Errorf(errMsgConvertXmlJson, err)
		}

		result = append(result, ValuteJson{
			NumCode:  numCode,
			CharCode: val.CharCode,
			Value:    num,
		})
	}

	return result, nil
}
