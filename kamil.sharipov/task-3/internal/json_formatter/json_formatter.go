package json

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	xml "github.com/kamilSharipov/task-3/internal/xml_parser"
)

type ValuteJSON struct {
	NumCode  int    `json:"num_code"`
	CharCode string `json:"char_code"`
	Value    string `json:"value"`
}

func FormateJSON(valCurs *xml.ValCurs) ([]byte, error) {

	valutes := make([]ValuteJSON, len(valCurs.Valutes))
	for i, valute := range valCurs.Valutes {
		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			fmt.Println("invalid NumCode")
			return nil, err
		}

		valutes[i].CharCode = valute.CharCode
		valutes[i].NumCode = numCode
		valutes[i].Value = valute.Value
	}

	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})

	bytes, err := json.Marshal(valutes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
