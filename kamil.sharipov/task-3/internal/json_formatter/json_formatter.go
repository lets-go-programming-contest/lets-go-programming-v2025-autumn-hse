package json

import (
	"encoding/json"
	"fmt"

	xml "github.com/kamilSharipov/task-3/internal/xml_parser"
)

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func FormateJSON(valutesXML []xml.Valute) ([]byte, error) {
	valutes := make([]ValuteJSON, len(valutesXML))

	for index, valute := range valutesXML {
		numCode := valute.NumCode
		charCode := valute.CharCode

		value := 0.0

		if valute.Value != "" {
			var err error

			value, err = valute.GetValue()
			if err != nil {
				return nil, fmt.Errorf("invalid Value %q: %w", valute.Value, err)
			}
		}

		valutes[index].CharCode = charCode
		valutes[index].NumCode = numCode
		valutes[index].Value = value
	}

	bytes, err := json.Marshal(valutes)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	return bytes, nil
}
