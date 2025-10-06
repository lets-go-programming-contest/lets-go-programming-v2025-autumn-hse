package valute

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	errMsgConvertXMLJSON = "while convering XML to JSON: %w"
)

type ValCursXML struct {
	Valutes []ValuteXML `xml:"Valute"`
}
type ValuteXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ConverteXMLtoJSON(valutes []ValuteXML) ([]ValuteJSON, error) {
	result := make([]ValuteJSON, 0, len(valutes))

	for _, val := range valutes {
		strFloat := strings.ReplaceAll(val.Value, ",", ".")

		num, err := strconv.ParseFloat(strFloat, 64)
		if err != nil {
			return result, fmt.Errorf(errMsgConvertXMLJSON, err)
		}

		numCode, err := strconv.Atoi(val.NumCode)
		if val.NumCode == "" {
			numCode = 0
		} else if err != nil {
			return result, fmt.Errorf(errMsgConvertXMLJSON, err)
		}

		result = append(result, ValuteJSON{
			NumCode:  numCode,
			CharCode: val.CharCode,
			Value:    num,
		})
	}

	return result, nil
}
