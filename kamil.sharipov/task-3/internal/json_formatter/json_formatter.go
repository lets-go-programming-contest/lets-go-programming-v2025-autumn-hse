package json

import (
	"encoding/json"
	"sort"

	xml "github.com/kamilSharipov/task-3/internal/xml_parser"
)

type ValuteJSON struct {
	NumCode  string `json:"num_code"`
	CharCode string `json:"char_code"`
	Value    string `json:"value"`
}

func FormateJSON(valCurs *xml.ValCurs) ([]byte, error) {

	valutes := make([]ValuteJSON, len(valCurs.Valutes))
	for i, valute := range valCurs.Valutes {
		valutes[i].CharCode = valute.CharCode
		valutes[i].NumCode = valute.NumCode
		valutes[i].Value = valute.Value
	}

	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})

	bytes, err := json.Marshal(valutes)
	if err != nil {
		return nil, nil
	}
	return bytes, nil
}
